package requests

import (
	"perfomance/test/queries"

	"github.com/hasura/go-graphql-client"
)

func (c Client) SendProductRequests() {
	variables := map[string]interface{}{
		"documentStatus": graphql.String(""),
		"page":           graphql.Int(1),
		"deletedAtOrder": graphql.String("DESC"),
	}
	c.SendAsync(&queries.GetDeletedProduct, "GetDeletedProduct", variables)

	variables = map[string]interface{}{
		"status":       graphql.String(""),
		"page":         graphql.Int(1),
		"itemsPerPage": graphql.Int(10), // Change this value as needed
	}
	c.SendAsync(&queries.GetProductAccessRequests, "GetProductAccessRequests", variables)

}
