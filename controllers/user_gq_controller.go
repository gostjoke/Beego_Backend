// controllers/graphql_controller.go
package controllers

import (
	"Beego_Backend/graphql"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/graphql-go/handler"
)

type GraphqlController struct {
	beego.Controller
}

/* windows curl example:
curl -X POST http://127.0.0.1:4173/graphql -H "Content-Type: application/json" -d "{\"query\":\"{ users { id name email } }\"}"
Single user query example:
curl -X POST http://127.0.0.1:4173/graphql -H "Content-Type: application/json" -d "{\"query\":\"{ user(id:1) { id name email } }\"}"
*/

/* POST /graphql
Invoke-WebRequest -Uri "http://127.0.0.1:4173/graphql" `
  -Method POST `
  -Headers @{ "Content-Type" = "application/json" } `
  -Body '{ "query": "{ users { id name email } }" }'
*/

func (c *GraphqlController) Post() {
	h := handler.New(&handler.Config{
		Schema: &graphql.Schema,
		Pretty: true,
	})
	h.ServeHTTP(c.Ctx.ResponseWriter, c.Ctx.Request)
}
