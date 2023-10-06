package pdu

import (
	"AdvancedFlightServer/util"
	"strings"
)

type Base struct {
	From string
	To   string
}

func NewPDUBase(from, to string) *Base {
	return &Base{
		From: from,
		To:   to,
	}
}

func Serialize(header string, tokens []string) string {
	payload := strings.Join(tokens, util.PayloadDelimiter)
	return header + payload + util.PacketDelimiter
}

func UnpackPitchBankHeading(pbh int64) []float64 {
	pitchInt := pbh >> 22

	pitch := float64(pitchInt) / 1024.0 * -360.0
	if pitch > 180.0 {
		pitch -= 360.0
	} else if pitch <= -180.0 {
		pitch += 360.0
	}

	bankInt := (pbh >> 12) & 0x3FF
	bank := float64(bankInt) / 1024.0 * -360.0
	if bank > 180.0 {
		bank -= 360.0
	} else if bank <= -180.0 {
		bank += 360.0
	}

	hdgInt := (pbh >> 2) & 0x3FF
	heading := float64(hdgInt) / 1024.0 * 360.0
	if heading < 0.0 {
		heading += 360.0
	} else if heading >= 360.0 {
		heading -= 360.0
	}

	return []float64{pitch, bank, heading}
}

func PackPitchBankHeading(pitch, bank, heading float64) int64 {
	p := pitch / -360.0
	if p < 0 {
		p += 1.0
	}
	p *= 1024.0

	b := bank / -360.0
	if b < 0 {
		b += 1.0
	}
	b *= 1024.0

	h := heading / 360.0 * 1024.0

	return int64(p)<<22 | int64(b)<<12 | int64(h)<<2
}
