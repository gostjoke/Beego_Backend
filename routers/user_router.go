// routers/user_router.go
package routers

import (
	"Beego_Backend/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func InitUserRouter() {
	beego.Router("/api/users", &controllers.UserController{}, "get:GetAll;post:Create")
	beego.Router("/api/users/:id", &controllers.UserController{}, "get:GetOne;put:Update;delete:Delete")
	beego.Router("/graphql", &controllers.GraphqlController{}, "post:Post")
}
