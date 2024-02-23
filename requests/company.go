package requests

import (
	"perfomance/test/queries"

	"github.com/hasura/go-graphql-client"
)

func (c Client) SendCompanyRequests() {

	variables := map[string]interface{}{
		"products_product_id_list": []string{
			string(queries.GetProducts.Products.Edges[len(queries.GetProducts.Products.Edges)-1].Node.Id),
			string(queries.GetProducts.Products.Edges[len(queries.GetProducts.Products.Edges)-2].Node.Id),
		},
		"status_list": []string{
			string(queries.GetClientCompanies.ClientCompanies.Collection[len(queries.GetClientCompanies.ClientCompanies.Collection)-1].Status),
		},
		"itemsPerPage": graphql.Int(500),
		"page":         graphql.Int(1),
	}
	c.SendAsync(&queries.GetClientCompaniesFilteredByProduct, "GetClientCompaniesFilteredByProduct", variables)

	variables = map[string]interface{}{
		"id": queries.GetClientCompanies.ClientCompanies.Collection[len(queries.GetClientCompanies.ClientCompanies.Collection)-1].Id,
	}
	c.SendAsync(&queries.GetClientCompanyProducts, "GetClientCompanyProducts", variables)

	variables = map[string]interface{}{
		"id": queries.GetClientCompanies.ClientCompanies.Collection[len(queries.GetClientCompanies.ClientCompanies.Collection)-1].Id,
	}
	c.SendAsync(&queries.GetCompanyOnboarding, "GetCompanyOnboarding", variables)

	variables = map[string]interface{}{
		"id":           queries.GetClientCompanies.ClientCompanies.Collection[len(queries.GetClientCompanies.ClientCompanies.Collection)-1].Id,
		"itemsPerPage": graphql.Int(100),
	}
	c.SendAsync(&queries.GetCompanyClientView, "GetCompanyClientView", variables)

	variables = map[string]interface{}{
		"id": queries.GetClientCompanies.ClientCompanies.Collection[len(queries.GetClientCompanies.ClientCompanies.Collection)-1].Id,
	}
	c.SendAsync(&queries.GetCompanyStatus, "GetCompanyStatus", variables)

	variables = map[string]interface{}{
		"id": queries.GetCompany.ClientCompany.PrimarySubEntity.Id,
	}
	c.SendAsync(&queries.GetSubEntity, "GetSubEntity", variables)

	variables = map[string]interface{}{
		"id": queries.GetCompany.ClientCompany.PrimarySubEntity.Id,
	}
	c.SendAsync(&queries.GetSubEntityBsupp, "GetSubEntityBsupp", variables)

	variables = map[string]interface{}{
		"page":         graphql.Int(1),
		"itemsPerPage": graphql.Int(500),
		"order":        graphql.String("ASC"),
	}
	c.SendAsync(&queries.GetZohoAccounts, "GetZohoAccounts", variables)

	variables = map[string]interface{}{
		"product":      graphql.String(queries.GetProducts.Products.Edges[len(queries.GetProducts.Products.Edges)-1].Node.Id),
		"status_list":  []graphql.String{},
		"itemsPerPage": graphql.Int(500),
		"client":       graphql.String(""),
	}
	c.SendAsync(&queries.GetSharedWithCompanies, "GetSharedWithCompanies", variables)

	variables = map[string]interface{}{
		"client":       graphql.String(queries.GetClientCompanies.ClientCompanies.Collection[len(queries.GetClientCompanies.ClientCompanies.Collection)-1].Id),
		"product":      graphql.String(queries.GetProducts.Products.Edges[len(queries.GetProducts.Products.Edges)-1].Node.Id),
		"itemsPerPage": graphql.Int(500),
		"page":         graphql.Int(1),
	}
	c.SendAsync(&queries.GetClientViewCompanies, "GetClientViewCompanies", variables)

	variables = map[string]interface{}{
		"product":      graphql.String(queries.GetProducts.Products.Edges[len(queries.GetProducts.Products.Edges)-1].Node.Id),
		"itemsPerPage": graphql.Int(500),
		"page":         graphql.Int(1),
		"client":       graphql.String(queries.GetClientCompanies.ClientCompanies.Collection[len(queries.GetClientCompanies.ClientCompanies.Collection)-1].Id),
	}
	c.SendAsync(&queries.GetClientName, "GetClientName", variables)

	variables = map[string]interface{}{
		"page":         graphql.Int(1),
		"itemsPerPage": graphql.Int(10),
	}
	c.SendAsync(&queries.GetAllClientCompanies, "GetAllClientCompanies", variables)
}
