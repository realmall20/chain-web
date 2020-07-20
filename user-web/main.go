package main

import (
	"chain-web/user-web/config"
	"chain-web/user-web/db"
	"chain-web/user-web/router"
	"github.com/gin-gonic/gin"
)

func main(){
	config.Init()
	db.SetUp()
	engine := gin.Default()
    router.Init(engine)
	engine.Run(":8089")
}

