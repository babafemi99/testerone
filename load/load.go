package load

import (
	"log"
	"net/http"
	"time"
)

func (r *Req) Run() ResponseData {
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
	log.Printf("Index\tTime\tDelivered")

	for result := range results {
		log.Printf("%d\t%f\t%t", result.Index, result.Time, result.Success)
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

	responseData.AverageResponseTime = sumTime / float64(r.NumberOfRequests)
	responseData.SuccessRate = float64(SuccessCount/r.NumberOfRequests) * 100
	responseData.ErrorRate = 100 - responseData.SuccessRate

	log.Println("response average time- ", responseData.AverageResponseTime)
	log.Println("response error rate- ", responseData.ErrorRate)
	log.Println("response success rate- ", responseData.SuccessRate)
	log.Println("response maximum time- ", responseData.MaximumTime)
	log.Println("response minimum time- ", responseData.MinimumTime)
	return responseData
}

func (r *Req) loadTarget(ch chan ResponseTime, done chan bool, index int) {

	defer func() {
		done <- true
	}()

	start := time.Now()
	res, err := http.Get(r.URL)
	if err != nil {
		log.Println("error hitting the server")
		ch <- ResponseTime{
			Time:    time.Since(start).Seconds(),
			Success: false,
		}
		return
	}
	defer res.Body.Close()

	ch <- ResponseTime{
		Time:    time.Since(start).Seconds(),
		Success: true,
	}

}
