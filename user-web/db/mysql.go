package db

import _ "github.com/go-sql-driver/mysql"
import "github.com/jinzhu/gorm"

var SqlDB *gorm.DB
// 先于main 方法执行，实现包级别的数据初始化
// 执行顺序，先变量初始化 -> init() 函数 -> main() 函数
func init() {
	//创建一个数据库的连接
	var err error
	SqlDB, err = gorm.Open("mysql", "root:BVYRDR2Z7X@tcp(192.168.8.105:3306)/test?charset=utf8&parseTime=True&loc=Local")
	SqlDB.LogMode(true)
	if err != nil {
		panic("failed to connect database")
	}
}
