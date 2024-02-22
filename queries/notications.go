package queries

import "github.com/hasura/go-graphql-client"

var GetNotifyModules struct {
	AccountNotificationSettings struct {
		Edges []struct {
			Node struct {
				Id        graphql.ID      `graphql:"id"`
				Module    graphql.String  `graphql:"module"`
				SendEmail graphql.Boolean `graphql:"sendEmail"`
				SendPush  graphql.Boolean `graphql:"sendPush"`
			} `graphql:"node"`
		} `graphql:"edges"`
	} `graphql:"accountNotificationSettings"`
}
