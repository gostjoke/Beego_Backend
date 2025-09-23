// routers/user_router.go
package routers

import (
	"Beego_Backend/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func InitPoRouter() {
	beego.Router("/api/pos/:id", &controllers.PoController{}, "get:GetPo") // http://127.0.0.1:4173/po/1
	beego.Router("/api/users", &controllers.UserController{}, "get:GetAll;post:Create")
	beego.Router("/graphql", &controllers.GraphqlController{}, "post:Post")
}
