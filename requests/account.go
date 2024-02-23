package requests

import (
	"perfomance/test/mutations"
	"perfomance/test/queries"

	"github.com/hasura/go-graphql-client"
)

func (c Client) SendAccountRequests() {
	variables := map[string]interface{}{
		"id":            graphql.ID(mutations.StartAuth.StartAuthentication.Authentication.Account.Id),
		"productStatus": graphql.String(""),
	}
	c.SendAsync(&queries.GetAccount, "GetAccount", variables)

	variables = map[string]interface{}{
		"status":          graphql.String("active"),
		"type":            graphql.String("company representative"),
		"page":            graphql.Int(1),
		"orderByFullName": graphql.String("ASC"),
	}
	c.SendAsync(&queries.GetUsers, "GetUsers", variables)

	variables = map[string]interface{}{
		"id": graphql.ID(mutations.StartAuth.StartAuthentication.Authentication.Account.Id),
	}
	c.SendAsync(&queries.GetAdditionalSAData, "GetAdditionalSAData", variables)
}
