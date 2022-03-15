package showmessage

type MessageType uint8

const (
	Error   MessageType = 1
	Warning MessageType = 2
	Info    MessageType = 3
	log     MessageType = 4
)

type Parameters struct {
	/*
		The message type. See {@link MessageType}
	*/
	MessageType *MessageType `json:"type,omitempty"`

	/*
		The actual message
	*/
	Message *string  `json:"message,omitempty"`
}
