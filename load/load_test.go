package load

import (
	"net/http"
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

func Test_Custom(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:1010/ping", nil)
	if err != nil {
		return
	}

	req22 := Req{
		ReqType:          "custom",
		NumberOfRequests: 10000,
		URL:              "",
		Interval:         0,
		Func:             req,
	}
	req22.Run()
}
