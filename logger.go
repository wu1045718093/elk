package main

import (
	"encoding/json"
	"fmt"

	"github.com/beego/beego/v2/core/logs"
)

func converLogLevel(level string) int {

	switch level {
	case "debug":
		return logs.LevelDebug
	case "warm":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}

	return logs.LevelDebug

}
func InitLogger() (err error) {

	config := make(map[string]interface{})
	config["filename"] = appConfig.LogPath
	config["level"] = converLogLevel(appConfig.logLevel)

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("Marshal fail err", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))

	return
}
