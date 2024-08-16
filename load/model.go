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
	// NumberOfRequests specifies the number of requests to be made concurrently
	NumberOfRequests int `json:"number_of_requests"`

	// Interval specifies intervals between displaying output for completed requests,
	// For example, if NumberOfRequests is 100 and Interval is 10,
	// the output will be displayed for the 1st, 10th, 20th, ..., 100th request
	Interval int `json:"interval"`

	//Func2 represents the endpoint to be hit during the simulation
	Func2 []CustomFunction `json:"func_2"`

	// RunAfterDuration specifies the duration to pause before another batch job continues.
	// If RunAfterDuration is set to, for example, 5 * time.Second, the simulation will run
	// the specified number of requests, then wait for 5 seconds before running the next batch.
	// This pattern continues until the total RunDuration is reached.
	RunAfterDuration time.Duration `json:"run_after_duration"`

	// RunDuration specifies the total number of times the batch job will run.
	// The simulation will continue running batch jobs, including the specified number of requests,
	// according to the RunAfterDuration pattern, for the total number of times indicated by RunDuration.
	RunDuration int `json:"run_duration"`

	// Timeout specifies the maximum duration the operation is allowed to take,
	// formatted as a string, e.g., "10s" for 10 seconds.
	Timeout string `json:"timeout"`
}

// validate checks if the configuration is valid and returns an error if not
func (c *CustomReq) validate() error {
	if c.NumberOfRequests < c.Interval {
		return errors.New("number of requests must be more than intervals")
	}

	return nil
}

// getDuration gets the duration out of the custom struct
func (c *CustomReq) getDuration() (time.Duration, error) {
	duration, err := getDuration(c.Timeout)
	if err != nil {
		return 0, err
	}
	return duration, nil
}

type CustomFunction struct {
	// Method specifies the HTTP method for the request, e.g., "GET", "POST", etc.
	Method string `json:"method"`

	// URL is the URL of the endpoint to be called.
	URL string `json:"url"`

	// Body contains the request body as a byte slice.
	Body []byte `json:"body"`

	// Timeout specifies the maximum duration the request is allowed to take,
	// formatted as a string, e.g., "10s" for 10 seconds.
	Timeout string `json:"timeout"`

	// Headers is a slice of Header representing the custom headers to be included in the request.
	Headers []Header `json:"headers"`

	// Cookies is a slice of Cookie representing the custom cookies to be included in the request.
	Cookies []Cookie `json:"cookies"`
}

type Header struct {
	// Name is the name of the header.
	Name string `json:"name"`

	// Value is the value of the header.
	Value string `json:"value"`
}

// validate checks if the configuration is valid and returns an error if not.
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
	// Name is the name of the cookie.
	Name string `json:"name"`

	// Value is the value of the cookie.
	Value string `json:"value"`

	// ExpiresAt specifies the expiration time for the cookie as a string.
	// The format of the string depends on the time package's parsing rules,
	// and it represents the moment at which the cookie should expire.
	ExpiresAt string `json:"expires_at"`
}

// validate checks if the configuration is valid and returns an error if not
// also returns time.Duration that would be used in contexts for later.
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
