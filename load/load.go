package load

import (
	"fmt"
	"math"
	"net/http"
	"time"
)

func (r *Req) Run() (ResponseData, error) {

	err := r.validate()
	if err != nil {
		return ResponseData{}, fmt.Errorf("r.validate: %w", err)
	}

	results := make(chan ResponseTime, r.NumberOfRequests)
	done := make(chan bool, r.NumberOfRequests)

	for i := 1; i <= r.NumberOfRequests; i++ {
		go r.loadTarget(results, done, i)
	}

	for i := 1; i <= r.NumberOfRequests; i++ {
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

	if r.Interval == 0 {
		r.Interval = 1
	}

	processReq(responseData.Responses, r.Interval)

	responseData.AverageResponseTime = sumTime / float64(r.NumberOfRequests)
	responseData.SuccessRate = float64(SuccessCount) / float64(r.NumberOfRequests) * 100
	responseData.ErrorRate = 100 - responseData.SuccessRate

	fmt.Printf("response average time :%.4fs\n", responseData.AverageResponseTime)
	fmt.Printf("response error rate: %.2f%%\n", responseData.ErrorRate)
	fmt.Printf("response success rate: %.2f%%\n", responseData.SuccessRate)
	fmt.Printf("response maximum time :%.4fs\n", responseData.MaximumTime)
	fmt.Printf("response minimum time :%.4fs\n", responseData.MinimumTime)

	return responseData, nil
}

var errCount int

func (r *Req) loadTarget(ch chan ResponseTime, done chan bool, index int) {

	defer func() {
		done <- true
	}()

	start := time.Now()
	res, err := http.Get(r.URL)
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
