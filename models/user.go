package models

import "github.com/beego/beego/v2/client/orm"

// User 表模型
type User struct {
	Id    int    `orm:"auto"`      // 自動遞增主鍵
	Name  string `orm:"size(100)"` // 名稱字段，最大長度 100
	Email string `orm:"size(100)"` // 電子郵件字段，最大長度 100
}

func init() {
	// 註冊模型
	orm.RegisterModel(new(User))
}
