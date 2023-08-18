package load

import (
	"testing"
)

func TestReq_Run(t *testing.T) {
	req := Req{
		NumberOfRequests: 250,
		URL:              "http://localhost:1010/ping",
		//URL:      "https://www.google.com/",
		//URL:      "http://localhost:2020/process",
		Interval: 1,
	}
	data, _ := req.Run()

	if len(data.Responses) != req.NumberOfRequests {
		t.Errorf("Failed number is meant to be equal")
	}
	//table.RenderTable(data)
}
