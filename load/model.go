package load

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
	NumberOfRequests int    `json:"number_of_requests"`
	URL              string `json:"url"`
}
