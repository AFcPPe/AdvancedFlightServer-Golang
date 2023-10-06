package pdu

import (
	"AdvancedFlightServer/model/server"
	"fmt"
	"strconv"
)

type PilotPosition struct {
	Base
	SquawkCode       int
	SquawkingModeC   bool
	Identing         bool
	Rating           server.NetworkRating
	Lat              float64
	Lon              float64
	TrueAltitude     int
	PressureAltitude int
	GroundSpeed      int
	Pitch            float64
	Heading          float64
	Bank             float64
}

func NewPDUPilotPosition(from string, txCode int, squawkingModeC bool, identing bool, rating int, lat float64, lon float64, trueAlt int, pressureAlt int, gs int, pitch float64, heading float64, bank float64) *PilotPosition {
	return &PilotPosition{
		Base: Base{
			From: from,
			To:   "",
		},
		SquawkCode:       txCode,
		SquawkingModeC:   squawkingModeC,
		Identing:         identing,
		Rating:           server.NetworkRating(rating),
		Lat:              lat,
		Lon:              lon,
		TrueAltitude:     trueAlt,
		PressureAltitude: pressureAlt,
		GroundSpeed:      gs,
		Pitch:            pitch,
		Heading:          heading,
		Bank:             bank,
	}
}

func PilotPositionFromTokens(tokens []string) (*PilotPosition, error) {
	if len(tokens) < 10 {
		return nil, fmt.Errorf("PDUError tokens length < 10")
	}
	pbh, _ := strconv.ParseInt(tokens[8], 10, 64)
	pbhValue := UnpackPitchBankHeading(int64(pbh))
	txCode, err := strconv.Atoi(tokens[2])
	if err != nil {
		return nil, err
	}
	rating, err := strconv.Atoi(tokens[3])
	if err != nil {
		return nil, err
	}
	trueAlt, err := strconv.Atoi(tokens[6])
	if err != nil {
		return nil, err
	}
	pressureAlt, err := strconv.Atoi(tokens[6])
	if err != nil {
		return nil, err
	}
	deltaAlt, _ := strconv.Atoi(tokens[9])
	pressureAlt += deltaAlt
	if err != nil {
		return nil, err
	}
	gs, err := strconv.Atoi(tokens[7])
	if err != nil {
		return nil, err
	}
	pitch := pbhValue[0]
	heading := pbhValue[2]
	bank := pbhValue[1]
	lat, _ := strconv.ParseFloat(tokens[4], 64)
	lon, _ := strconv.ParseFloat(tokens[5], 64)
	return NewPDUPilotPosition(tokens[1], txCode, tokens[0] == "N", tokens[0] == "I", rating,
		lat, lon, trueAlt, pressureAlt, gs, pitch, heading, bank), nil
}

func (p *PilotPosition) ToTokens() []string {
	identing := "S"
	if p.Identing {
		identing = "I"
	} else if p.SquawkingModeC {
		identing = "N"
	}

	return []string{
		identing,
		p.From,
		strconv.Itoa(p.SquawkCode),
		strconv.Itoa(int(p.Rating)),
		strconv.FormatFloat(p.Lat, 'f', -1, 64),
		strconv.FormatFloat(p.Lon, 'f', -1, 64),
		strconv.Itoa(p.TrueAltitude),
		strconv.Itoa(p.GroundSpeed),
		strconv.FormatInt(PackPitchBankHeading(p.Pitch, p.Bank, p.Heading), 10),
		strconv.Itoa(p.PressureAltitude - p.TrueAltitude),
	}
}

func (p *PilotPosition) GetHeader() string {
	return "@"
}
