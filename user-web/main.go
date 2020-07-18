package main

import (
	"chain-web/user-web/router"
	"github.com/gin-gonic/gin"
)

func main(){
	engine := gin.Default()
    router.Init(engine)
	engine.Run(":8089")
}
