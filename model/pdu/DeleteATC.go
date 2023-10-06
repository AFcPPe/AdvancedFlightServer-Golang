package pdu

import "fmt"

type DeleteATC struct {
	Base
	Cid string
}

func NewPDUDeleteATC(from, cid string) *DeleteATC {
	return &DeleteATC{
		Base: Base{
			From: from,
			To:   "*",
		},
		Cid: cid,
	}
}

func DeleteATCFromTokens(tokens []string) (*DeleteATC, error) {
	if len(tokens) < 2 {
		return nil, fmt.Errorf("")
	}

	return NewPDUDeleteATC(tokens[0], tokens[1]), nil
}

func (p *DeleteATC) ToTokens() []string {
	return []string{p.From, p.To, p.Cid}
}

func (p *DeleteATC) GetHeader() string {
	return "#DA"
}
