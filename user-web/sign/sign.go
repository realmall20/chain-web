package sign

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"sort"
)

//验证签名
func VerifySign(c *gin.Context) bool {
	if c.PostForm("timestamp") == ""{
		return false
	}
	if c.PostForm("nonce") == ""{
		return false
	}
	if c.PostForm("sign") == ""{
		return false
	}
	if c.PostForm("device_num") == ""{
		return false
	}
	if c.PostForm("mobile_os") == ""{
		return false
	}
	if c.PostForm("mobile_os_version") == ""{
		return false
	}
	if c.PostForm("lng") == ""{
		return false
	}
	if c.PostForm("lat") == ""{
		return false
	}
	a := make([]string, 0)
	for k := range c.Request.Form {
		if k != "nonce" && k != "sign" {
			a=append(a, k)
		}
	}
	sort.Strings(a)
	paramStr := ""
	for _,v := range a {
		paramStr += v
		paramStr += "="
		paramStr += c.PostForm(v)
	}
	sign := CreateSign(c.PostForm("nonce"), paramStr)
	paramSign := c.PostForm("sign")
	return sign == paramSign

}

//生成验证码
func CreateSign(nonce string, paramStr string) string {
	key := []byte(nonce)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(paramStr))
	return fmt.Sprintf("%x", mac.Sum(nil))

}
