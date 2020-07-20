package model

import (
	 "chain-web/user-web/db"
)
type Test struct {
	ID   int64 `json:"id"`
	Name  string  `form:"name" json:"name" binding:"required"`
}

//操作用户指针，不需要返回ID
func (test *Test) Insert() (err error) {
	//添加数据
	err = db.SqlDB.Create(&test).Error
	return err
}


