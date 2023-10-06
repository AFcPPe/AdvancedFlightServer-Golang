package pdu

import "fmt"

type TextMessage struct {
	Base
	Message string
}

func NewPDUTextMessage(from, to, message string) *TextMessage {
	return &TextMessage{
		Base: Base{
			From: from,
			To:   to,
		},
		Message: message,
	}
}

func TextMessageFromTokens(tokens []string) (*TextMessage, error) {
	if len(tokens) < 3 {
		return nil, fmt.Errorf("PDUError tokens length < 3")
	}
	message := tokens[2]
	for i := 3; i < len(tokens); i++ {
		message += ":" + tokens[i]
	}
	return NewPDUTextMessage(tokens[0], tokens[1], message), nil
}

func (pdu *TextMessage) ToTokens() []string {
	return []string{pdu.From, pdu.To, pdu.Message}
}

func (pdu *TextMessage) GetHeader() string {
	return "#TM"
}
