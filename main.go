package main

import (
	// 引用 models 包
	_ "Beego_Backend/routers"

	_ "github.com/mattn/go-sqlite3"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql" // 匿名導入 MySQL 驅動
)

func init() {
	// 註冊 MySQL 驅動
	// orm.RegisterDriver("mysql", orm.DRMySQL)
	// 註冊 SQLite3 驅動
	orm.RegisterDriver("sqlite3", orm.DRSqlite)

	// 註冊默認資料庫
	// orm.RegisterDataBase("default", "mysql", "root:ad112345@tcp(127.0.0.1:3306)/Will?charset=utf8&parseTime=True&loc=Local")
	orm.RegisterDataBase("default", "sqlite3", "./data/database.db")

	// 自動創建表 only when you need to create table
	// orm.RunSyncdb("default", false, true)
}

func main() {
	// o := orm.NewOrm()

	// // 測試插入資料
	// user := models.User{Name: "Alice", Email: "alice@example.com"}
	// id, err := o.Insert(&user)
	// if err == nil {
	// 	fmt.Printf("資料插入成功 ID: %d\n", id)
	// } else {
	// 	fmt.Printf("資料插入失敗: %s\n", err.Error())
	// }
	beego.Run()
}
