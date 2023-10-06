package pdu

import "fmt"

type PlaneInfo struct {
	Base
	Type    string
	Payload []string
}

func NewPDUPlaneInfo(from, to, planeType string, payload []string) *PlaneInfo {
	return &PlaneInfo{
		Base: Base{
			From: from,
			To:   to,
		},
		Type:    planeType,
		Payload: payload,
	}
}

func PlaneInfoFromTokens(tokens []string) (*PlaneInfo, error) {
	if len(tokens) < 3 {
		return nil, fmt.Errorf("PDUError tokens length < 3")
	}

	payload := make([]string, len(tokens)-3)
	for i := 3; i < len(tokens); i++ {
		payload[i-3] = tokens[i]
	}

	return NewPDUPlaneInfo(tokens[0], tokens[1], tokens[2], payload), nil
}

func (pdu *PlaneInfo) ToTokens() []string {
	tokens := make([]string, len(pdu.Payload)+3)
	tokens[0] = pdu.From
	tokens[1] = pdu.To
	tokens[2] = pdu.Type
	for i := 0; i < len(pdu.Payload); i++ {
		tokens[i+3] = pdu.Payload[i]
	}
	return tokens
}

func (pdu *PlaneInfo) GetHeader() string {
	return "#SB"
}
