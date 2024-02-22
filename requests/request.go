package requests

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/hasura/go-graphql-client"
)

type Client struct {
	GqlClient *graphql.Client
	Wg        *sync.WaitGroup
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

	fmt.Printf("Request %s took: %d ms\n", name, elapsedTime)
}
