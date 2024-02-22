package requests

import (
	"perfomance/test/queries"

	"github.com/hasura/go-graphql-client"
)

func (c Client) SendDocumentRequests() {
	variables := map[string]interface{}{
		"page":           graphql.Int(1),
		"itemsPerPage":   graphql.Int(10),
		"updatedAtOrder": graphql.String("DESC"), // Change to "ASC" if needed
		"createdAtOrder": graphql.String("DESC"), // Change to "ASC" if needed
	}
	c.SendAsync(&queries.GetProductCategories, "GetProductCategories", variables)

	variables = map[string]interface{}{
		"id": graphql.ID("your_product_document_id_here"),
	}
	c.SendAsync(&queries.GetClientDocumentContentUrl, "GetClientDocumentContentUrl", variables)

	variables = map[string]interface{}{
		"id": graphql.ID("your_client_document_id_here"),
	}
	c.SendAsync(&queries.GetProductDocumentContentUrl, "GetProductDocumentContentUrl", variables)

	variables = map[string]interface{}{
		"clientCompany": queries.GetCompany.ClientCompany.Id,
		"page":          graphql.Int(1),
	}
	c.SendAsync(&queries.GetClientDocumentCategories, "GetClientDocumentCategories", variables)

	variables = map[string]interface{}{
		"clientCompany":  graphql.String("your_client_company_id_here"),
		"id":             graphql.String("your_client_document_category_id_here"),
		"documentStatus": graphql.String(""),
		"page":           graphql.Int(1),
		"updatedAtOrder": graphql.String("DESC"), // Change to "ASC" if needed
		"itemsPerPage":   graphql.Int(10),        // Change as needed
	}
	c.SendAsync(&queries.GetClientCategoryDocuments, "GetClientCategoryDocuments", variables)

	variables = map[string]interface{}{
		"id":             graphql.String("your_product_document_category_id_here"),
		"documentStatus": graphql.String(""),
		"page":           graphql.Int(1),
	}
	c.SendAsync(&queries.GetProductCategoryDocuments, "GetProductCategoryDocuments", variables)

	variables = map[string]interface{}{
		"id":             graphql.String("your_product_document_category_id_here"),
		"documentStatus": graphql.String(""),
		"page":           graphql.Int(1),
	}
	c.SendAsync(&queries.GetProductCategoryDocumentsForCR, "GetProductCategoryDocumentsForCR", variables)

	variables = map[string]interface{}{
		"documentStatus": graphql.String(""),
		"page":           graphql.Int(1),
		"itemsPerPage":   graphql.Int(10),        // Change as needed
		"createdAtOrder": graphql.String("DESC"), // Change to "ASC" if needed
		"client":         graphql.String(""),
		"clientCompany":  graphql.String(""),
	}
	c.SendAsync(&queries.GetAllClientsDocuments, "GetAllClientsDocuments", variables)

	variables = map[string]interface{}{
		"itemsPerPage":   graphql.Int(15),
		"documentStatus": graphql.String(""),
		"page":           graphql.Int(1),
		"createdAtOrder": graphql.String("DESC"),
	}
	c.SendAsync(&queries.GetAllProductDocuments, "GetAllProductDocuments", variables)

	variables = map[string]interface{}{
		"productId":      graphql.String(""),
		"documentStatus": graphql.String(""),
		"page":           graphql.Int(1),
		"itemsPerPage":   graphql.Int(15),
		"updatedAtOrder": graphql.String("DESC"),
	}
	c.SendAsync(&queries.GetProductCategoriesForCR, "GetProductCategoriesForCR", variables)

	variables = map[string]interface{}{
		"clientCompany":  graphql.String(""),
		"documentStatus": graphql.String(""),
		"updatedAtOrder": graphql.String("DESC"),
		"itemsPerPage":   graphql.Int(5),
	}
	c.SendAsync(&queries.GetClientDocumentAllCategories, "GetClientDocumentAllCategories", variables)

	variables = map[string]interface{}{
		"clientCompany":  graphql.String(""),
		"documentStatus": graphql.String(""),
		"updatedAtOrder": graphql.String("DESC"),
		"itemsPerPage":   graphql.Int(5),
	}
	c.SendAsync(&queries.GetAllMyDocuments, "GetAllMyDocuments", variables)

	variables = map[string]interface{}{
		"documentStatus": graphql.String(""),
		"page":           graphql.Int(1),
		"deletedAtOrder": graphql.String("DESC"),
	}
	c.SendAsync(&queries.GetAllClientsDeletedDocuments, "GetAllClientsDeletedDocuments", variables)
}
