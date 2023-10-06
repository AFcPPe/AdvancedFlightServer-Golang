package pdu

type NetworkError int

const (
	Ok NetworkError = iota
	CallsignInUse
	CallsignInvalid
	AlreadyRegistered
	SyntaxError
	PDUSourceInvalid
	InvalidLogon
	NoSuchCallsign
	NoFlightPlan
	NoWeatherProfile
	InvalidProtocolRevision
	RequestedLevelTooHigh
	ServerFull
	CertificateSuspended
	InvalidControl
	InvalidPositionForRating
	UnauthorizedSoftware
)

func (ne NetworkError) String() string {
	switch ne {
	case Ok:
		return "Ok"
	case CallsignInUse:
		return "CallsignInUse"
	case CallsignInvalid:
		return "CallsignInvalid"
	case AlreadyRegistered:
		return "AlreadyRegistered"
	case SyntaxError:
		return "SyntaxError"
	case PDUSourceInvalid:
		return "PDUSourceInvalid"
	case InvalidLogon:
		return "InvalidLogon"
	case NoSuchCallsign:
		return "NoSuchCallsign"
	case NoFlightPlan:
		return "NoFlightPlan"
	case NoWeatherProfile:
		return "NoWeatherProfile"
	case InvalidProtocolRevision:
		return "InvalidProtocolRevision"
	case RequestedLevelTooHigh:
		return "RequestedLevelTooHigh"
	case ServerFull:
		return "ServerFull"
	case CertificateSuspended:
		return "CertificateSuspended"
	case InvalidControl:
		return "InvalidControl"
	case InvalidPositionForRating:
		return "InvalidPositionForRating"
	case UnauthorizedSoftware:
		return "UnauthorizedSoftware"
	default:
		return ""
	}
}
