package redisDb

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

func RedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle		: 3,//最大空闲数
		MaxActive	: 100,//最大活跃数
		IdleTimeout	: 240 * time.Second,//最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func () (redis.Conn, error) {
			//此处对应redis ip及端口号
			conn, err := redis.Dial("tcp", "123.206.201.139:6379")
			if err != nil {
				return nil, err
			}
			//此处1234对应redis密码
			if _, err := conn.Do("AUTH", "123456"); err != nil {
				conn.Close()
				return nil, err
			}
			return conn,err
		},
	}
}
