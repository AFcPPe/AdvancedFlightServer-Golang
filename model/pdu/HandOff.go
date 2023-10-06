package pdu

import "fmt"

type HandOff struct {
	Base
	Target string
}

func NewPDUHandOff(from, to, target string) *HandOff {
	return &HandOff{
		Base: Base{
			From: from,
			To:   to,
		},
		Target: target,
	}
}

func HandOffFromTokens(tokens []string) (*HandOff, error) {
	if len(tokens) < 3 {
		return nil, fmt.Errorf("Invalid number of tokens")
	}

	return NewPDUHandOff(tokens[0], tokens[1], tokens[2]), nil
}

func (p *HandOff) ToTokens() []string {
	return []string{p.From, p.To, p.Target}
}

func (p *HandOff) GetHeader() string {
	return "$HO"
}
