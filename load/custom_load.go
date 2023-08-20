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

func (c *CustomReq) Run() (ResponseData, error) {
	err := c.validate()
	if err != nil {
		return ResponseData{}, err
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

	fmt.Printf("response average time :%.4fs\n", responseData.AverageResponseTime)
	fmt.Printf("response error rate: %.2f%%\n", responseData.ErrorRate)
	fmt.Printf("response success rate: %.2f%%\n", responseData.SuccessRate)
	fmt.Printf("response maximum time :%.4fs\n", responseData.MaximumTime)
	fmt.Printf("response minimum time :%.4fs\n", responseData.MinimumTime)

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
			return err
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
			log.Println("error", err)
			return err
		}
	case "GET":
		req, err = http.NewRequest(cf.Method, cf.URL, nil)
		if err != nil {
			log.Println("error", err)
			return err
		}

	default:
		return errors.New(`Method must be "POST" or "GET" `)

	}

	cl := http.Client{}

	res, err := cl.Do(req)
	if err != nil {
		return err
	}

	res.Body.Close()
	return nil
}
