package handler

import (
	"chain-web/user-web/model"
	"chain-web/user-web/response"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
import . "chain-web/user-web/nt"

//通过用户手机号码+身份证号码+EID编码
func CreateChainAddress(c *gin.Context) {
	var user model.User
	user.Phone = c.PostForm("phone")
	user.IdCard = c.PostForm("idCard")
	user.Eid = c.PostForm("eid")
	//TODO 生成一个伪地址保存到数据库里面
	user.FakeChainAddr = fmt.Sprintf("%x", md5.Sum([]byte(user.Phone+user.IdCard+user.Eid)))
	err :=user.Insert()
	if err != nil{
		response.FailWithMessage("数据保存失败",c)
	}else {
		response.Ok(c)
	}
}

//通过“用户手机号”，到“全民数据链”--即国金公链 查询得到用户真实的区块链地址，存到数据库中 chain_addr: 区块链地址 中
func UserDetail(c *gin.Context) {
	phone := c.Query("phone")
	result := GetNtUserDetailResp(phone)
	if result.Code == 0 {
		//TODO 获取 result.Data数据里面的区块链地址，保存到数据库
		var user model.User
		user.ChainAddr = result.Data["chain_addr"]
		user.SpaceUsed=result.Data["space_used"]
		user.Update(phone)
		c.JSON(http.StatusOK, res(0, result.Msg, result.Data))
	} else {
		c.JSON(http.StatusOK, res(result.Code, result.Msg, result.Data))
	}
}
