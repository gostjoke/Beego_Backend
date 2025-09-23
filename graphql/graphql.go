// graphql/schema.go
package graphql

import (
	"Beego_Backend/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.Int},
			"name":  &graphql.Field{Type: graphql.String},
			"email": &graphql.Field{Type: graphql.String},
		},
	},
)

var poType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Po",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.Int},
			"poNo":        &graphql.Field{Type: graphql.String},
			"description": &graphql.Field{Type: graphql.String},
		},
	},
)

var pnType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Pn",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.Int},
			"pnNo":        &graphql.Field{Type: graphql.String},
			"description": &graphql.Field{Type: graphql.String},
		},
	},
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"users": &graphql.Field{
			Type: graphql.NewList(userType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				o := orm.NewOrm()
				var users []models.User
				_, err := o.QueryTable("user").All(&users)
				return users, err
			},
		},
		"user": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if !ok {
					return nil, nil
				}
				o := orm.NewOrm()
				user := models.User{Id: id}
				err := o.Read(&user)
				if err != nil {
					return nil, err
				}
				return user, nil
			},
		},
		// curl -X POST http://127.0.0.1:4173/graphql -H "Content-Type: application/json" -d "{\"query\":\"{ pos { id poNo description } }\"}"
		"pos": &graphql.Field{
			Type: graphql.NewList(poType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				o := orm.NewOrm()
				var pos []*models.Po_table
				_, err := o.QueryTable(new(models.Po_table)).All(&pos)
				return pos, err
			},
		},
		"po": &graphql.Field{
			Type: poType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if !ok {
					return nil, nil
				}
				o := orm.NewOrm()
				po := models.Po_table{Id: id}
				err := o.Read(&po)
				if err != nil {
					return nil, err
				}
				o.LoadRelated(&po, "Pn_tables")
				return &po, nil
			},
		},
		"pns": &graphql.Field{
			Type: graphql.NewList(pnType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				o := orm.NewOrm()
				var pns []*models.Pn_table
				_, err := o.QueryTable(new(models.Pn_table)).All(&pns)
				return pns, err
			},
		},
		"pn": &graphql.Field{
			Type: pnType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if !ok {
					return nil, nil
				}
				o := orm.NewOrm()
				pn := models.Pn_table{Id: id}
				err := o.Read(&pn)
				if err != nil {
					return nil, err
				}
				o.Read(&pn.Po_table)
				return &pn, nil
			},
		},
	},
})

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createPo": &graphql.Field{
			Type: poType,
			Args: graphql.FieldConfigArgument{
				"poNo": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				poNo := p.Args["poNo"].(string)
				description, _ := p.Args["description"].(string)

				o := orm.NewOrm()
				po := models.Po_table{PoNo: poNo, Description: description}
				_, err := o.Insert(&po)
				return &po, err
			},
		},
		"createPn": &graphql.Field{
			Type: pnType,
			Args: graphql.FieldConfigArgument{
				"pnNo": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"poId": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				pnNo := p.Args["pnNo"].(string)
				description, _ := p.Args["description"].(string)
				poId, _ := p.Args["poId"].(int)

				o := orm.NewOrm()
				var po *models.Po_table
				if poId > 0 {
					po = &models.Po_table{Id: poId}
					if err := o.Read(po); err != nil {
						return nil, err
					}
				}
				pn := models.Pn_table{PnNo: pnNo, Description: description, Po_table: po}
				_, err := o.Insert(&pn)
				return &pn, err
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    RootQuery,
	Mutation: RootMutation,
})
