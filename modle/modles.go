package modles

import (
	"apiTools/libs/config"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	RedisPool *redis.Pool
)

func InitRedis() (err error) {
	RedisPool = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", fmt.Sprintf("%s:%s",
				config.GetString("redis::host"), config.GetString("redis::port")))
			if err != nil {
				return nil, err
			}
			if config.GetString("redis::password") != "" {
				if _, err := conn.Do("AUTH", config.GetString("redis::password")); err != nil {
					conn.Close()
					return nil, err
				}
			}
			return
		},
		MaxIdle:     16,              // 最大空闲连接数
		MaxActive:   32,              // 最大活跃连接数
		IdleTimeout: time.Second * 3, // 最大空闲超时时间
	}
	pool := RedisPool.Get()
	defer pool.Close()
	if err := pool.Err(); err != nil {
		return fmt.Errorf("init redis failed, err: %v", err)
	}
	return
}