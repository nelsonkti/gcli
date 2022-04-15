/**
** @创建时间 : 2021/12/6 15:29
** @作者 : fzy
 */
package xetcd

import (
	"crypto/tls"
	"crypto/x509"
	"demod/lib/logger"
	clientv3 "go.etcd.io/etcd/client/v3"
	"io/ioutil"
	"time"
)

var Client *clientv3.Client
var Pfx string

type Config struct {
	Endpoints   []string
	DialTimeout time.Duration
	OpenTLS     bool
	TlsPath     string
	CAConfig    CAConfig
	Pfx         string
	Env         string
}

type CAConfig struct {
	Path    string
	Cert    string
	CertKey string
	Ca      string
}

func New(cfg Config) {
	// 启动 etcd 服务

	var err error

	env(&cfg)

	Client, err = clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: dialTimeout(cfg.DialTimeout),
		TLS:         tlsConfig(cfg),
	})

	etcdPfx(cfg)

	if err != nil {
		logger.Sugar.Info("启动失败")
		logger.Sugar.Error(err)
		panic(err)
	}
}

func dialTimeout(duration time.Duration) time.Duration {

	if duration == 0 {
		duration = defaultDialTimeout()
	}

	return duration
}

// 设置tls
func tlsConfig(cfg Config) *tls.Config {

	cafg := cfg.CAConfig

	// 没有开启
	if !cfg.OpenTLS {
		return nil
	}

	if cfg.TlsPath == "" {
		panic("缺少文件路径")
	}

	cafg.Path = cfg.TlsPath

	defaultTlsCa(&cafg)

	caData, err := ioutil.ReadFile(cafg.Ca)
	if err != nil {
		panic(err)
	}

	cert, err := tls.LoadX509KeyPair(cafg.Cert, cafg.CertKey)
	if err != nil {
		panic(err)
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caData)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	}

	return tlsConfig
}

func Close() {
	defer Client.Close()
}

func etcdPfx(cfg Config) {
	Pfx = cfg.Pfx
}

func defaultDialTimeout() time.Duration {
	return 5 * time.Second
}

func defaultTlsCa(cafg *CAConfig) {

	if cafg.CertKey == "" {
		cafg.CertKey = "Key"
	}

	cafg.CertKey = cafg.Path + cafg.CertKey

	if cafg.Cert == "" {
		cafg.Cert = "Cert"
	}
	cafg.Cert = cafg.Path + cafg.Cert

	if cafg.Ca == "" {
		cafg.Ca = "CAcert"
	}

	cafg.Ca = cafg.Path + cafg.Ca
}

func env(cfg *Config) {
	if cfg.Env == "production" {
		cfg.TlsPath = cfg.TlsPath + "prod" + "/"
	} else {
		cfg.TlsPath = cfg.TlsPath + "dev" + "/"
	}
}
