package model

import (
	"chain-web/user-web/db"
	"log"
)

type User struct {
	ID            int64  `form:"id" json:"id"`
	Phone         string `form:"phone" json:"phone"`
	IdCard        string `form:"id_card" json:"id_card"`
	Eid           string `form:"eid"  json:"eid"`
	ChainAddr     string `form:"chain_addr" json:"chain_addr"`
	FakeChainAddr string `form:"fake_chain_addr" json:"fake_chain_addr"`
	UserSpace     string `form:"user_space" json:"user_space"`
	SpaceUsed     string `form:"space_used" json:"space_used"`
}

var Users []User

//操作用户指针，不需要返回ID
func (user *User) Insert() (err error) {
	//添加数据
	err = db.SqlDB.Create(&user).Error
	return
}

//修改
func (user *User) Update(phone string) (updateUser User, err error) {
	if err = db.SqlDB.Where("phone=?", phone).First(&updateUser).Error; err != nil {
		log.Fatal("获取用户信息失败，phone={}", phone)
		return
	}
	if err = db.SqlDB.Model(&updateUser).Updates(&user).Error; err != nil {
		return
	}
	return
}
