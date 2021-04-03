package main

import (
	"elk/models/kafka"
	"elk/models/tailf"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

func serverRun() (err error) {
	for {
		msg := tailf.GetOneLine()
		err := SendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
}

func SendToKafka(msg *tailf.TextMsg) (err error) {
	err = kafka.SendToKafka(msg.Msg, msg.Topic)
	return
}
