package db

import "github.com/bradfitz/gomemcache/memcache"

var Memcache *memcache.Client

//连接memcache
func ConnectMemcache(server []string) *memcache.Client {
	Memcache = memcache.New(server...)
	return Memcache
}
