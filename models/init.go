package models

import "github.com/beego/beego/v2/client/orm"

func init() {
	// 註冊所有模型
	orm.RegisterModel(
		new(User),
		new(Sample),
		new(Po_table),
		new(Pn_table),
		new(Task),
	)
}

func (u *User) TableName() string {
	return "user"
}
