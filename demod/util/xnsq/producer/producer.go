/**
** @创建时间 : 2021/11/13 16:05
** @作者 : fzy
 */
package producer

import (
	"demod/lib/helper"
	"demod/lib/logger"
	"demod/util/xnsq/service/registry"
	"github.com/nsqio/go-nsq"
	"time"
)

var Separator = "@"
var producer *nsq.Producer
var Options registry.Options

func StartNsqProducer(opt registry.Options) {
	Options = opt
	if producer != nil {
		return
	}

	var err error
	cfg := nsq.NewConfig()
	producer, err = nsq.NewProducer(opt.NsqAddress, cfg)
	if nil != err {
		logger.Sugar.Info(err)
		panic("nsq new panic")
	}

	err = producer.Ping()
	if nil != err {
		logger.Sugar.Info(err)
		panic("nsq ping panic")
	}

	SetOngoingIp(opt.LocalAddress)

	// 清理包含ip的旧topic
	go ClearContainIpTopic()
}

type Producer struct {
}

func (p *Producer) Publish(topic, data string) {
	err := producer.Publish(topic, []byte(Separator+data))
	if err != nil {
		logger.Sugar.Error(err)
	}
}

func (p *Producer) DelayPublish(topic, data string, delay time.Duration) {
	err := producer.DeferredPublish(topic, delay, []byte(Separator+data))
	if err != nil {
		logger.Sugar.Error(err)
	}
}

func (p *Producer) DeferredPublish(deviceId, serverAddress, topic, data string, delay time.Duration) {
	if serverAddress == "" {
		//logger.Sugar.Infow("device not online:", "device_id:", deviceId, "topic:", topic)
		return
	}

	err := producer.DeferredPublish(serverAddress+"."+topic, delay, []byte(deviceId+Separator+data))
	if err != nil {
		logger.Sugar.Error(err)
	}
}

// 指定服务发布
func (p *Producer) AssignServerPublish(serverAddress, topic, data string) {
	err := producer.Publish(serverAddress+"."+topic, []byte(Separator+data))
	if err != nil {
		logger.Sugar.Error(err)
	}
}

func (p *Producer) AssignUuidPublish(uuid, topic, data string) {
	if uuid == "" {
		logger.Sugar.Error("uuid为空")
		return
	}
	err := producer.Publish(topic, []byte(uuid+Separator+data))
	if err != nil {
		logger.Sugar.Error(err)
	}
}

func (p *Producer) StopProducer() {
	if producer != nil {
		producer.Stop()
	}

	SetEndedIp(helper.GetLocalIP())

	logger.Sugar.Info("stop nsq producer")
}
