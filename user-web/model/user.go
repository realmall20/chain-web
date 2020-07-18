package model

import (
	 "chain-web/user-web/db"
	"log"
)
type User struct {
	ID   int64 `json:"id"`
	Phone  string `json:"phone"`
	IdCard string `json:"id_card"`
	Eid string `json:"eid"`
	ChainAddr string `json:"chain_addr"`
	FakeChainAddr string `json:"fake_chain_addr"`
	UserSpace string `json:"user_space"`
	SpaceUsed string `json:"space_used"`
}

var Users []User

//操作用户指针，不需要返回ID
func (user *User) Insert() (err error) {
	//添加数据
	err = db.SqlDB.Create(&user).Error
	return err
}

//修改
func (user *User) Update(phone string) (updateUser User, err error) {
	if err = db.SqlDB.Where("phone=?",phone).First(&updateUser).Error; err != nil{
	    log.Fatal("获取用户信息失败，phone={}",phone)
		return
	}
	if err = db.SqlDB.Model(&updateUser).Updates(&user).Error; err != nil {
		return
	}
	return
}

