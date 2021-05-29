package view

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/BrightHao/seckill/g"
	m "github.com/BrightHao/seckill/model"
)

var (
	pro   = m.Product{}
	store int
)

func init() {
	err := g.Db.Table("products").Select("store").Where("id = ?", 1).Scan(&pro)
	if err != nil {
		fmt.Println("查询余量失败", err)
	}
	store = pro.Store
	fmt.Println("剩余量：", store)
}

func Seckill() {
	pro := m.Product{}
	err := g.Db.Table("products").Select("store").Where("id = ?", 1).Scan(&pro)
	if err != nil {
		fmt.Println("查询余量失败", err)
	}
	store := pro.Store
	if store > 0 {
		tx := g.Db.Begin()
		if err != nil {
			fmt.Println("事务开启失败!", err)
		}
		if tx != nil {
			err := g.Db.Table("products").Where("id = ?", 1).Update("store", gorm.Expr("store - ?", 1))
			if err != nil {
				fmt.Println("秒杀失败", err)
				// 回滚
				tx.Rollback()
			} else {
				fmt.Println("秒杀成功!")
			}
		}
		tx.Commit()
	} else {
		fmt.Println("秒杀结束！")
	}
}

// 乐观锁
func Watch() {
	rc := g.RedisPool.Get()
	defer rc.Close()

	// 监听销量 sales
	_, err2 := rc.Do("watch", "sales")
	if err2 != nil {
		fmt.Println("监听失败！", err2)
	}
	// 获取销量
	sales, err3 := g.RediInt(rc.Do("Get", "sales"))
	if err3 != nil {
		fmt.Println("获取销量失败", err3)
	}
	fmt.Println("当前销量：", sales)
	if sales < store {
		// 抢购成功  扣库存

		// 开启redis事务  支持自动回滚
		rc.Send("multi")
		// 销量+1
		rc.Send("incr", "sales")
		// 执行事务
		rel, err := rc.Do("exec")
		fmt.Println("事务执行后sales的值： ", rel)
		if err != nil {
			fmt.Println("事务执行失败", err)
			// 出现意外BUG。。销毁事务
			_, _ = rc.Do("discard")
		}
		if rel != nil {
			// 事务执行成功
			// 执行扣库存
			g.Db.Table("products").Where("id = ?", 1).Update("store", gorm.Expr("store - ?", 1))
			fmt.Println("扣库存成功！")
		}
	} else {
		fmt.Println("秒杀活动已经结束!")
	}
}
