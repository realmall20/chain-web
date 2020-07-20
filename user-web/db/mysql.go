package db

import (
	"bytes"
	"chain-web/user-web/config"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)
import "github.com/jinzhu/gorm"

var SqlDB *gorm.DB
var (
	DbType   string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
)
func SetUp() {
	//创建一个数据库的连接
	var err error
	DbType = config.DatabaseConfig.DBType
	Host = config.DatabaseConfig.Host
	Port = config.DatabaseConfig.Port
	Name = config.DatabaseConfig.Name
	Username = config.DatabaseConfig.Username
	Password = config.DatabaseConfig.Password
	var conn bytes.Buffer
	conn.WriteString(Username)
	conn.WriteString(":")
	conn.WriteString(Password)
	conn.WriteString("@tcp(")
	conn.WriteString(Host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(Port))
	conn.WriteString(")")
	conn.WriteString("/")
	conn.WriteString(Name)
	conn.WriteString("?charset=utf8&parseTime=True&loc=Local&timeout=1000ms")

	SqlDB, err = gorm.Open(DbType,  conn.String())
	SqlDB.LogMode(true)
	if err != nil {
		panic("failed to connect database")
	}
}

