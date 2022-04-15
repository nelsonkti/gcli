/**
** @创建时间 : 2021/12/6 17:38
** @作者 : fzy
 */
package xetcd

import (
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	"sync"
)

func Locker(key string) sync.Locker {
	session, err := concurrency.NewSession(Client)
	if err != nil {
		log.Println("session err :", err)
	}

	locker := concurrency.NewLocker(session, Pfx+key)

	return locker
}
