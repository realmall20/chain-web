package main

import (
	"github.com/gin-gonic/gin"
)
import ."chain-web/user-web/handler"



func main(){
	router := gin.Default()
	v1 := router.Group("/api/v1/")
	v1.GET("/detail", UserDetail)
	v1.POST("/createChainAddress", CreateChainAddress)
	router.Run(":8089")
}
