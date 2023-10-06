package util

import (
	"github.com/robfig/cron/v3"
)

func SetupCron() {
	go SyncWeather()
	c := cron.New(cron.WithSeconds())
	spec := "0 */5 * * * *"
	c.AddFunc(spec, SyncWeather)
	c.Start()
}
