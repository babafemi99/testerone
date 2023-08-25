package load

import (
	"fmt"
	"time"
)

func processReq(responseArr []ResponseTime, interval int) []ResponseTime {
	var output []ResponseTime
	for i := 0; i < len(responseArr); i += interval {
		//log.Printf("%d\t%f\t%t", responseArr[i].Index, responseArr[i].Time, responseArr[i].Success)
		output = append(output, responseArr[i])
	}

	return output
}

func getDuration(str string) (time.Duration, error) {
	duration, err := time.ParseDuration(str)
	if err != nil {
		return 0, fmt.Errorf(str)
	}
	return duration, nil
}
