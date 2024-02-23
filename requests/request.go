package requests

import (
	"context"
	"encoding/csv"
	"fmt"
	"perfomance/test/queries"
	"sync"
	"time"

	"github.com/hasura/go-graphql-client"
)

type Client struct {
	GqlClient *graphql.Client
	Wg        *sync.WaitGroup
	CsvWriter *csv.Writer
}

func (c *Client) SendAsync(q interface{}, name string, v map[string]interface{}) {
	c.Wg.Add(1)
	go c.SendReq(q, name, v, true)
}

func (c *Client) SendReq(q interface{}, name string, v map[string]interface{}, async bool) {
	if async {
		defer c.Wg.Done()
	}

	startTime := time.Now().UnixMilli()

	err := c.GqlClient.Query(context.Background(), q, v)
	if err != nil {
		panic(err.Error())
	}

	endTime := time.Now().UnixMilli()

	elapsedTime := endTime - startTime

	c.CsvWriter.Write([]string{name, fmt.Sprintf("%d", elapsedTime)})

	fmt.Printf("Request %s took: %d ms\n", name, elapsedTime)
}

func (c *Client) SendSyncRequests() {
	c.SendReq(&queries.GetClientCompanies, "GetCompanies", nil, false)
	c.SendReq(&queries.GetProducts, "GetProducts", nil, false)

	variables := map[string]interface{}{
		"page":           graphql.Int(1),
		"itemsPerPage":   graphql.Int(10),
		"updatedAtOrder": graphql.String("DESC"), // Change to "ASC" if needed
		"createdAtOrder": graphql.String("DESC"), // Change to "ASC" if needed
	}
	c.SendReq(&queries.GetProductCategories, "GetProductCategories", variables, false)

	variables = map[string]interface{}{
		"id":           queries.GetClientCompanies.ClientCompanies.Collection[len(queries.GetClientCompanies.ClientCompanies.Collection)-1].Id,
		"page":         graphql.Int(1),
		"itemsPerPage": graphql.Int(100),
	}
	c.SendReq(&queries.GetCompany, "GetCompany", variables, false)

	variables = map[string]interface{}{
		"page":           graphql.Int(1),
		"itemsPerPage":   graphql.Int(10),        // Change as needed
		"createdAtOrder": graphql.String("DESC"), // Change to "ASC" if needed
	}
	c.SendReq(&queries.GetAllClientsDocuments, "GetAllClientsDocuments", variables, false)

	variables = map[string]interface{}{
		"itemsPerPage":   graphql.Int(15),
		"page":           graphql.Int(1),
		"createdAtOrder": graphql.String("DESC"),
	}
	c.SendReq(&queries.GetAllProductDocuments, "GetAllProductDocuments", variables, false)

	variables = map[string]interface{}{
		"clientCompany": graphql.String(queries.GetCompany.ClientCompany.Id),
	}
	c.SendReq(&queries.GetClientDocumentCategories, "GetClientDocumentCategories", variables, false)
}
