package mutations

import (
	"go-server/types"

	"github.com/graphql-go/graphql"
)

func CreatePropertyMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.PropertyType,
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"price": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			property := &types.Property{
				Title: params.Args["title"].(string),
				Price:  params.Args["price"].(int),
			}

			return property, nil
		},
	}
}

func UpdatePropertyMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.PropertyType,
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"price": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			property := &types.Property{
				Title: params.Args["title"].(string),
				Price:  params.Args["price"].(int),
			}

			return property, nil
		},
	}
}