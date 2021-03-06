package router

import "github.com/gin-gonic/gin"
import . "chain-web/pkg/handler"

func Init(router *gin.Engine) {
	outAuthRouter(router)
}

// 不需要认证的接口
func outAuthRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	v1.POST("/syncNt", SyncNt)
	v1.POST("/createFakeChainAddr", CreateChainAddress)
}
