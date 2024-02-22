package main

import (
	"context"
	"fmt"
	"net/http"
	"perfomance/test/mutations"
	"perfomance/test/mutations/inputs"
	"perfomance/test/queries"
	"perfomance/test/requests"
	"time"

	"github.com/hasura/go-graphql-client"
)

const url = "https://alpha-portal-api.afsenergy.nl/api/graphql"

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
	portalClient := requests.Client{GqlClient: client}

	portalClient.SendReq(&queries.GetClientCompanies, "GetCompanies", nil, false)
	portalClient.SendReq(&queries.GetProducts, "GetProducts", nil, false)

	variables = map[string]interface{}{
		"id":           queries.GetClientCompanies.ClientCompanies.Collection[len(queries.GetClientCompanies.ClientCompanies.Collection)-1].Id,
		"page":         graphql.Int(1),
		"itemsPerPage": graphql.Int(100),
	}
	portalClient.SendReq(&queries.GetCompany, "GetCompany", variables, false)

	portalClient.SendAccountRequests()
	portalClient.SendCompanyRequests()
	portalClient.SendDocumentRequests()
	portalClient.SendModuleRequests()
	portalClient.SendCalendarRequests()
	portalClient.SendChatRequests()
	portalClient.SendCommodityRequests()
	portalClient.Wg.Wait()
}
