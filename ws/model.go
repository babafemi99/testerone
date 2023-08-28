package ws

type ChanWs struct {
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
	Error       error  `json:"error"`
}
