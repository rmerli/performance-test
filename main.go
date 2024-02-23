package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"perfomance/test/mutations"
	"perfomance/test/mutations/inputs"
	"perfomance/test/requests"
	"sync"
	"time"

	"github.com/hasura/go-graphql-client"
)

const url = "https://alpha-portal-api.afsenergy.nl/api/graphql"

func main() {
	wg := sync.WaitGroup{}
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
	csvFile, err := os.Create(fmt.Sprintf("async_request_times_%d.csv", time.Now().Unix()))
	if err != nil {
		fmt.Printf("failed creating file: %s", err)
	}

	csvWriter := csv.NewWriter(csvFile)
	csvWriter.Write([]string{"Request", "Time ms"})

	//Authenticated client
	client = graphql.NewClient(url, nil).WithRequestModifier(func(r *http.Request) {
		r.Header.Add("Authorization", token)
	})
	portalClient := requests.Client{GqlClient: client, Wg: &wg, CsvWriter: csvWriter}

	portalClient.SendSyncRequests()
	portalClient.SendAccountRequests()
	portalClient.SendCompanyRequests()
	portalClient.SendDocumentRequests()
	portalClient.SendModuleRequests()
	portalClient.SendCalendarRequests()
	portalClient.SendCommodityRequests()

	portalClient.Wg.Wait()
	csvWriter.Flush()
	csvFile.Close()
}
