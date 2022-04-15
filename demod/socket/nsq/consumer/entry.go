/**
** @创建时间 : 2021/11/15 15:14
** @作者 : fzy
 */
package consumer

import (
	"demod/util/xnsq/consumer"
	"demod/util/xnsq/server"
	"demod/util/xnsq/service/registry"
)

func SocketConsumerHandler(opt registry.Options) server.ConsumerHandler {
	consumer.Options = opt
	return &SocketNsqConsumer{Options: opt}
}

type SocketNsqConsumer struct {
	registry.Options
}

func (l *SocketNsqConsumer) Run() {

}
