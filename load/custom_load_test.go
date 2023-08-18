package load

import (
	"net/http"
	"testing"
)

func Test_Custom(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:1010/ping", nil)
	if err != nil {
		return
	}

	req22 := CustomReq{
		ReqType:          "custom",
		NumberOfRequests: 10000,
		URL:              "",
		Interval:         0,
		Func:             req,
	}
	req22.Run()
}
