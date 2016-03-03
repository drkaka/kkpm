package kkpm

// MessageInfo for private message.
type MessageInfo struct {
	MessageID string `json:"messageid"`
	Message   string `json:"message"`
	FromUser  int32  `json:"fromuser,omitempty"`
	ToUser    int32  `json:"touser,omitempty"`
	At        int32  `json:"at"`
}

// SendMessage to send a message.
func SendMessage(from, to int32, message string) {

}
