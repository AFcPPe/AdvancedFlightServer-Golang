package pdu

import "fmt"

type ControllerInfo struct {
	Base
	Payload []string
}

func NewPDUControllerInfo(from, to string, payload []string) *ControllerInfo {
	return &ControllerInfo{
		Base: Base{
			From: from,
			To:   to,
		},
		Payload: payload,
	}
}

func ControllerInfoFromTokens(tokens []string) (*ControllerInfo, error) {
	if len(tokens) < 2 {
		return nil, fmt.Errorf("PDUError tokens length < 2")
	}

	payload := make([]string, len(tokens)-2)
	for i := 2; i < len(tokens); i++ {
		payload[i-2] = tokens[i]
	}

	return NewPDUControllerInfo(tokens[0], tokens[1], payload), nil
}

func (pdu *ControllerInfo) ToTokens() []string {
	tokens := make([]string, len(pdu.Payload)+2)
	tokens[0] = pdu.From
	tokens[1] = pdu.To
	for i := 0; i < len(pdu.Payload); i++ {
		tokens[i+2] = pdu.Payload[i]
	}
	return tokens
}

func (pdu *ControllerInfo) GetHeader() string {
	return "#PC"
}
