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
