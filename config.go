package main

import (
	"elk/models/tailf"
	"errors"
	"fmt"

	"github.com/beego/beego/v2/core/config"
)

type Config struct {
	logLevel string
	LogPath  string

	chanSize    int
	kafkaAddr   string
	collectConf []tailf.CollectConf
}

var (
	appConfig *Config
)

func loadCollectConf(conf config.Configer) (err error) {
	var cc tailf.CollectConf

	cc.LogPath, err = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::log_path")
		return
	}

	cc.Topic, err = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		err = errors.New("invalid collect::log_topic")
		return
	}

	appConfig.collectConf = append(appConfig.collectConf, cc)
	return
}

func LoadConf(confType, filename string) (err error) {

	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("NewConfig fail err", err)
		return
	}

	appConfig = &Config{}

	appConfig.logLevel, err = conf.String("logs::log_level")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug"
	}

	appConfig.LogPath, err = conf.String("logs::log_path")
	if len(appConfig.LogPath) == 0 {
		appConfig.LogPath = "./logs"
	}

	appConfig.chanSize, err = conf.Int("collect::chan_size")
	if err != nil {
		appConfig.chanSize = 100
		return
	}

	err = loadCollectConf(conf)
	if err != nil {
		return fmt.Errorf("Invalid kafka addr")
	}

	appConfig.kafkaAddr, err = conf.String("kafka::srever_addr")
	if len(appConfig.LogPath) == 0 {
		fmt.Printf("load collect conf failed err:%v\n", err)
		return
	}

	return
}
