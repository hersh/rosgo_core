package ros

import (
	"bytes"
)

type MessageType interface {
	Text() string
	MD5Sum() string
	Name() string
	NewMessage() Message
}

type Message interface {
	Serialize(buf *bytes.Buffer) error
	Deserialize(buf *bytes.Reader) error
}
