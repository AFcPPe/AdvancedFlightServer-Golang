package server

type NetworkRating int

const (
	SUS NetworkRating = iota
	OBS
	S1
	S2
	S3
	C1
	C2
	C3
	I1
	I2
	I3
	SUP
	ADM
)

func (nr NetworkRating) String() string {
	switch nr {
	case SUS:
		return "SUS"
	case OBS:
		return "OBS"
	case S1:
		return "S1"
	case S2:
		return "S2"
	case S3:
		return "S3"
	case C1:
		return "C1"
	case C2:
		return "C2"
	case C3:
		return "C3"
	case I1:
		return "I1"
	case I2:
		return "I2"
	case I3:
		return "I3"
	case SUP:
		return "SUP"
	case ADM:
		return "ADM"
	default:
		return ""
	}
}
