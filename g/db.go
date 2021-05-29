package g

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func dbConn(addr string) *gorm.DB {
	db, err := gorm.Open("mysql", addr)
	if err != nil {
		panic(err)
	}
	return db
}

