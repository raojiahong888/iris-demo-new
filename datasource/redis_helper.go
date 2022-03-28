package datasource

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"iris-demo-new/config"
	"iris-demo-new/slog"
	"sync"
	"time"
)

var once sync.Once
var cacheInstance *redisConn

type redisConn struct {
	Pool      *redis.Pool
	ShowDebug bool
}

func (r *redisConn) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := r.Pool.Get()
	defer conn.Close()
	start := time.Now().UnixNano()
	reply, err = conn.Do(commandName, args...)
	if err != nil {
		if e := conn.Err(); e != nil {
			slog.Errorf("redis connect failed, %s", e)
		}
	}
	if config.Setting.Redis.Debug {
		end := time.Now().UnixNano()
		_ = fmt.Sprintf("execution time:%d", (end-start)/1000)
	}
	return
}

func InstanceRedis() *redisConn {
	once.Do(func() {
		cacheInstance = newCache()
	})
	return cacheInstance
}

func newCache() *redisConn {
	pool := redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			password := redis.DialPassword(config.Setting.Redis.Password)
			conn, err = redis.Dial("tcp", fmt.Sprintf("%s", config.Setting.Redis.Host), password)
			if err != nil {
				slog.Errorf("redis connect failed, %s", err)
				return nil, err
			}
			return conn, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         config.Setting.Redis.MaxIdle,
		MaxActive:       config.Setting.Redis.MaxActive,
		IdleTimeout:     config.Setting.Redis.IdleTimeout,
		Wait:            config.Setting.Redis.Wait,
		MaxConnLifetime: config.Setting.Redis.MaxConnLifetime,
	}
	cacheInstance = &redisConn{
		Pool:      &pool,
	}
	return cacheInstance
}