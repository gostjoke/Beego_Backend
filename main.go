package main

import (
	// 引用 models 包
	_ "Beego_Backend/routers"
	"context"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql" // 匿名導入 MySQL 驅動
	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	Rdb *redis.Client
)

func initRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 預設 port
		Password: "",               // 如果有設定密碼要填
		DB:       0,                // 使用 DB0
	})

	// 測試連線
	pong, err := Rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Redis 連線失敗: %v", err))
	}
	fmt.Println("連線成功:", pong)
}

func init() {
	// 註冊 MySQL 驅動
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 註冊 SQLite3 驅動
	// orm.RegisterDriver("sqlite3", orm.DRSqlite)

	// 註冊默認資料庫
	// orm.RegisterDataBase("default", "mysql", "root:ad112345@tcp(127.0.0.1:3306)/Will?charset=utf8&parseTime=True&loc=Local")
	orm.RegisterDataBase("default", "mysql", "root:will123@tcp(127.0.0.1:3306)/beegodb?charset=utf8&parseTime=True&loc=Local")
	// orm.RegisterDataBase("default", "sqlite3", "./data/database.db")

	// 自動創建表 only when you need to create table
	orm.RunSyncdb("default", false, true) // 第二個參數表示是否強制重新創建表，第三個參數表示是否輸出詳細信息

	// 初始化 Redis
	initRedis()
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
