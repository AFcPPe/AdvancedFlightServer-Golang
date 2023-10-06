package tasks

import (
	"AdvancedFlightServer/service"
	"encoding/json"
	"github.com/panjf2000/gnet/v2"
	"github.com/spf13/viper"
	"io/ioutil"
)

func Output() {
	var data = make(map[string]any)
	var pilot = make(map[string]any)
	var atc = make(map[string]any)
	service.Connections.Range(func(key, value interface{}) bool {
		conn := key.(gnet.Conn)
		ctx := service.GetConnContext(conn)
		if ctx.Authorized == false {
			return true
		}
		if ctx.Type == 1 {
			jCtx, _ := json.Marshal(ctx)
			m := make(map[string]interface{})
			json.Unmarshal(jCtx, &m)
			delete(m, "Authorized")
			delete(m, "Facility")
			delete(m, "Frequencies")
			delete(m, "IncomingPacket")
			delete(m, "Type")
			delete(m, "Range")
			delete(m, "Rating")
			pilot[ctx.Callsign] = m
		} else if ctx.Type == 0 {
			jCtx, _ := json.Marshal(ctx)
			m := make(map[string]interface{})
			json.Unmarshal(jCtx, &m)
			delete(m, "Authorized")
			delete(m, "Bank")
			delete(m, "FlightPlan")
			delete(m, "GroundSpeed")
			delete(m, "Heading")
			delete(m, "Identing")
			delete(m, "IncomingPacket")
			delete(m, "Pitch")
			delete(m, "PressureAltitude")
			delete(m, "SquawkCode")
			delete(m, "SquawkingModeC")
			delete(m, "TrueAltitude")
			delete(m, "Type")
			atc[ctx.Callsign] = m
		}
		return true
	})
	data["pilots"] = pilot
	data["controllers"] = atc
	js, _ := json.Marshal(data)
	err := ioutil.WriteFile(viper.GetString("server.output_dir"), js, 0644)
	if err != nil {
		return
	}
}
