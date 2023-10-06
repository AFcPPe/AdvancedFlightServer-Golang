package main

import (
	"AdvancedFlightServer/config"
	"AdvancedFlightServer/service"
	"AdvancedFlightServer/tasks"
	"AdvancedFlightServer/util"
)

func main() {
	config.ReadConfig()
	util.ConnectToDatabase()
	util.SetupCron()
	tasks.SetupCron()
	service.StartServer()
}
