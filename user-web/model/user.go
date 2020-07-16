package model

import (
	 "chain-web2/user-web/db"
)
type User struct {
	ID   int64 `json:"id"`
	Phone  string `json:"phone"`
	IdCard string `json:"id_card"`
	Eid string `json:"eid"`
	ChainAddress string `json:"chain_address"`
	NtChainAddress string `json:"nt_chain_address"`
}

var Users []User

func (user User) Insert() (id int64, err error) {
	//添加数据
	result := db.SqlDB.Create(&user)
	id =user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//修改
func (user *User) Update(phone string) (updateUser User, err error) {
	if err = db.SqlDB.Select([]string{"phone"}).First(&updateUser, phone).Error; err != nil {
		return
	}
	if err = db.SqlDB.Model(&updateUser).Updates(&user).Error; err != nil {
		return
	}
	return
}

