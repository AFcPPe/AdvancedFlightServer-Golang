package pdu

import "fmt"

type ClientQueryResponse struct {
	Base
	Type    string
	Payload []string
}

func NewPDUClientQueryResponse(from, to, typeStr string, payload []string) *ClientQueryResponse {
	return &ClientQueryResponse{
		Base: Base{
			From: from,
			To:   to,
		},
		Type:    typeStr,
		Payload: payload,
	}
}

func ClientQueryResponseFromTokens(tokens []string) (*ClientQueryResponse, error) {
	if len(tokens) < 3 {
		return nil, fmt.Errorf("PDUError tokens length < 3")
	}

	payload := make([]string, len(tokens)-3)
	for i := 3; i < len(tokens); i++ {
		payload[i-3] = tokens[i]
	}

	return NewPDUClientQueryResponse(tokens[0], tokens[1], tokens[2], payload), nil
}

func (p *ClientQueryResponse) ToTokens() []string {
	tokens := make([]string, len(p.Payload)+3)
	tokens[0] = p.From
	tokens[1] = p.To
	tokens[2] = p.Type
	for i := 0; i < len(p.Payload); i++ {
		tokens[i+3] = p.Payload[i]
	}
	return tokens
}

func (p *ClientQueryResponse) GetHeader() string {
	return "$CR"
}
