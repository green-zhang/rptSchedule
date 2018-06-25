package main

import (
	"github.com/green/rptSchedule/config"
	"github.com/green/rptSchedule/crontab"
	"github.com/green/rptSchedule/task"
)

func main() {
	var ch = make(chan string, 1)

	crontab.CronManager.AddFunc(config.GenRptTaskTime, task.GenRpt)
	crontab.StartCron()

	<-ch
}
