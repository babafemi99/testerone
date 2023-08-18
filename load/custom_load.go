package load

import (
	"fmt"
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
		go c.loadCustomTarget(results, done)
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

	processReq(responseData.Responses, c.Interval)

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

func (c *CustomReq) loadCustomTarget(ch chan ResponseTime, done chan bool) {
	defer func() {
		done <- true
	}()

	start := time.Now()

	cl := http.Client{}

	res, err := cl.Do(c.Func)
	if err != nil {
		//log.Println("error hitting the server", err)
		ch <- ResponseTime{
			Time:    time.Since(start).Seconds(),
			Success: false,
		}
		errCount++
		return
	}
	defer res.Body.Close()

	ch <- ResponseTime{
		Time:    time.Since(start).Seconds(),
		Success: true,
	}
}
