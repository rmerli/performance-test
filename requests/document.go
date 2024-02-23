package requests

import (
	"perfomance/test/queries"

	"github.com/hasura/go-graphql-client"
)

func (c Client) SendDocumentRequests() {

	variables := map[string]interface{}{
		"id": graphql.ID(queries.GetAllClientsDocuments.ClientCompanyDocuments.Collection[len(queries.GetAllClientsDocuments.ClientCompanyDocuments.Collection)-1].Id),
	}
	c.SendAsync(&queries.GetClientDocumentContentUrl, "GetClientDocumentContentUrl", variables)

	variables = map[string]interface{}{
		"id": graphql.ID(queries.GetAllProductDocuments.ProductDocuments.Collection[len(queries.GetAllProductDocuments.ProductDocuments.Collection)-1].Id),
	}
	c.SendAsync(&queries.GetProductDocumentContentUrl, "GetProductDocumentContentUrl", variables)

	variables = map[string]interface{}{
		"clientCompany":  graphql.String(queries.GetCompany.ClientCompany.Id),
		"id":             graphql.ID(queries.GetClientDocumentCategories.ClientDocumentCategories.Edges[len(queries.GetClientDocumentCategories.ClientDocumentCategories.Edges)-1].Node.Id),
		"page":           graphql.Int(1),
		"updatedAtOrder": graphql.String("DESC"), // Change to "ASC" if needed
		"itemsPerPage":   graphql.Int(10),        // Change as needed
	}
	c.SendAsync(&queries.GetClientCategoryDocuments, "GetClientCategoryDocuments", variables)

	variables = map[string]interface{}{
		"id": graphql.ID(queries.GetProductCategories.ProductDocumentCategories.Edges[len(queries.GetProductCategories.ProductDocumentCategories.Edges)-1].Node.Id),
	}
	c.SendAsync(&queries.GetProductCategoryDocuments, "GetProductCategoryDocuments", variables)

	variables = map[string]interface{}{
		"productId":      graphql.String(queries.GetProducts.Products.Edges[len(queries.GetProducts.Products.Edges)-1].Node.Id),
		"itemsPerPage":   graphql.Int(15),
		"updatedAtOrder": graphql.String("DESC"),
	}
	c.SendAsync(&queries.GetProductCategoriesForCR, "GetProductCategoriesForCR", variables)

	variables = map[string]interface{}{
		"clientCompany":  graphql.String(queries.GetClientCompanies.ClientCompanies.Collection[len(queries.GetClientCompanies.ClientCompanies.Collection)-1].Id),
		"updatedAtOrder": graphql.String("DESC"),
		"itemsPerPage":   graphql.Int(5),
	}
	c.SendAsync(&queries.GetClientDocumentAllCategories, "GetClientDocumentAllCategories", variables)

	variables = map[string]interface{}{
		"clientCompany":  graphql.String(queries.GetClientCompanies.ClientCompanies.Collection[len(queries.GetClientCompanies.ClientCompanies.Collection)-1].Id),
		"updatedAtOrder": graphql.String("DESC"),
		"itemsPerPage":   graphql.Int(5),
	}
	c.SendAsync(&queries.GetAllMyDocuments, "GetAllMyDocuments", variables)

	variables = map[string]interface{}{
		"page":           graphql.Int(1),
		"deletedAtOrder": graphql.String("DESC"),
	}
	c.SendAsync(&queries.GetAllClientsDeletedDocuments, "GetAllClientsDeletedDocuments", variables)
}
