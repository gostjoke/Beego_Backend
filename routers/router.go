package routers

import (
	"Beego_Backend/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// user routers
	InitUserRouter()
	InitPoRouter()
	beego.Router("/", &controllers.MainController{})
	beego.Router("/items/", &controllers.ItemController{}, "get:Get")
	beego.Router("/items/:id", &controllers.ItemController{}, "get:GetOne;put:Put;delete:Delete")

}
