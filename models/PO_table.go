package models

type Po_table struct {
	Id          int    `orm:"auto"`      // 自動遞增主鍵
	PoNo        string `orm:"size(30)"`  // 名稱字段
	Description string `orm:"size(100)"` //

	Pn_tables []*Pn_table `orm:"reverse(many)"` // 反向關聯，表示一對多
}

type Pn_table struct {
	Id          int    `orm:"auto"`      // 自動遞增主鍵
	PnNo        string `orm:"size(30)"`  // 名稱字段
	Description string `orm:"size(100)"` //

	Po_table *Po_table `orm:"rel(fk)"` // 外鍵，關聯到 PO_table
}
