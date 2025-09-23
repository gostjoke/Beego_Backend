package controllers

import (
	"Beego_Backend/graphql"

	"github.com/graphql-go/handler"
)

func (c *GraphqlController) PO_PN_Post() {
	schema := graphql.Schema

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	h.ServeHTTP(c.Ctx.ResponseWriter, c.Ctx.Request)
}
