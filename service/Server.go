package service

import (
	"AdvancedFlightServer/model/pdu"
	"AdvancedFlightServer/model/server"
	"AdvancedFlightServer/util"
	"fmt"
	"github.com/panjf2000/gnet/v2"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type echoServer struct {
	gnet.BuiltinEventEngine
	eng       gnet.Engine
	addr      string
	multicore bool
}

func StartServer() {
	echo := &echoServer{addr: fmt.Sprintf("tcp://:%d", viper.GetInt("server.port")), multicore: true}

	log.Fatal(gnet.Run(echo, echo.addr, gnet.WithMulticore(true), gnet.WithReusePort(true), gnet.WithReuseAddr(true)))
}

func (es *echoServer) OnBoot(eng gnet.Engine) gnet.Action {
	es.eng = eng
	log.Printf("Server listening on port %d.", viper.GetInt("server.port"))
	return gnet.None
}

func (es *echoServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	c.SetContext(server.Context{
		IncomingPacket: "",
		Authorized:     false,
		Callsign:       "",
	})
	Connections.Store(c, true)
	return
}

func (es *echoServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	Connections.Delete(c)
	if err != nil {
		log.Println(err)
	}
	ctx := GetConnContext(c)
	if ctx.Authorized {
		if ctx.Type == server.ATC {
			aPdu := pdu.NewPDUDeleteATC(ctx.Callsign, ctx.Cid)
			sendMessageToAll(c, pdu.Serialize(aPdu.GetHeader(), aPdu.ToTokens()))
		} else {
			aPdu := pdu.NewPDUDeletePilot(ctx.Callsign, ctx.Cid)
			sendMessageToAll(c, pdu.Serialize(aPdu.GetHeader(), aPdu.ToTokens()))
		}
	}

	return
}

func (es *echoServer) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1)
	str := string(buf)
	if c.Context() == nil {
		c.SetContext(server.Context{
			IncomingPacket: "",
			Authorized:     false,
			Callsign:       "",
		})
	}
	ctx := c.Context().(server.Context)
	strToProcess := ctx.IncomingPacket + str
	splitStr := strings.Split(strToProcess, util.PacketDelimiter)
	if strings.HasSuffix(strToProcess, util.PacketDelimiter) {
		ctx.IncomingPacket = ""
	} else {
		ctx.IncomingPacket = splitStr[len(splitStr)-1]
		splitStr = splitStr[:len(splitStr)-1]
	}
	c.SetContext(ctx)
	ServerHandler(splitStr, c)
	return gnet.None
}
