package load

func processReq(responseArr []ResponseTime, interval int) []ResponseTime {
	var output []ResponseTime
	for i := 0; i < len(responseArr); i += interval {
		//log.Printf("%d\t%f\t%t", responseArr[i].Index, responseArr[i].Time, responseArr[i].Success)
		output = append(output, responseArr[i])
	}

	return output
}
