package models

import "github.com/beego/beego/v2/client/orm"

func init() {
	// 註冊所有模型
	orm.RegisterModel(
		new(User),
		new(Sample),
	)
}
