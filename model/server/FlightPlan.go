package server

type FlightPlan struct {
	FlightRules   FlightRule
	Type          string
	TAS           string
	Dep           string
	DepTime       string
	ActualDepTime string
	CruiseAlt     string
	Dest          string
	EnrouteHour   string
	EnrouteMin    string
	FobHour       string
	FobMin        string
	AlterDest     string
	Remark        string
	Route         string
	Locked        bool
}

type FlightRule int

const (
	IFR FlightRule = iota
	VFR
)
