package kkpm

// MessageInfo for private message.
type MessageInfo struct {
	MessageID string `json:"messageid"`
	Message   string `json:"message"`
	FromUser  int    `json:"fromuser,omitempty"`
	ToUser    int    `json:"touser,omitempty"`
	At        int    `json:"at"`
}

// SendMessage to send a message.
func SendMessage(from, to int, message string) {

}
