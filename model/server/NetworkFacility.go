package server

type NetworkFacility int

const (
	Observer NetworkFacility = iota
	FSS
	DEL
	GND
	TWR
	APP
	CTR
)

func (nf NetworkFacility) String() string {
	switch nf {
	case Observer:
		return "OBS"
	case FSS:
		return "FSS"
	case DEL:
		return "DEL"
	case GND:
		return "GND"
	case TWR:
		return "TWR"
	case APP:
		return "APP"
	case CTR:
		return "CTR"
	default:
		return ""
	}
}
