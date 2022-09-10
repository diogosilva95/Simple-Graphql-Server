package types

import (
	"github.com/graphql-go/graphql"
)

type Property struct {
	ID        int    `db:"id" json:"id"`
	Title 		string `db:"title" json:"title"`
	Price  		int		 `db:"price" json:"price"`
}

var PropertyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Property",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.Int},
		"title": 		  &graphql.Field{Type: graphql.String},
		"price": 		  &graphql.Field{Type: graphql.Int},
	},
})
