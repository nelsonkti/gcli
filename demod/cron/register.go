package cron

import (
	"demod/cron/base"
)

func RegisterCrons(address, password string) {
	cron := base.StartCronTab(address, password)

	//注册定时任务

	go cron.Run()
}
