package config

import (
	"AdvancedFlightServer/util"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func defaultConfig() {
	viper.SetDefault("server.port", 6809)
	viper.SetDefault("db.addr", "127.0.0.1")
	viper.SetDefault("db.port", 3306)
	viper.SetDefault("db.username", "root")
	viper.SetDefault("db.password", "123456")
	viper.SetDefault("db.dbname", "flightserver")
	viper.SetDefault("server.motd", fmt.Sprintf("Advanced Flight Server v%s. Open source by AFcPPe under CC-BY-NC-SA 4.0", util.Version))
	viper.SetDefault("server.weather_url", "http://metar.flightsim.top/getall")
	viper.SetDefault("server.output_dir", "./whazzup.json")
	err := viper.SafeWriteConfigAs("config.ini")
	if err != nil {
		return
	}
}

func ReadConfig() {
	log.Println("Reading config file from config.ini...")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	defaultConfig()
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error while reading config file. Using default config.")
		log.Fatal(err)
		return
	}

}
