package pdu

import "fmt"

type DeletePilot struct {
	Base
	Cid string
}

func NewPDUDeletePilot(from, cid string) *DeletePilot {
	return &DeletePilot{
		Base: Base{
			From: from,
			To:   "*",
		},
		Cid: cid,
	}
}

func DeletePilotFromTokens(tokens []string) (*DeletePilot, error) {
	if len(tokens) < 2 {
		return nil, fmt.Errorf("")
	}

	return NewPDUDeletePilot(tokens[0], tokens[1]), nil
}

func (p *DeletePilot) ToTokens() []string {
	return []string{p.From, p.Cid}
}

func (p *DeletePilot) GetHeader() string {
	return "#DP"
}
