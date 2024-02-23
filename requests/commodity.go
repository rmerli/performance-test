package requests

import (
	"perfomance/test/queries"

	"github.com/hasura/go-graphql-client"
)

func (c Client) SendCommodityRequests() {
	variables := map[string]interface{}{
		"page":       graphql.Int(1),
		"company_id": graphql.String(queries.GetClientCompanies.ClientCompanies.Collection[0].Id),
		"productId":  graphql.String(queries.GetProducts.Products.Edges[0].Node.Id),
	}
	c.SendAsync(&queries.GetAllYears, "GetAllYears", variables)
}
