package pdu

import "fmt"

type ClientQuery struct {
	Base
	Type    string
	Payload []string
}

func NewPDUClientQuery(from, to, typeStr string, payload []string) *ClientQuery {
	return &ClientQuery{
		Base: Base{
			From: from,
			To:   to,
		},
		Type:    typeStr,
		Payload: payload,
	}
}

func ClientQueryFromTokens(tokens []string) (*ClientQuery, error) {
	if len(tokens) < 3 {
		return nil, fmt.Errorf("PDUError tokens length < 3")
	}

	payload := make([]string, len(tokens)-3)
	for i := 3; i < len(tokens); i++ {
		payload[i-3] = tokens[i]
	}

	return NewPDUClientQuery(tokens[0], tokens[1], tokens[2], payload), nil
}

func (p *ClientQuery) ToTokens() []string {
	tokens := make([]string, len(p.Payload)+3)
	tokens[0] = p.From
	tokens[1] = p.To
	tokens[2] = p.Type
	for i := 0; i < len(p.Payload); i++ {
		tokens[i+3] = p.Payload[i]
	}
	return tokens
}

func (p *ClientQuery) GetHeader() string {
	return "$CQ"
}
