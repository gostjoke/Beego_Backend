package models

// User 表模型
type User struct {
	Id    int    `orm:"auto"`      // 自動遞增主鍵
	Name  string `orm:"size(100)"` // 名稱字段
	Email string `orm:"size(100)"` // 電子郵件字段

	Samples []*Sample `orm:"reverse(many)"` // 反向關聯 Sample 表
}

// Sample 表模型
type Sample struct {
	Id        int    `orm:"auto"`                        // 自動遞增主鍵
	Name      string `orm:"size(100)"`                   // 名稱字段
	User      *User  `orm:"rel(fk)"`                     // 外鍵，關聯 User 表
	CreatedAt string `orm:"auto_now_add;type(datetime)"` // 自動填充創建時間
}
