package load

import (
	"log"
	"time"
)

// I want to simulate a request / Requests for n number of times but with a t time interval for a total on x times

// write a fn that runs n times but after a time interval of x times : RunXN

func RunXn(timeInterval time.Duration, n int) {
	ticker := time.NewTicker(timeInterval)

	for i := 0; i < n; i++ {
		CustomFN(i)
		<-ticker.C
	}
}

func CustomFN(in int) {
	log.Println("simulating work !!! ", in+1)
}
