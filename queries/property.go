package queries

import (
	"go-server/types"

	"log"

	"github.com/graphql-go/graphql"
)

func GetPropertyQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.PropertyType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			log.Printf("[query] user\n")
			var properties []types.Property

			return properties, nil
		},
	}
}
