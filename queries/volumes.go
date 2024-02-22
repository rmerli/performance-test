package queries

import "github.com/hasura/go-graphql-client"

var GetAllYears struct {
	CommodityVolumesYears struct {
		PaginationInfo struct {
			ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
			LastPage     graphql.Int `graphql:"lastPage"`
			TotalCount   graphql.Int `graphql:"totalCount"`
		}
		Collection []struct {
			Id            graphql.ID     `graphql:"id"`
			Year          graphql.String `graphql:"year"`
			SavingLimit   graphql.Float  `graphql:"savingLimit"`
			TotalExpected graphql.Float  `graphql:"totalExpected"`
		}
	} `graphql:"commodityVolumesYears(page: $page, company_id: $company_id, product_id: $productId)"`
}
