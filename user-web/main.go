package main

import "github.com/gin-gonic/gin"
import ."chain-web2/user-web/handler"

func main(){
	router := gin.Default()
	router.GET("/detail", UserDetail)
	router.POST("/createChainAddress", CreateChainAddress)
	router.Run(":8080")
}
