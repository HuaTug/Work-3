package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
	"time"
)

var DB *gorm.DB

func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功")
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 20)
	DB = db
	migration()
}
