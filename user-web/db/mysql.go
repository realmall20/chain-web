package db

import _ "github.com/go-sql-driver/mysql"
import "github.com/jinzhu/gorm"

var SqlDB *gorm.DB

func init() {
	//创建一个数据库的连接
	var err error
	SqlDB, err = gorm.Open("mysql", "root:BVYRDR2Z7X@tcp(192.168.8.105:3306)/test?charset=utf8&parseTime=True&loc=Local")
	SqlDB.LogMode(true)
	if err != nil {
		panic("failed to connect database")
	}
}
