package requests

import (
	"perfomance/test/queries"

	"github.com/hasura/go-graphql-client"
)

func (c Client) SendCommodityRequests() {
	variables := map[string]interface{}{
		"page":       graphql.Int(1),
		"company_id": graphql.String("your_company_id_here"),
		"productId":  graphql.String("your_product_id_here"),
	}
	c.SendAsync(&queries.GetAllYears, "GetAllYears", variables)
}
