package main

import (
	"context"
	"fmt"
	"net/http"
	"perfomance/test/mutations"
	"perfomance/test/mutations/inputs"
	"perfomance/test/queries"
	"sync"
	"time"

	"github.com/hasura/go-graphql-client"
)

const url = "https://alpha-portal-api.afsenergy.nl/api/graphql"

var wg sync.WaitGroup

func main() {
	variables := map[string]interface{}{
		"input": inputs.StartAuthenticationInputValues,
	}

	client := graphql.NewClient(url, nil)

	startTime := time.Now().UnixMilli()
	//Authentication
	err := client.Mutate(context.Background(), &mutations.StartAuth, variables)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	endTime := time.Now().UnixMilli()
	elapsedTime := endTime - startTime
	fmt.Printf("Auth took: %d ms\n", elapsedTime)

	token := fmt.Sprintf("Bearer %s", mutations.StartAuth.StartAuthentication.Authentication.AccessToken)

	//Authenticated client
	client = graphql.NewClient(url, nil).WithRequestModifier(func(r *http.Request) {
		r.Header.Add("Authorization", token)
	})

	sendReq(client, &queries.GetProduct, "GetProduct", nil, false)

	sendReq(client, &queries.GetClientCompanies, "GetCompanies", nil, false)

	variables = map[string]interface{}{
		"id":           queries.GetClientCompanies.ClientCompanies.Collection[len(queries.GetClientCompanies.ClientCompanies.Collection)-1].Id,
		"page":         graphql.Int(1),
		"itemsPerPage": graphql.Int(100),
	}
	sendReq(client, &queries.GetCompany, "GetCompany", variables, false)

	variables = map[string]interface{}{
		"status":          graphql.String("active"),
		"type":            graphql.String("company representative"),
		"page":            graphql.Int(1),
		"orderByFullName": graphql.String("ASC"),
	}

	sendAsync(client, &queries.GetUsers, "GetUsers", variables)

	variables = map[string]interface{}{
		"id":            graphql.ID(mutations.StartAuth.StartAuthentication.Authentication.Account.Id),
		"productStatus": graphql.String(""),
	}
	sendAsync(client, &queries.GetAccount, "GetAccount", variables)

	variables = map[string]interface{}{
		"products_product_id_list": []string{
			string(queries.GetProduct.Products.Edges[len(queries.GetProduct.Products.Edges)-1].Node.Id),
			string(queries.GetProduct.Products.Edges[len(queries.GetProduct.Products.Edges)-2].Node.Id),
		},
		"status_list": []string{
			string(queries.GetClientCompanies.ClientCompanies.Collection[len(queries.GetClientCompanies.ClientCompanies.Collection)-1].Status),
		},
		"itemsPerPage": graphql.Int(500),
		"page":         graphql.Int(1), // or any page number you need
	}
	sendAsync(client, &queries.GetClientCompaniesFilteredByProduct, "GetCompanyFilterByProduct", variables)

	variables = map[string]interface{}{
		"id": queries.GetClientCompanies.ClientCompanies.Collection[len(queries.GetClientCompanies.ClientCompanies.Collection)-1].Id,
	}
	sendAsync(client, &queries.GetClientCompanyProducts, "GetClientCompanyProducts", variables)

	variables = map[string]interface{}{
		"id": queries.GetClientCompanies.ClientCompanies.Collection[len(queries.GetClientCompanies.ClientCompanies.Collection)-1].Id,
	}
	sendAsync(client, &queries.GetCompanyOnboarding, "GetCompanyOnboarding", variables)

	variables = map[string]interface{}{
		"id":           queries.GetClientCompanies.ClientCompanies.Collection[len(queries.GetClientCompanies.ClientCompanies.Collection)-1].Id,
		"itemsPerPage": graphql.Int(100),
	}
	sendAsync(client, &queries.GetCompanyClientView, "GetCompanyClientView", variables)

	variables = map[string]interface{}{
		"id": queries.GetClientCompanies.ClientCompanies.Collection[len(queries.GetClientCompanies.ClientCompanies.Collection)-1].Id,
	}
	sendAsync(client, &queries.GetCompanyStatus, "GetCompanyStatus", variables)

	variables = map[string]interface{}{
		"id": queries.GetCompany.ClientCompany.PrimarySubEntity.Id,
	}
	sendAsync(client, &queries.GetSubEntity, "GetSubEntity", variables)

	variables = map[string]interface{}{
		"id": queries.GetCompany.ClientCompany.PrimarySubEntity.Id,
	}
	sendAsync(client, &queries.GetSubEntityBsupp, "GetSubEntityBsupp", variables)

	sendAsync(client, &queries.GetModules, "GetModules", nil)
	sendAsync(client, &queries.GetModulesWithGroups, "GetModulesWithGroups", nil)

	variables = map[string]interface{}{
		"id": mutations.StartAuth.StartAuthentication.Authentication.Account.Id,
	}
	sendAsync(client, &queries.GetCalendar, "GetCalendar", variables)

	variables = map[string]interface{}{
		"page":         graphql.Int(1),
		"itemsPerPage": graphql.Int(500),
		"order":        graphql.String("ASC"),
	}

	sendAsync(client, &queries.GetZohoAccounts, "GetZohoAccounts", variables)

	variables = map[string]interface{}{
		"page":           graphql.Int(1),
		"itemsPerPage":   graphql.Int(10),
		"updatedAtOrder": graphql.String("DESC"), // Change to "ASC" if needed
		"createdAtOrder": graphql.String("DESC"), // Change to "ASC" if needed
	}
	sendAsync(client, &queries.GetProductCategories, "GetProductCategories", variables)

	variables = map[string]interface{}{
		"id": graphql.ID("your_product_document_id_here"),
	}
	sendAsync(client, &queries.GetClientDocumentContentUrl, "GetClientDocumentContentUrl", variables)

	variables = map[string]interface{}{
		"id": graphql.ID("your_client_document_id_here"),
	}
	sendAsync(client, &queries.GetProductDocumentContentUrl, "GetProductDocumentContentUrl", variables)

	variables = map[string]interface{}{
		"clientCompany": queries.GetCompany.ClientCompany.Id,
		"page":          graphql.Int(1),
	}
	sendAsync(client, &queries.GetClientDocumentCategories, "GetClientDocumentCategories", variables)

	variables = map[string]interface{}{
		"clientCompany":  graphql.String("your_client_company_id_here"),
		"id":             graphql.String("your_client_document_category_id_here"),
		"documentStatus": graphql.String(""),
		"page":           graphql.Int(1),
		"updatedAtOrder": graphql.String("DESC"), // Change to "ASC" if needed
		"itemsPerPage":   graphql.Int(10),        // Change as needed
	}
	sendAsync(client, &queries.GetClientCategoryDocuments, "GetClientCategoryDocuments", variables)

	variables = map[string]interface{}{
		"id":             graphql.String("your_product_document_category_id_here"),
		"documentStatus": graphql.String(""),
		"page":           graphql.Int(1),
	}
	sendAsync(client, &queries.GetProductCategoryDocuments, "GetProductCategoryDocuments", variables)

	variables = map[string]interface{}{
		"id":             graphql.String("your_product_document_category_id_here"),
		"documentStatus": graphql.String(""),
		"page":           graphql.Int(1),
	}
	sendAsync(client, &queries.GetProductCategoryDocumentsForCR, "GetProductCategoryDocumentsForCR", variables)

	variables = map[string]interface{}{
		"documentStatus": graphql.String(""),
		"page":           graphql.Int(1),
		"itemsPerPage":   graphql.Int(10),        // Change as needed
		"createdAtOrder": graphql.String("DESC"), // Change to "ASC" if needed
		"client":         graphql.String(""),
		"clientCompany":  graphql.String(""),
	}
	sendAsync(client, &queries.GetAllClientsDocuments, "GetAllClientsDocuments", variables)

	variables = map[string]interface{}{
		"id": graphql.String("your_account_id_here"),
	}
	sendAsync(client, &queries.GetAdditionalSAData, "GetAdditionalSAData", variables)

	variables = map[string]interface{}{
		"product":      graphql.String(""),
		"status_list":  []graphql.String{},
		"itemsPerPage": graphql.Int(500),
		"client":       graphql.String(""),
	}
	sendAsync(client, &queries.GetSharedWithCompanies, "GetSharedWithCompanies", variables)

	variables = map[string]interface{}{
		"client":       graphql.String(""),
		"product":      graphql.String(""),
		"status_list":  []graphql.String{},
		"itemsPerPage": graphql.Int(500),
		"module":       graphql.String(""),
		"page":         graphql.Int(1),
		"search":       graphql.String(""),
	}
	sendAsync(client, &queries.GetClientViewCompanies, "GetClientViewCompanies", variables)

	variables = map[string]interface{}{
		"itemsPerPage":   graphql.Int(15),
		"documentStatus": graphql.String(""),
		"page":           graphql.Int(1),
		"createdAtOrder": graphql.String("DESC"),
	}
	sendAsync(client, &queries.GetAllProductDocuments, "GetAllProductDocuments", variables)

	variables = map[string]interface{}{
		"productId":      graphql.String(""),
		"documentStatus": graphql.String(""),
		"page":           graphql.Int(1),
		"itemsPerPage":   graphql.Int(15),
		"updatedAtOrder": graphql.String("DESC"),
	}
	sendAsync(client, &queries.GetProductCategoriesForCR, "GetProductCategoriesForCR", variables)

	variables = map[string]interface{}{
		"product":      graphql.String(""),
		"status_list":  []graphql.String{},
		"itemsPerPage": graphql.Int(500),
		"search":       graphql.String(""),
		"page":         graphql.Int(1),
		"client":       graphql.String(""),
	}
	sendAsync(client, &queries.GetClientName, "GetClientName", variables)

	variables = map[string]interface{}{
		"clientCompany":  graphql.String(""),
		"documentStatus": graphql.String(""),
		"updatedAtOrder": graphql.String("DESC"),
		"itemsPerPage":   graphql.Int(5),
	}
	sendAsync(client, &queries.GetClientDocumentAllCategories, "GetClientDocumentAllCategories", variables)

	variables = map[string]interface{}{
		"clientCompany":  graphql.String(""),
		"documentStatus": graphql.String(""),
		"updatedAtOrder": graphql.String("DESC"),
		"itemsPerPage":   graphql.Int(5),
	}
	sendAsync(client, &queries.GetAllMyDocuments, "GetAllMyDocuments", variables)

	variables = map[string]interface{}{
		"documentStatus": graphql.String(""),
		"page":           graphql.Int(1),
		"deletedAtOrder": graphql.String("DESC"),
	}
	sendAsync(client, &queries.GetAllClientsDeletedDocuments, "GetAllClientsDeletedDocuments", variables)

	sendAsync(client, &queries.GetNotifyModules, "GetNotifyModules", nil)

	variables = map[string]interface{}{
		"documentStatus": graphql.String(""),
		"page":           graphql.Int(1),
		"deletedAtOrder": graphql.String("DESC"),
	}
	sendAsync(client, &queries.GetDeletedProduct, "GetDeletedProduct", variables)

	variables = map[string]interface{}{
		"status_list":  []string{},
		"page":         graphql.Int(1),
		"itemsPerPage": graphql.Int(10),
		"search":       graphql.String(""),
		"client":       graphql.String(""),
	}
	sendAsync(client, &queries.GetAllClientCompanies, "GetAllClientCompanies", variables)

	variables = map[string]interface{}{
		"id": graphql.ID("your_client_company_id"),
	}
	sendAsync(client, &queries.GetClientDashboardGroups, "GetClientDashboardGroups", variables)
	sendAsync(client, &queries.GetDashboardGroups, "GetDashboardGroups", nil)

	variables = map[string]interface{}{
		"status":       graphql.String(""),
		"page":         graphql.Int(1),
		"itemsPerPage": graphql.Int(10), // Change this value as needed
	}
	sendAsync(client, &queries.GetProductAccessRequests, "GetProductAccessRequests", variables)

	variables = map[string]interface{}{
		"id": "YOUR_BLOG_POST_ID_HERE",
	}
	sendAsync(client, &queries.GetSingleBlogPost, "GetSingleBlogPost", variables)

	variables = map[string]interface{}{
		"product_id":       graphql.String("YOUR_PRODUCT_ID_HERE"),
		"status":           graphql.String("YOUR_STATUS_HERE"),
		"page":             graphql.Int(1),
		"itemsPerPage":     graphql.Int(10),
		"publishedAtOrder": graphql.String("DESC"),
		"createdAtOrder":   graphql.String("YOUR_CREATED_AT_ORDER_HERE"),
	}
	sendAsync(client, &queries.GetBlogPosts, "GetBlogPosts", variables)

	variables = map[string]interface{}{
		"product_id":   graphql.String("YOUR_PRODUCT_ID_HERE"),
		"status":       graphql.String("YOUR_STATUS_HERE"),
		"page":         graphql.Int(1),
		"itemsPerPage": graphql.Int(10),
	}
	sendAsync(client, &queries.GetBlogPostsAsClient, "GetBlogPostsAsClient", variables)

	variables = map[string]interface{}{
		"id": "YOUR_BLOG_POST_ID_HERE",
	}
	sendAsync(client, &queries.GetSingleBlogPostAsClient, "GetSingleBlogPostAsClient", variables)

	variables = map[string]interface{}{
		"page":      graphql.Int(1),
		"eventDate": graphql.String("your_event_date_here"),
		"perPage":   graphql.Int(10),
		"dateRange": graphql.String(""),
	}
	sendAsync(client, &queries.CalendarEvents, "CalendarEvents", variables)

	variables = map[string]interface{}{
		"page":      graphql.Int(1),
		"eventDate": graphql.String("your_event_date_here"),
		"perPage":   graphql.Int(10),
		"dateRange": graphql.String(""),
	}
	sendAsync(client, &queries.CalendarEventsForCR, "CalendarEventsForCR", variables)

	variables = map[string]interface{}{
		"eventDate": graphql.String("your_event_date_here"),
		"page":      graphql.Int(1),
	}
	sendAsync(client, &queries.AppEvents, "AppEvents ", variables)

	variables = map[string]interface{}{
		"eventDate": graphql.String("your_event_date_here"),
		"page":      graphql.Int(1),
	}
	sendAsync(client, &queries.AppEventsForCR, "AppEventsForCR", variables)

	variables = map[string]interface{}{
		"participants_firstName": graphql.String("first_name_here"),
		"participants_lastName":  graphql.String("last_name_here"),
		"last":                   graphql.Int(10), // or any desired value
	}
	sendAsync(client, &queries.GetChatsRooms, "GetChatsRooms", variables)

	variables = map[string]interface{}{
		"chatRoom":     graphql.String("your_chat_room_id_here"),
		"page":         graphql.Int(1),
		"itemsPerPage": graphql.Int(10),
	}
	sendAsync(client, &queries.GetChatsMessages, "GetChatsMessages", variables)

	variables = map[string]interface{}{
		"status": graphql.String("your_status_here"),
		"type":   graphql.String("your_type_here"),
		"search": graphql.String("your_search_here"),
		"sustainabilityAdvisorProducts_product_id": graphql.String("your_product_id_here"),
	}
	sendAsync(client, &queries.GetUsersChat, "GetUsersChat", variables)

	variables = map[string]interface{}{
		"page":       graphql.Int(1),
		"company_id": graphql.String("your_company_id_here"),
		"productId":  graphql.String("your_product_id_here"),
	}
	sendAsync(client, &queries.GetAllYears, "GetAllYears", variables)

	variables = map[string]interface{}{
		"products_id":      graphql.String("your_products_id_here"),
		"status_list":      []graphql.String{"your_status_list_here"},
		"status":           graphql.String("your_status_here"),
		"page":             graphql.Int(1),
		"itemsPerPage":     graphql.Int(10),
		"publishedAtOrder": graphql.String("DESC"),
		"createdAtOrder":   graphql.String(""), // Specify the order if needed, e.g., "ASC" or "DESC"
	}
	sendAsync(client, &queries.GetNewsletters, "GetNewsletters", variables)

	wg.Wait()
}

func sendAsync(client *graphql.Client, q interface{}, name string, v map[string]interface{}) {
	wg.Add(1)
	go sendReq(client, q, name, v, true)
}

func sendReq(client *graphql.Client, q interface{}, name string, v map[string]interface{}, async bool) {
	if async {
		defer wg.Done()
	}

	startTime := time.Now().UnixMilli()

	err := client.Query(context.Background(), q, v)
	if err != nil {
		panic(err.Error())
	}

	endTime := time.Now().UnixMilli()
	elapsedTime := endTime - startTime

	fmt.Printf("Request %s took: %d ms\n", name, elapsedTime)
}
