package pdu

import (
	"fmt"
	"strconv"
)

type Error struct {
	Base
	Error NetworkError
	Msg   string
	Param string
}

func NewPDUError(from, to string, errorVal int, param, msg string) *Error {
	return &Error{
		Base: Base{
			From: from,
			To:   to,
		},
		Error: NetworkError(errorVal),
		Msg:   msg,
		Param: param,
	}
}

func ErrorFromTokens(tokens []string) (*Error, error) {
	if len(tokens) < 5 {
		return nil, fmt.Errorf("invalid number of tokens")
	}

	errorVal, err := strconv.Atoi(tokens[2])
	if err != nil {
		return nil, fmt.Errorf("invalid error value")
	}

	return NewPDUError(tokens[0], tokens[1], errorVal, tokens[3], tokens[4]), nil
}

func (p *Error) ToTokens() []string {
	return []string{p.From, p.To, strconv.Itoa(int(p.Error)), p.Param, p.Msg}
}

func (p *Error) GetHeader() string {
	return "$ER"
}
