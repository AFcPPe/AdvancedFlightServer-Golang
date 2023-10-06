package pdu

import "fmt"

type MetarResponse struct {
	Base
	Type  string
	Metar string
}

func NewPDUMetarResponse(from, to, typeVal, metar string) *MetarResponse {
	return &MetarResponse{
		Base: Base{
			From: from,
			To:   to,
		},
		Type:  typeVal,
		Metar: metar,
	}
}

func (p *MetarResponse) FromTokens(tokens []string) (*MetarResponse, error) {
	if len(tokens) < 4 {
		return nil, fmt.Errorf("")
	}

	return NewPDUMetarResponse(tokens[0], tokens[1], tokens[2], tokens[3]), nil
}

func (p *MetarResponse) ToTokens() []string {
	return []string{p.From, p.To, p.Type, p.Metar}
}

func (p *MetarResponse) GetHeader() string {
	return "$AR"
}
