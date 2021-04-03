package kafka

import (
	"time"

	"github.com/Shopify/sarama"
	"github.com/beego/beego/v2/adapter/logs"
)

var (
	client sarama.SyncProducer
)

func InitKafka(addr string) (err error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logs.Error("sarama.NewSyncProducer err:%v", err)
		return
	}

	logs.Debug("init kafka succ")
	return
}

func SendToKafka(data, topic string) (err error) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}

	part, offset, err := client.SendMessage(msg)
	if err != nil {
		logs.Error("send message failed, err:%v data:%v topic:%v", data, err, topic)
		return
	}

	logs.Debug("send succ, part:%v offset:%v topic:%v", part, offset, topic)

	return
}
