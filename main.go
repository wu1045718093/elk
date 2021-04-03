package main

import (
	"elk/models/kafka"
	"elk/models/tailf"
	"fmt"

	_ "github.com/beego/beego/v2/core/config/xml"
	"github.com/beego/beego/v2/core/logs"
)

func main() {

	filename := "./conf/logagent.conf"

	err := LoadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed, err:%v", err)
		panic("load conf failed")
	}

	err = InitLogger()
	if err != nil {
		fmt.Printf("load logger failed, err:%v", err)
		panic("load logger failed")
	}

	logs.Debug("load conf success, config:%v", appConfig)

	err = tailf.InitTail(appConfig.collectConf, appConfig.chanSize)
	if err != nil {
		logs.Error("init tail failed, err:%v", err)
		return
	}
	logs.Debug("init tailf success")

	err = kafka.InitKafka(appConfig.kafkaAddr)
	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
		return
	}

	logs.Debug("init all success")

	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed, err:%v", err)
		return
	}

	logs.Info("LogAgent exit")
}
