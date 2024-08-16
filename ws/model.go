package ws

type ChanWs struct {
	User         string `json:"user"`
	Message      string `json:"message"`
	MessageType  string `json:"message_type"`
	Error        error  `json:"error"`
	ResponseTime string `json:"response_time"`
}
