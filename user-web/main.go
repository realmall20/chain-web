package main

import (
	"chain-web/user-web/router"
	"chain-web/user-web/validate"
	"github.com/gin-gonic/gin"
)

func main(){
	engine := gin.Default()
	validate.Init()
	router.Init(engine)
	engine.Run(":8089")
}

