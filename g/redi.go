package g

import (
	"time"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func RediInt(reply interface{}, err error) (int, error) {
	return redis.Int(reply, err)
}

func GetRedis(addr string) *redis.Pool {
	fmt.Println("get redis addr ", addr)
	return &redis.Pool{
		MaxIdle: 200,
		MaxActive: 0,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("ping")
			return err
		},
	}
}
