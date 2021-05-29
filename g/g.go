package g

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/garyburd/redigo/redis"
)

var (
	Db *gorm.DB
	RedisPool *redis.Pool
	addr = "127.0.0.1:6379"
)

func init()  {
	Db = dbConn("root:rootroot@tcp(localhost:3306)/laravel")
	RedisPool = GetRedis(addr)
	if Db != nil && RedisPool != nil {
		fmt.Println("redis和mysql准备就绪")
	}
}
