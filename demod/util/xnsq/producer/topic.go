/**
** @创建时间 : 2022/3/22 17:50
** @作者 : fzy
 */
package producer

import (
	"demod/lib/logger"
	"demod/util/xetcd"
	"demod/util/xnsq/api"
	string2 "demod/util/xnsq/util/string"
	"time"
)

var nsqApiClient *api.Client

func DeleteTopic(topic string) {
	err := nsqApiClient.Topic().Delete(topic)
	if err != nil {
		logger.Sugar.Error(err)
	}
}

// 清理包含ip的旧topic
func ClearContainIpTopic() {
	time.Sleep(time.Second * 20)
	locker := xetcd.Locker("clear:contain:ip:topic")
	defer locker.Unlock()
	locker.Lock()

	// 获取差集的ip
	ips := GetEndedOldTp()

	// 清楚所有的旧ip记录
	for _, ip := range GetEndedIp() {
		DelEndedIp(ip)
	}

	// 清楚所有的新ip记录
	for _, ip1 := range GetOngoingIp() {
		DelOngoingIp(ip1)
	}

	if len(ips) == 0 {
		return
	}

	// 获取所有的 topic
	nsqApiClient = api.NewClient(Options)
	topics, _ := nsqApiClient.Topic().QueryAll()

	for _, topic := range topics {
		if string2.ContainArray(ips, topic) {
			go DeleteTopic(topic)
		}
	}

}
