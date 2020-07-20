package handler

import (
	"chain-web/user-web/model"
	"chain-web/user-web/response"
	"github.com/gin-gonic/gin"
)

func Insert(c *gin.Context) {
	var user model.Test
	c.ShouldBind(&user)
	err :=user.Insert()
	if err != nil{
		response.FailWithMessage("数据保存失败",c)
	}else {
		response.Ok(c)
	}
}



