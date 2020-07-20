package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)
import ."chain-web/user-web/handler"

func Init(router *gin.Engine)  {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notEmpty", notEmpty)
	}
	outAuthRouter(router)
}
// 不需要认证的接口
func outAuthRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.POST("/test/insert", Insert)
}


var notEmpty validator.Func = func(fl validator.FieldLevel) bool {
	str, ok := fl.Field().Interface().(string)
	if ok {
		return str != ""
	}
	return true
}