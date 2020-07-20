package router

import "github.com/gin-gonic/gin"
import ."chain-web/user-web/handler"

func Init(router *gin.Engine)  {
	outAuthRouter(router)
}
// 不需要认证的接口
func outAuthRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.POST("/test/insert", Insert)
}
