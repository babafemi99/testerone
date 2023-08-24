package load

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"
)

func (c *CustomReq) RunAfter() ([]ResponseData, error) {
	return c.runX(c.RunAfterDuration, c.RunDuration)
}

func (c *CustomReq) Run() (ResponseData, error) {
	err := c.validate()
	if err != nil {
		return ResponseData{}, fmt.Errorf(" Run: %w", err)
	}

	results := make(chan ResponseTime, c.NumberOfRequests)
	done := make(chan bool, c.NumberOfRequests)

	for i := 1; i <= c.NumberOfRequests; i++ {
		go c.loadCustomTarget2(results, done)
	}

	for i := 1; i <= c.NumberOfRequests; i++ {
		<-done
	}

	close(results)
	close(done)

	var sumTime float64
	var errorCount int
	var SuccessCount int
	var responseData ResponseData
	responseData.MinimumTime = math.Inf(1)
	//log.Printf("Index\tTime\tDelivered")

	for result := range results {
		//log.Printf("%d\t%f\t%t", result.Index, result.Time, result.Success)
		sumTime += result.Time
		if result.Time > responseData.MaximumTime {
			responseData.MaximumTime = result.Time
		}

		if result.Time < responseData.MinimumTime {
			responseData.MinimumTime = result.Time
		}

		if result.Success {
			SuccessCount += 1
		} else {
			errorCount += 1
		}

		responseData.Responses = append(responseData.Responses, result)
	}

	if c.Interval == 0 {
		c.Interval = 1
	}

	responseData.Responses = processReq(responseData.Responses, c.Interval)

	responseData.AverageResponseTime = sumTime / float64(c.NumberOfRequests)
	responseData.SuccessRate = float64(SuccessCount) / float64(c.NumberOfRequests) * 100
	responseData.ErrorRate = 100 - responseData.SuccessRate

	return responseData, nil

}

func (c *CustomReq) loadCustomTarget2(ch chan ResponseTime, done chan bool) {
	defer func() {
		done <- true
	}()

	// start time
	start := time.Now()

	// call the custom function method and handle in case of errors
	err := c.callCustomFunc()
	if err != nil {
		ch <- ResponseTime{
			Time:    time.Since(start).Seconds(),
			Success: false,
		}
		errCount++
		return
	}

	ch <- ResponseTime{
		Time:    time.Since(start).Seconds(),
		Success: true,
	}
}

func (c *CustomReq) callCustomFunc() error {

	for _, fn := range c.Func2 {
		err := fn.hitReq()
		if err != nil {
			return fmt.Errorf("fn.hitReq error: %w", err)
		}
	}
	return nil
}

func (cf *CustomFunction) hitReq() error {

	var req *http.Request
	var err error

	switch cf.Method {

	case "POST":
		req, err = http.NewRequest(cf.Method, cf.URL, bytes.NewBuffer(cf.Body))
		if err != nil {
			return fmt.Errorf("http.NewRequest Error: %w", err)
		}
	case "GET":
		req, err = http.NewRequest(cf.Method, cf.URL, nil)
		if err != nil {
			return fmt.Errorf("http.NewRequest Error: %w", err)
		}

	default:
		return errors.New(`Method must be "POST" or "GET" `)

	}

	cl := http.Client{}

	for _, header := range cf.Headers {
		err := header.validate()
		if err != nil {
			return fmt.Errorf("header.validate: %w", err)
		}

		req.Header.Set(header.Name, header.Value)
	}

	for _, cookie := range cf.Cookies {

		dur, err := cookie.validate()
		if err != nil {
			return fmt.Errorf("cookie.Validate: %w", err)
		}

		ck := &http.Cookie{
			Name:    cookie.Name,
			Value:   cookie.ExpiresAt,
			Expires: time.Now().Add(dur),
		}

		req.AddCookie(ck)
	}

	res, err := cl.Do(req)
	if err != nil {
		return fmt.Errorf("cl.Do error: %w", err)
	}

	res.Body.Close()
	return nil
}

func (c *CustomReq) runX(timeInterval time.Duration, n int) ([]ResponseData, error) {
	var data []ResponseData
	ticker := time.NewTicker(timeInterval)
	for i := 0; i < n; i++ {
		run, err := c.Run()
		if err != nil {
			return []ResponseData{}, fmt.Errorf("c.Run > %w", err)
		}
		data = append(data, run)
		log.Println("request"+" "+"completed", i+1)
		<-ticker.C
	}
	return data, nil
}
