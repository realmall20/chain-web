package main

import (
	"chain-web/pkg/config"
	"chain-web/pkg/db"
	"chain-web/pkg/router"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	db.SetUp()
	engine := gin.Default()
	router.Init(engine)
	engine.Run(":8181")
}
