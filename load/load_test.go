package load

import (
	"testing"
)

func TestReq_Run(t *testing.T) {
	req := Req{
		NumberOfRequests: 50,
		URL:              "http://localhost:1010/ping",
	}
	data := req.Run()

	if len(data.Responses) != req.NumberOfRequests {
		t.Errorf("Failed number is meant to be equal")
	}
}
