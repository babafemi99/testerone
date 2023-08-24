package load

import (
	"errors"
	"time"
)

type ResponseData struct {
	AverageResponseTime float64        `json:"average_response_time"`
	ErrorRate           float64        `json:"error_rate"`
	SuccessRate         float64        `json:"success_rate"`
	MinimumTime         float64        `json:"minimum_time"`
	MaximumTime         float64        `json:"maximum_time"`
	Responses           []ResponseTime `json:"responses"`
}
type ResponseTime struct {
	Index   int     `json:"index"`
	Time    float64 `json:"time"`
	Success bool    `json:"success"`
}

type Req struct {
	NumberOfRequests int           `json:"number_of_requests"`
	URL              string        `json:"url"`
	Interval         int           `json:"interval"`
	RunAfterDuration time.Duration `json:"run_after_duration"`
	RunDuration      int           `json:"run_duration"`
}

func (r *Req) validate() error {
	if r.NumberOfRequests < r.Interval {
		return errors.New("number of requests must be more than intervals")
	}
	return nil
}

// CustomReq :custom requests have more options for making a http Requests
type CustomReq struct {
	ReqType          string           `json:"req_type"`
	NumberOfRequests int              `json:"number_of_requests"`
	URL              string           `json:"url"`
	Interval         int              `json:"interval"`
	Func2            []CustomFunction `json:"func_2"`
	RunAfterDuration time.Duration    `json:"run_after_duration"`
	RunDuration      int              `json:"run_duration"`
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (h *Header) validate() error {
	if h.Name == "" {
		return errors.New("header must have a name")
	}

	if h.Value == "" {
		return errors.New("header must have a value")
	}

	return nil
}

type Cookie struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	ExpiresAt string `json:"expires_at"`
}

func (c *Cookie) validate() (time.Duration, error) {
	if c.Name == "" {
		return 0, errors.New("cookie must have a name")
	}

	if c.Value == "" {
		return 0, errors.New("cookie must have a value")
	}

	dur, err := time.ParseDuration(c.ExpiresAt)
	if err != nil {
		return 0, errors.New("invalid cookie duration")
	}

	return dur, nil
}

type CustomFunction struct {
	Method  string   `json:"method"`
	URL     string   `json:"url"`
	Body    []byte   `json:"body"`
	Headers []Header `json:"headers"`
	Cookies []Cookie `json:"cookies"`
}

func (c *CustomReq) validate() error {
	if c.NumberOfRequests < c.Interval {
		return errors.New("number of requests must be more than intervals")
	}
	return nil
}
