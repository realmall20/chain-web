package handler

import (
	"chain-web/user-web/model"
	"chain-web/user-web/response"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"sort"
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
		var user model.User
		user.ChainAddr = result.Data["chain_addr"]
		user.SpaceUsed=result.Data["space_used"]
		user.Update(phone)
		response.OkWithData(result,c)
	} else {
		response.FailWithMessage("获取数据失败",c)
	}
}

//验证签名
func verify(c *gin.Context){
	timestamp := c.Query("timestamp") // 13位时间戳
	nonce := c.Query("nonce") //11位字符串
	sign := c.Query("sign") //
	device_num := c.Query("device_num") // 设备号码
	mobile_os := c.Query("mobile_os_version") // 设备操作系统版本号
	lng := c.Query("lng")
	lat := c.Query("lat")
}
