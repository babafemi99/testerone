package yaml

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/babafemi99/testerone/load"
	"gopkg.in/yaml.v2"
)

type CustomFunction struct {
	Method string                 `yaml:"method"`
	URL    string                 `yaml:"url"`
	Body   map[string]interface{} `yaml:"body"`
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

func LoadYAMLFile(filePath string) (load.CustomReq, error) {
	var req CustomReq

	yamlContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return load.CustomReq{}, err
	}

	err = yaml.Unmarshal(yamlContent, &req)
	if err != nil {
		return load.CustomReq{}, err
	}

	var func2Custom []load.CustomFunction
	for _, cf := range req.Func2 {

		bodyJSON, err := json.Marshal(cf.Body)
		if err != nil {
			return load.CustomReq{}, err
		}

		// Convert bodyJSON to byte slice
		bodyBytes := []byte(bodyJSON)

		func2Custom = append(func2Custom, load.CustomFunction{
			Method: cf.Method,
			URL:    cf.URL,
			Body:   bodyBytes,
		})
	}

	loadReq := load.CustomReq{
		ReqType:          req.ReqType,
		NumberOfRequests: req.NumberOfRequests,
		URL:              req.URL,
		Interval:         req.Interval,
		Func2:            func2Custom,
		RunAfterDuration: req.RunAfterDuration,
		RunDuration:      req.RunDuration,
	}

	return loadReq, nil
}
