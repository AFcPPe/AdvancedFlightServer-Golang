package service

import (
	"AdvancedFlightServer/model/orm"
	"AdvancedFlightServer/model/pdu"
	"AdvancedFlightServer/model/server"
	"AdvancedFlightServer/util"
	"github.com/panjf2000/gnet/v2"
	"github.com/spf13/viper"
	"log"
	"math"
	"strings"
)

func ServerHandler(packet []string, c gnet.Conn) {
	for _, s := range packet {
		if len(s) <= 3 {
			continue
		}
		header := s[:3]
		tokens := strings.Split(s, util.PayloadDelimiter)
		switch s[0] {
		case '@':
			tokens[0] = tokens[0][1:]
			aPdu, err := pdu.PilotPositionFromTokens(tokens)
			if err != nil {
				log.Println(err)
				continue
			}
			onPilotPositionReceived(aPdu, c)
			continue
		case '%':
			tokens[0] = tokens[0][1:]
			aPdu, err := pdu.ATCPositionFromTokens(tokens)
			if err != nil {
				log.Println(err)
				continue
			}
			onATCPositionReceived(aPdu, c)
			continue
		case '#':
			tokens[0] = tokens[0][3:]
			if header == "#AA" {
				aPdu, err := pdu.AddATCFromTokens(tokens)
				if err != nil {
					log.Println(err)
					continue
				}
				success := onAddATCReceived(aPdu, c)
				if !success {
					err := c.Close()
					if err != nil {
						return
					}
				}
				continue
			}
			if header == "#AP" {
				aPdu, err := pdu.AddPilotFromTokens(tokens)
				if err != nil {
					log.Println(err)
					continue
				}
				success := onAddPilotReceived(aPdu, c)
				if !success {
					err := c.Close()
					if err != nil {
						return
					}
				}
				continue
			}
			if header == "#DP" {
				aPdu, err := pdu.DeletePilotFromTokens(tokens)
				if err != nil {
					log.Println(err)
					continue
				}
				onDeletePilotReceived(aPdu, c)
				continue
			}
			if header == "#DA" {
				aPdu, err := pdu.DeleteATCFromTokens(tokens)
				if err != nil {
					log.Println(err)
					continue
				}
				onDeleteATCReceived(aPdu, c)
				continue
			}
			if header == "#TM" {
				aPdu, err := pdu.TextMessageFromTokens(tokens)
				if err != nil {
					log.Println(err)
					continue
				}
				onTextMessageReceived(aPdu, c)
				continue
			}
			if header == "#SB" {
				aPdu, err := pdu.PlaneInfoFromTokens(tokens)
				if err != nil {
					log.Println(err)
					continue
				}
				onPlaneInfoReceived(aPdu, c)
				continue
			}
			if header == "#PC" {
				aPdu, err := pdu.ControllerInfoFromTokens(tokens)
				if err != nil {
					log.Println(err)
					continue
				}
				onControllerInfoReceived(aPdu, c)
				continue
			}
			break
		case '$':
			tokens[0] = tokens[0][3:]
			if header == "$CQ" {
				aPdu, err := pdu.ClientQueryFromTokens(tokens)
				if err != nil {
					log.Println(err)
					continue
				}
				onClientQueryReceived(aPdu, c)
				continue
			}
			if header == "$CR" {
				aPdu, err := pdu.ClientQueryResponseFromTokens(tokens)
				if err != nil {
					log.Println(err)
					continue
				}
				onClientQueryResponseReceived(aPdu, c)
				continue
			}
			if header == "$AX" {
				aPdu, err := pdu.MetarRequestFromTokens(tokens)
				if err != nil {
					log.Println(err)
					continue
				}
				onMetarRequestReceived(aPdu, c)
				continue
			}
			if header == "$FP" {
				aPdu, err := pdu.FlightPlanFromTokens(tokens)
				if err != nil {
					log.Println(err)
					continue
				}
				onFlightPlanReceived(aPdu, c)
				continue
			}
			if header == "$HO" {
				aPdu, err := pdu.HandOffFromTokens(tokens)
				if err != nil {
					log.Println(err)
					continue
				}
				onHandOffReceived(aPdu, c)
				continue
			}
			if header == "$HA" {
				aPdu, err := pdu.HandOffAcceptFromTokens(tokens)
				if err != nil {
					log.Println(err)
					continue
				}
				onHandOffAcceptReceived(aPdu, c)
				continue
			}
			break
		}
	}
}

func onHandOffAcceptReceived(aPdu *pdu.HandOffAccept, c gnet.Conn) {
	ctx := GetConnContext(c)
	if !ctx.Authorized {
		return
	}
	if ctx.Callsign != aPdu.From {
		return
	}
	sendMessageToTarget(aPdu.To, pdu.Serialize(aPdu.GetHeader(), aPdu.ToTokens()))
}

func onHandOffReceived(aPdu *pdu.HandOff, c gnet.Conn) {
	ctx := GetConnContext(c)
	if !ctx.Authorized {
		return
	}
	if ctx.Callsign != aPdu.From {
		return
	}
	sendMessageToTarget(aPdu.To, pdu.Serialize(aPdu.GetHeader(), aPdu.ToTokens()))
}

func onControllerInfoReceived(aPdu *pdu.ControllerInfo, c gnet.Conn) {
	ctx := GetConnContext(c)
	if !ctx.Authorized {
		return
	}
	if ctx.Callsign != aPdu.From {
		return
	}
	sendMessageToTarget(aPdu.To, pdu.Serialize(aPdu.GetHeader(), aPdu.ToTokens()))
}

func onPlaneInfoReceived(aPdu *pdu.PlaneInfo, c gnet.Conn) {
	ctx := GetConnContext(c)
	if !ctx.Authorized {
		return
	}
	if ctx.Callsign != aPdu.From {
		return
	}
	sendMessageToTarget(aPdu.To, pdu.Serialize(aPdu.GetHeader(), aPdu.ToTokens()))
}

func onAddATCReceived(aPdu *pdu.AddATC, c gnet.Conn) bool {
	user := orm.User{}
	result := util.Db.Where("username = ?", aPdu.Cid).First(&user)
	if len(aPdu.Callsign) > 12 {
		sendErrorMessage("unknown", pdu.CallsignInvalid, aPdu.Cid, "Invalid callsign", c)
		return false
	}
	if result.Error != nil {
		sendErrorMessage("unknown", pdu.InvalidLogon, aPdu.Cid, "Wrong password or username", c)
		return false
	}
	if user.Password != aPdu.Password {
		sendErrorMessage("unknown", pdu.InvalidLogon, aPdu.Cid, "Wrong password or username", c)
		return false
	}
	if user.Rating < int(aPdu.Rating) {
		sendErrorMessage("unknown", pdu.RequestedLevelTooHigh, aPdu.Cid, "Requested level is too high", c)
		return false
	}
	if checkSameCallsign(aPdu.Callsign) {
		sendErrorMessage("unknown", pdu.CallsignInUse, aPdu.Cid, "Callsign in use", c)
		return false
	}
	ctx := GetConnContext(c)
	ctx.Cid = aPdu.Cid
	ctx.Callsign = aPdu.Callsign
	ctx.Rating = aPdu.Rating
	ctx.Authorized = true
	ctx.Type = server.ATC
	c.SetContext(ctx)
	readMotd(c)
	return true
}

func onAddPilotReceived(aPdu *pdu.AddPilot, c gnet.Conn) bool {
	user := orm.User{}
	result := util.Db.Where("username = ?", aPdu.Cid).First(&user)
	if len(aPdu.Callsign) > 12 {
		sendErrorMessage("unknown", pdu.CallsignInvalid, aPdu.Cid, "Invalid callsign", c)
		return false
	}
	if result.Error != nil {
		sendErrorMessage("unknown", pdu.InvalidLogon, aPdu.Cid, "Wrong password or username", c)
		return false
	}
	if user.Password != aPdu.Password {
		sendErrorMessage("unknown", pdu.InvalidLogon, aPdu.Cid, "Wrong password or username", c)
		return false
	}
	if user.Rating < int(aPdu.Rating) {
		sendErrorMessage("unknown", pdu.RequestedLevelTooHigh, aPdu.Cid, "Requested level is too high", c)
		return false
	}
	if checkSameCallsign(aPdu.Callsign) {
		sendErrorMessage("unknown", pdu.CallsignInUse, aPdu.Cid, "Callsign in use", c)
		return false
	}
	ctx := GetConnContext(c)
	ctx.Cid = aPdu.Cid
	ctx.Callsign = strings.ToUpper(aPdu.Callsign)
	ctx.Rating = aPdu.Rating
	ctx.Authorized = true
	ctx.Type = server.Pilot
	c.SetContext(ctx)
	readMotd(c)
	return true
}

func onATCPositionReceived(aPdu *pdu.ATCPosition, c gnet.Conn) {
	ctx := GetConnContext(c)
	if !ctx.Authorized {
		return
	}
	if aPdu.From != ctx.Callsign {
		return
	}
	ctx.Lat = aPdu.Lat
	ctx.Lon = aPdu.Lon
	ctx.Facility = aPdu.Facility
	ctx.Frequencies = aPdu.Frequencies
	ctx.Range = aPdu.VisibilityRange
	sendMessageToInRange(c, pdu.Serialize(aPdu.GetHeader(), aPdu.ToTokens()))
	c.SetContext(ctx)
}

func onPilotPositionReceived(aPdu *pdu.PilotPosition, c gnet.Conn) {
	ctx := GetConnContext(c)
	if !ctx.Authorized {
		return
	}
	if aPdu.From != ctx.Callsign {
		return
	}
	ctx.Lat = aPdu.Lat
	ctx.Lon = aPdu.Lon
	ctx.SquawkCode = aPdu.SquawkCode
	ctx.SquawkingModeC = aPdu.SquawkingModeC
	ctx.Identing = aPdu.Identing
	ctx.TrueAltitude = aPdu.TrueAltitude
	ctx.PressureAltitude = aPdu.PressureAltitude
	ctx.GroundSpeed = aPdu.GroundSpeed
	ctx.Heading = aPdu.Heading
	ctx.Pitch = aPdu.Pitch
	ctx.Bank = aPdu.Bank
	sendMessageToInRange(c, pdu.Serialize(aPdu.GetHeader(), aPdu.ToTokens()))
	c.SetContext(ctx)
}

func onDeletePilotReceived(aPdu *pdu.DeletePilot, c gnet.Conn) {
	ctx := GetConnContext(c)
	if !ctx.Authorized {
		return
	}
	if aPdu.From != ctx.Callsign {
		return
	}
	sendMessageToAll(c, pdu.Serialize(aPdu.GetHeader(), aPdu.ToTokens()))
}

func onDeleteATCReceived(aPdu *pdu.DeleteATC, c gnet.Conn) {
	ctx := GetConnContext(c)
	if !ctx.Authorized {
		return
	}
	if aPdu.From != ctx.Callsign {
		return
	}
	sendMessageToAll(c, pdu.Serialize(aPdu.GetHeader(), aPdu.ToTokens()))
}

func onTextMessageReceived(aPdu *pdu.TextMessage, c gnet.Conn) {
	ctx := GetConnContext(c)
	if !ctx.Authorized {
		return
	}
	if ctx.Callsign != aPdu.From {
		return
	}
	if strings.ToUpper(aPdu.To) == "*S" {
		sendMessageToSupervisor(c, pdu.Serialize(aPdu.GetHeader(), aPdu.ToTokens()))
		return
	}
	if strings.ToUpper(aPdu.To) == "*" {
		if ctx.Rating < server.SUP {
			return
		}
		sendMessageToAll(c, pdu.Serialize(aPdu.GetHeader(), aPdu.ToTokens()))
		return
	}
	if strings.ToUpper(aPdu.To) == "@" {
		sendMessageToAll(c, pdu.Serialize(aPdu.GetHeader(), aPdu.ToTokens()))
		return
	}
	if strings.ToUpper(aPdu.To) == "*A" {
		sendMessageToATC(c, pdu.Serialize(aPdu.GetHeader(), aPdu.ToTokens()))
		return
	}

	if strings.ToUpper(aPdu.To) == "FP" {
		return
	}
}

func onClientQueryReceived(aPdu *pdu.ClientQuery, c gnet.Conn) {
	ctx := GetConnContext(c)
	if !ctx.Authorized {
		return
	}
	if aPdu.From != ctx.Callsign {
		return
	}
	if strings.ToLower(aPdu.To) == "server" {
		dealServerQuery(aPdu, c)
		return
	}
	if aPdu.To[0] == '@' {
		sendMessageToAll(c, pdu.Serialize(aPdu.GetHeader(), aPdu.ToTokens()))
		return
	}
	sendMessageToTarget(aPdu.To, pdu.Serialize(aPdu.GetHeader(), aPdu.ToTokens()))
}

func onClientQueryResponseReceived(aPdu *pdu.ClientQueryResponse, c gnet.Conn) {
	ctx := GetConnContext(c)
	if !ctx.Authorized {
		return
	}
	if aPdu.From != ctx.Callsign {
		return
	}
	if aPdu.To == "server" {
		return
	}
	if aPdu.To[0] == '@' {
		sendMessageToAll(c, pdu.Serialize(aPdu.GetHeader(), aPdu.ToTokens()))
		return
	}
	sendMessageToTarget(aPdu.To, pdu.Serialize(aPdu.GetHeader(), aPdu.ToTokens()))
}

func onFlightPlanReceived(aPdu *pdu.FlightPlan, c gnet.Conn) {
	ctx := GetConnContext(c)
	if !ctx.Authorized {
		return
	}
	if ctx.Callsign != aPdu.From {
		return
	}
	ctx.FlightPlan.FlightRules = aPdu.FlightRules
	ctx.FlightPlan.Type = aPdu.Type
	ctx.FlightPlan.TAS = aPdu.TAS
	ctx.FlightPlan.Dep = aPdu.Dep
	ctx.FlightPlan.DepTime = aPdu.DepTime
	ctx.FlightPlan.ActualDepTime = aPdu.ActualDepTime
	ctx.FlightPlan.CruiseAlt = aPdu.CruiseAlt
	ctx.FlightPlan.Dest = aPdu.Dest
	ctx.FlightPlan.EnrouteHour = aPdu.EnrouteHour
	ctx.FlightPlan.EnrouteMin = aPdu.EnrouteMin
	ctx.FlightPlan.FobHour = aPdu.FobHour
	ctx.FlightPlan.FobMin = aPdu.FobMin
	ctx.FlightPlan.AlterDest = aPdu.AlterDest
	ctx.FlightPlan.Remark = aPdu.Remark
	ctx.FlightPlan.Route = aPdu.Route
	c.SetContext(ctx)
	sendMessageToAll(c, pdu.Serialize(aPdu.GetHeader(), aPdu.ToTokens()))
}

func onMetarRequestReceived(aPdu *pdu.MetarRequest, c gnet.Conn) {
	ctx := GetConnContext(c)
	if !ctx.Authorized {
		return
	}
	if ctx.Callsign != aPdu.From {
		return
	}
	metar := util.GetWeatherByICAO(aPdu.ICAO)
	if metar == "" {
		sendErrorMessage(aPdu.From, pdu.NoWeatherProfile, ctx.Cid, "No such weather data", c)
		return
	}
	bPdu := pdu.NewPDUMetarResponse("server", aPdu.From, aPdu.ICAO, metar)
	sendConnMessage(c, pdu.Serialize(bPdu.GetHeader(), bPdu.ToTokens()))
}

func sendErrorMessage(callsign string, eType pdu.NetworkError, cid string, msg string, c gnet.Conn) {
	errPdu := pdu.NewPDUError("server", callsign, int(eType), cid, msg)
	sendConnMessage(c, pdu.Serialize(errPdu.GetHeader(), errPdu.ToTokens()))
}

func GetConnContext(c gnet.Conn) server.Context {
	if c.Context() == nil {
		c.SetContext(server.Context{
			IncomingPacket: "",
			Authorized:     false,
			Callsign:       "",
		})
	}
	ctx := c.Context().(server.Context)
	return ctx
}

func readMotd(c gnet.Conn) {
	ctx := GetConnContext(c)
	msgPdu := pdu.NewPDUTextMessage("server", ctx.Callsign, util.SoftwareTitle)
	sendConnMessage(c, pdu.Serialize(msgPdu.GetHeader(), msgPdu.ToTokens()))
	msgPdu = pdu.NewPDUTextMessage("server", ctx.Callsign, viper.GetString("server.motd"))
	sendConnMessage(c, pdu.Serialize(msgPdu.GetHeader(), msgPdu.ToTokens()))

}

func sendMessageToTarget(targetCallsign string, msg string) {
	Connections.Range(func(key, value interface{}) bool {
		conn := key.(gnet.Conn)
		ctx := GetConnContext(conn)
		if !ctx.Authorized {
			return true
		}
		if ctx.Authorized == false {
			return true
		}
		if ctx.Callsign == targetCallsign {
			sendConnMessage(conn, msg)
			return false
		}
		return true
	})
}

func sendMessageToAll(c gnet.Conn, msg string) {
	Connections.Range(func(key, value interface{}) bool {
		conn := key.(gnet.Conn)
		ctx := GetConnContext(conn)
		if ctx.Authorized == false {
			return true
		}
		if c == conn {
			return true
		}
		sendConnMessage(conn, msg)
		return true
	})
}

func sendMessageToInRange(c gnet.Conn, msg string) {
	Connections.Range(func(key, value interface{}) bool {
		conn := key.(gnet.Conn)
		ctx := GetConnContext(conn)
		if ctx.Authorized == false {
			return true
		}
		if c == conn {
			return true
		}
		ctx1 := GetConnContext(c)
		ctx2 := GetConnContext(conn)
		if int(calculateDistance(ctx1, ctx2)) > ctx1.Range+ctx2.Range {
			return true
		}
		sendConnMessage(conn, msg)
		return true
	})
}

func sendMessageToSupervisor(c gnet.Conn, msg string) {
	Connections.Range(func(key, value interface{}) bool {
		conn := key.(gnet.Conn)
		ctx := GetConnContext(conn)
		if ctx.Authorized == false {
			return true
		}
		if c == conn {
			return true
		}
		if ctx.Rating < server.SUP {
			return true
		}
		sendConnMessage(conn, msg)
		return true
	})
}

func sendMessageToATC(c gnet.Conn, msg string) {
	Connections.Range(func(key, value interface{}) bool {
		conn := key.(gnet.Conn)
		ctx := GetConnContext(conn)
		if ctx.Authorized == false {
			return true
		}
		if c == conn {
			return true
		}
		if ctx.Rating <= server.OBS {
			return true
		}
		sendConnMessage(conn, msg)
		return true
	})
}

func sendConnMessage(c gnet.Conn, msg string) {
	//log.Println(msg)
	err := c.AsyncWrite([]byte(msg), nil)
	if err != nil {
		return
	}

}

func rad(d float64) float64 {
	return d * math.Pi / 180.0
}

func calculateDistance(ctx1, ctx2 server.Context) float64 {
	r := 6371.393
	lat1 := ctx1.Lat
	lat2 := ctx2.Lat
	lon1 := ctx1.Lon
	lon2 := ctx2.Lon
	radLat1 := rad(lat1)
	radLat2 := rad(lat2)
	a := radLat1 - radLat2
	b := rad(lon1) - rad(lon2)
	s := 2 * math.Asin(math.Sqrt(math.Pow(math.Sin(a/2), 2)+math.Cos(radLat1)*math.Cos(radLat2)*math.Pow(math.Sin(b/2), 2)))
	s = s * r
	s = math.Round(s*10000.0) / 10000.0
	return s / 1.852
}

func checkSameCallsign(callsign string) bool {
	result := false
	Connections.Range(func(key, value interface{}) bool {
		conn := key.(gnet.Conn)
		ctx := GetConnContext(conn)
		if ctx.Callsign == callsign {
			result = true
			return false
		}
		return true
	})
	return result
}

func getConnByCallsign(callsign string) gnet.Conn {
	var c gnet.Conn = nil
	Connections.Range(func(key, value interface{}) bool {
		conn := key.(gnet.Conn)
		ctx := GetConnContext(conn)
		if ctx.Authorized == false {
			return true
		}
		if ctx.Callsign == callsign {
			c = conn
			return false
		}
		return true
	})
	return c
}

func dealServerQuery(aPdu *pdu.ClientQuery, c gnet.Conn) {
	ctx := GetConnContext(c)
	switch aPdu.Type {
	case "ATC":
		if ctx.Type == server.ATC && ctx.Facility != server.Observer && ctx.Rating != server.OBS {
			bPdu := pdu.NewPDUClientQueryResponse("server", aPdu.From, "ATC", []string{"Y"})
			sendConnMessage(c, pdu.Serialize(bPdu.GetHeader(), bPdu.ToTokens()))
		}
		return
	case "CAPS":
		bPdu := pdu.NewPDUClientQueryResponse("server", aPdu.From, "CAPS", util.Caps)
		sendConnMessage(c, pdu.Serialize(bPdu.GetHeader(), bPdu.ToTokens()))
		return
	case "FP":
		if len(aPdu.Payload) > 0 {
			cUser := getConnByCallsign(aPdu.Payload[0])
			if cUser != nil {
				ctx := GetConnContext(cUser)
				if ctx.FlightPlan.Dep != "" {
					bPdu := pdu.FlightPlan{
						Base: pdu.Base{
							From: ctx.Callsign,
							To:   "",
						},
						Callsign:      ctx.Callsign,
						FlightRules:   ctx.FlightPlan.FlightRules,
						Type:          ctx.FlightPlan.Type,
						TAS:           ctx.FlightPlan.TAS,
						Dep:           ctx.FlightPlan.Dep,
						DepTime:       ctx.FlightPlan.DepTime,
						ActualDepTime: ctx.FlightPlan.ActualDepTime,
						CruiseAlt:     ctx.FlightPlan.CruiseAlt,
						Dest:          ctx.FlightPlan.Dest,
						EnrouteHour:   ctx.FlightPlan.EnrouteHour,
						EnrouteMin:    ctx.FlightPlan.EnrouteMin,
						FobHour:       ctx.FlightPlan.FobHour,
						FobMin:        ctx.FlightPlan.FobMin,
						AlterDest:     ctx.FlightPlan.AlterDest,
						Remark:        ctx.FlightPlan.Remark,
						Route:         ctx.FlightPlan.Route,
					}
					sendConnMessage(c, pdu.Serialize(bPdu.GetHeader(), bPdu.ToTokens()))
				}

			}
		}
	}

}
