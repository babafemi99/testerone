package yaml

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/babafemi99/testerone/load"
	"gopkg.in/yaml.v2"
)

type BasicReq struct {
	NumberOfRequests int           `yaml:"number_of_requests"`
	URL              string        `yaml:"url"`
	Interval         int           `yaml:"interval"`
	RunAfterDuration time.Duration `yaml:"run_after_duration"`
	RunDuration      int           `yaml:"run_duration"`
}

type CustomFunction struct {
	Method  string                 `yaml:"method"`
	URL     string                 `yaml:"url"`
	Body    map[string]interface{} `yaml:"body"`
	Timeout string                 `yaml:"timeout"`
	Headers []Header               `yaml:"headers"`
	Cookies []Cookie               `yaml:"cookies"`
}

type Header struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type Cookie struct {
	Name      string `yaml:"name"`
	Value     string `yaml:"value"`
	ExpiresAt string `yaml:"expires_at"`
}

type CustomReq struct {
	ReqType          string           `yaml:"req_type"`
	NumberOfRequests int              `yaml:"number_of_requests"`
	URL              string           `yaml:"url"`
	Interval         int              `yaml:"interval"`
	Func2            []CustomFunction `yaml:"func_2"`
	RunAfterDuration time.Duration    `yaml:"run_after_duration"`
	RunDuration      int              `yaml:"run_duration"`
}

type RequestEntry struct {
	ReqType       string     `yaml:"req_type"`
	CustomRequest *CustomReq `yaml:"custom_request"`
	BasicRequest  *BasicReq  `yaml:"basic_request"`
}

func LoadYAMLFile(filePath string) (*load.CustomReq, *load.Req, error) {
	var req []RequestEntry

	yamlContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	err = yaml.Unmarshal(yamlContent, &req)
	if err != nil {
		return nil, nil, err
	}

	if len(req) > 0 {
		entry := req[0]

		switch entry.ReqType {
		case "custom":

			if entry.CustomRequest != nil {
				customReq := *entry.CustomRequest
				var func2Custom []load.CustomFunction
				for _, cf := range customReq.Func2 {

					bodyJSON, err := json.Marshal(cf.Body)
					if err != nil {
						return nil, nil, err
					}

					// Convert bodyJSON to byte slice
					bodyBytes := []byte(bodyJSON)

					// Create headers and cookies slices
					var headers []load.Header
					for _, h := range cf.Headers {
						headers = append(headers, load.Header{
							Name:  h.Name,
							Value: h.Value,
						})
					}

					var cookies []load.Cookie
					for _, yc := range cf.Cookies {
						cookies = append(cookies, load.Cookie{
							Name:      yc.Name,
							Value:     yc.Value,
							ExpiresAt: yc.ExpiresAt,
						})
					}

					func2Custom = append(func2Custom, load.CustomFunction{
						Method:  cf.Method,
						URL:     cf.URL,
						Body:    bodyBytes,
						Headers: headers,
						Cookies: cookies,
					})
				}

				loadReq := &load.CustomReq{
					// ReqType:          customReq.ReqType,
					NumberOfRequests: customReq.NumberOfRequests,
					Interval:         customReq.Interval,
					Func2:            func2Custom,
					RunAfterDuration: customReq.RunAfterDuration,
					RunDuration:      customReq.RunDuration,
				}
				return loadReq, nil, nil
			}

		case "basic":
			if entry.BasicRequest != nil {
				basicReq := *entry.BasicRequest

				loadReq := &load.Req{
					NumberOfRequests: basicReq.NumberOfRequests,
					URL:              basicReq.URL,
					Interval:         basicReq.Interval,
					RunAfterDuration: basicReq.RunAfterDuration,
					RunDuration:      basicReq.RunDuration,
				}

				return nil, loadReq, nil
			}

		}
	}

	return nil, nil, err
}
