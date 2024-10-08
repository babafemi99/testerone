package load

import (
	"testing"
)

func TestReq_Run(t *testing.T) {
	req := Req{
		NumberOfRequests: 100,
		//URL:              "http://localhost:1010/ping",
		//URL:      "https://www.google.com/",
		URL:      "http://localhost:2020",
		Interval: 1,
	}
	data, _ := req.run()

	if len(data.Responses) != req.NumberOfRequests {
		t.Errorf("Failed number is meant to be equal")
	}
	//table.RenderTable(data)
}
