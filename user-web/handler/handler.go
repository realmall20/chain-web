package handler

import (
    "crypto/md5"
    "database/sql"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)
import ."chain-web2/user-web/nt"
import _ "github.com/go-sql-driver/mysql"
//通过用户手机号码+身份证号码+EID编码
func CreateChainAddress(c *gin.Context){
   phone :=c.PostForm("phone")
   idCard := c.PostForm("idCard")
   eid := c.PostForm("eid")
   //TODO 生成一个伪地址保存到数据库里面
   chainAddr := fmt.Sprintf("%x", md5.Sum([]byte(phone+idCard+eid)))
   db, _ := sql.Open("mysql","root:BVYRDR2Z7X@tcp(192.168.8.105)/test")
   db.Exec(
        "INSERT INTO user (phone,id_card,eid,chain_addr) VALUES (?, ?,?,?)",
        phone,
        idCard,
        eid,
        chainAddr,
    )
   c.JSON(http.StatusOK,res(0,"success",nil))
}


//通过“用户手机号”，到“全民数据链”--即国金公链 查询得到用户真实的区块链地址，存到数据库中 chain_addr: 区块链地址 中
func UserDetail(c *gin.Context) {
    phone :=  c.Query("phone_number")
    result := GetNtUserDetailResp(phone)
    if  result.Code == 0 {
        //TODO 获取 result.Data数据里面的区块链地址，保存到数据库
        c.JSON(http.StatusOK, res(0, result.Msg, result.Data))
    }  else {
        c.JSON(http.StatusOK, res(result.Code, result.Msg, result.Data))
    }
}





