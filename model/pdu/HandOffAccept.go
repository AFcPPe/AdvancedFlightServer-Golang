package pdu

import "fmt"

type HandOffAccept struct {
	Base
	Target string
}

func NewPDUHandOffAccept(from, to, target string) *HandOffAccept {
	return &HandOffAccept{
		Base: Base{
			From: from,
			To:   to,
		},
		Target: target,
	}
}

func HandOffAcceptFromTokens(tokens []string) (*HandOffAccept, error) {
	if len(tokens) < 3 {
		return nil, fmt.Errorf("Invalid number of tokens")
	}

	return NewPDUHandOffAccept(tokens[0], tokens[1], tokens[2]), nil
}

func (p *HandOffAccept) ToTokens() []string {
	return []string{p.From, p.To, p.Target}
}

func (p *HandOffAccept) GetHeader() string {
	return "$HA"
}
