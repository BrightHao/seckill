package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/BrightHao/seckill/view"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		// 普通方式
		v1.GET("/index", func(c *gin.Context)  {
			view.Seckill()
			c.String(http.StatusOK, "秒杀程序!")
		})
		// 乐观锁方式
		v1.GET("watch", func(c *gin.Context) {
			view.Watch()
			c.String(http.StatusOK, "乐观锁秒杀!")
		})
	}
	r.Run()
}

