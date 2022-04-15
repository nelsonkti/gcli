package main

import (
	"fmt"
	"github.com/judwhite/go-svc"
	"demod/config"
	"demod/lib/db"
	"demod/lib/helper"
	"demod/lib/localtion"
	"demod/lib/logger"
	"demod/socket/nsq/consumer"
	"demod/socket/nsq/producer"
	"demod/socket/socketio_server"
	"demod/util/xetcd"
	"demod/util/xnsq"
	"demod/util/xnsq/service/registry"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"sync"
	"syscall"
)

type socketProgram struct {
	once sync.Once
}

func main() {
	p := &socketProgram{}
	if err := svc.Run(p, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL); err != nil {
		logger.Sugar.Fatal(err)
	}
}

// svc 服务运行框架 程序启动时执行Init+Start, 服务终止时执行Stop
func (p *socketProgram) Init(env svc.Environment) error {
	if env.IsWindowsService() {
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}
	return nil
}

func (p *socketProgram) Start() error {

	// 日志
	logger.New(
		logger.Base(config.AppConf.App.Env, config.AppConf.App.Path.LogPath),
		logger.FileName(helper.CurrentFileName()),
	)

	//性能分析
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	db.InitMysql()

	//连接 memcache
	db.ConnectMemcache(config.AppConf.Data.Memcache.Host)

	//连接redis
	db.ConnectRedis(config.AppConf.Data.Redis.Addr, config.AppConf.Data.Redis.Password, 0, "default")

	xetcd.New(xetcd.Config{
		Endpoints: config.AppConf.Etcd.Host,
		OpenTLS:   config.AppConf.Etcd.OpenTls,
		TlsPath:   config.AppConf.Etcd.TlsPath,
		Env:       config.AppConf.App.Env,
		Pfx:       fmt.Sprintf("%s-%s", config.AppConf.App.Name, config.AppConf.App.Env),
	})

	// 启动消息队列服务
	go func() {
		defer helper.RecoverPanic()

		server := xnsq.NewNsqServer(registry.Options{
			NsqAddress:      config.AppConf.Mq.Nsq.Host,
			NSQConsumers:    config.AppConf.Mq.Nsq.Consumer,
			NSQAdminAddress: config.AppConf.Mq.Nsq.AdminAddress,
			Env:             config.AppConf.App.Env,
			LocalAddress:    localtion.GetLocalIP(),
		})

		server.Run(consumer.SocketConsumerHandler(server.Opt))
	}()

	go func() {
		socketio_server.Start(int(config.AppConf.Server.Socket.Port))
	}()

	return nil
}

func (p *socketProgram) Stop() error {
	p.once.Do(func() {
		defer db.DisconnectMysql()
		defer db.DisconnectRedis()
		defer xetcd.Close()
		defer producer.SocketProducer.StopProducer()
		defer socketio_server.StopDevice()
	})
	return nil
}
