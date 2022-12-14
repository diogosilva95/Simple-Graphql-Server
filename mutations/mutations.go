package mutations

import (
	"github.com/graphql-go/graphql"
)

func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"createProperty": CreatePropertyMutation(),
		"updateProperty": UpdatePropertyMutation(),
	}
}
