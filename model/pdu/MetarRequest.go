package pdu

import "fmt"

type MetarRequest struct {
	Base
	Type string
	ICAO string
}

func NewPDUMetarRequest(from, to, typeVal, icao string) *MetarRequest {
	return &MetarRequest{
		Base: Base{
			From: from,
			To:   to,
		},
		Type: typeVal,
		ICAO: icao,
	}
}

func MetarRequestFromTokens(tokens []string) (*MetarRequest, error) {
	if len(tokens) < 4 {
		return nil, fmt.Errorf("")
	}

	return NewPDUMetarRequest(tokens[0], tokens[1], tokens[2], tokens[3]), nil
}

func (p *MetarRequest) ToTokens() []string {
	return []string{p.From, p.To, p.Type, p.ICAO}
}

func (p *MetarRequest) GetHeader() string {
	return "$AX"
}
