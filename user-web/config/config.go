package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var cfgDatabase *viper.Viper

func Init() {
	viper.AddConfigPath("./user-web/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}
	cfgDatabase = viper.Sub("settings.database")
	if cfgDatabase == nil {
		panic("config not found settings.database")
	}
	DatabaseConfig = InitDatabase(cfgDatabase)
}
type Database struct {
	Name     string
	DBType   string
	Host     string
	Port     int
	Username string
	Password string
}

func InitDatabase(cfg *viper.Viper) *Database {
	return &Database{
		Port:     cfg.GetInt("port"),
		Name:     cfg.GetString("name"),
		Host:     cfg.GetString("host"),
		Username: cfg.GetString("username"),
		Password: cfg.GetString("password"),
		DBType:   cfg.GetString("dbType"),
	}
}

var DatabaseConfig = new(Database)
