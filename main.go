package main

import (
	_ "Beego_Backend/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

