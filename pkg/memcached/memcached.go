package memcached

import (
	"fmt"
	"redis/pkg/conf"

	"github.com/bradfitz/gomemcache/memcache"
)

func MemcachedConn(config *conf.Conf) (*memcache.Client, error) {
	mc := memcache.New(fmt.Sprintf("%s:%d", config.Memcached.Host, config.Memcached.Port))

	if err := mc.Ping(); err != nil {
		return nil, err
	}

	return mc, nil
}