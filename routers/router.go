package routers

import (
	"Beego_Backend/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// user routers
	InitUserRouter()

	beego.Router("/", &controllers.MainController{})
	beego.Router("/items/", &controllers.ItemController{}, "get:Get")
	beego.Router("/items/:id", &controllers.ItemController{}, "get:GetOne;put:Put;delete:Delete")
	beego.Router("/po/:id", &controllers.PoController{}, "get:GetPo") // http://127.0.0.1:4173/po/1
}
