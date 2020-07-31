package nt

// 全民数据链调用
import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	Url = "http://tt.nt-chain.cn"
	//Url = "http://ntapp.nt-chain.cn/service"
	AccessKey = "DJ"
	SecretKey = "acd84337bac34c9f5edd2cd9c7fc6f3f"
)

//http client
var Client = &http.Client{Timeout: 10 * time.Second}

type Response struct {
	Code       int               `json:"code"`
	Msg        string            `json:"msg"`
	ServerTime int64             `json:"server_time"`
	RequestId  string            `json:"requst_id"`
	Data       map[string]string `json:"data"`
}

//通过“用户手机号”，到“全民数据链”--即国金公链 查询得到用户真实的区块链地址，存到数据库中 chain_addr: 区块链地址 中
func GetNtUserDetailResp(phone string) Response {
	attr := "phone_number,chain_addr,spaced_used"
	url := Url + "/service/api/v1/users/detail" + "?" + "attributes=" + attr + "&" + "phone_number=" + phone
	req, err := http.NewRequest("GET", url, nil)
	println(err)
	now := strconv.FormatInt(time.Now().Unix(), 10)
	originSignature := AccessKey + now + SecretKey
	signature := fmt.Sprintf("%x", md5.Sum([]byte(originSignature)))
	req.Header.Add("X-Access-Key", AccessKey)
	req.Header.Add("X-Time", now)
	req.Header.Add("X-Signature", signature)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-App-Type", "Android")
	resp, _ := Client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		bytes, _ := ioutil.ReadAll(resp.Body)
		var response Response
		json.Unmarshal(bytes, &response)
		return response
	} else {
		return Response{Code: 1, Msg: "获取数据错误", RequestId: "21313", Data: nil}
	}
}
