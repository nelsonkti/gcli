package main

import (
	"fmt"
	"demod/config"
	"demod/lib/db"
	"demod/lib/logger"
	"demod/logic/http/model"
)

func main() {
	// 日志
	logger.New(logger.Base(config.AppConf.App.Env, config.AppConf.App.Path.LogPath))

	defer db.DisconnectMysql()
	db.InitMysql()

	err := db.Mysql("db").AutoMigrate(&model.EmployeesBase{})
	fmt.Println(err)
}
