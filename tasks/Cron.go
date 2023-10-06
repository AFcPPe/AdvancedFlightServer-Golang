package tasks

import (
	"github.com/robfig/cron/v3"
)

func SetupCron() {
	c := cron.New(cron.WithSeconds())
	spec := "0/5 * * * * ? "
	c.AddFunc(spec, Output)
	c.Start()
}
