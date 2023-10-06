package server

type Context struct {
	IncomingPacket   string
	Authorized       bool
	Callsign         string
	Cid              string
	Type             ClientType
	Lat              float64
	Lon              float64
	Range            int
	Rating           NetworkRating
	Facility         NetworkFacility
	Pitch            float64
	Bank             float64
	Heading          float64
	FlightPlan       FlightPlan
	Frequencies      []string
	SquawkCode       int
	SquawkingModeC   bool
	Identing         bool
	TrueAltitude     int
	PressureAltitude int
	GroundSpeed      int
}
