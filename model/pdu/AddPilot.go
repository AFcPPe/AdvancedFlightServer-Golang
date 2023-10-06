package pdu

import (
	"AdvancedFlightServer/model/server"
	"fmt"
	"strconv"
)

type AddPilot struct {
	Base
	Callsign        string
	Cid             string
	Password        string
	Rating          server.NetworkRating
	ProtocolVersion int
	RealName        string
	SimType         string
}

func NewPDUAddPilot(from, to, cid, password string, networkRating, protocolVersion int, simType int, realName string) *AddPilot {
	return &AddPilot{
		Base: Base{
			From: from,
			To:   to,
		},
		Callsign:        from,
		Cid:             cid,
		Password:        password,
		Rating:          server.NetworkRating(networkRating),
		ProtocolVersion: protocolVersion,
		RealName:        realName,
		SimType:         strconv.Itoa(simType),
	}
}

func AddPilotFromTokens(tokens []string) (*AddPilot, error) {
	if len(tokens) < 8 {
		return nil, fmt.Errorf("PDUError tokens length < 8")
	}
	networkRating, err := strconv.Atoi(tokens[4])
	if err != nil {
		return nil, err
	}
	protocolVersion, err := strconv.Atoi(tokens[5])
	if err != nil {
		return nil, err
	}
	simType, err := strconv.Atoi(tokens[6])
	if err != nil {
		return nil, err
	}
	return NewPDUAddPilot(tokens[0], tokens[1], tokens[2], tokens[3], networkRating, protocolVersion, simType, tokens[7]), nil
}

func (pdu *AddPilot) ToTokens() []string {
	return []string{
		pdu.From,
		pdu.To,
		pdu.Cid,
		pdu.Password,
		strconv.Itoa(int(pdu.Rating)),
		strconv.Itoa(pdu.ProtocolVersion),
		pdu.SimType,
		pdu.RealName,
	}
}

func (pdu *AddPilot) GetHeader() string {
	return "#AP"
}
