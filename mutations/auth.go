package mutations

import (
	"github.com/hasura/go-graphql-client"
)

var StartAuth struct {
	StartAuthentication struct {
		Authentication struct {
			AccessToken     graphql.String
			QrCode          graphql.String
			MercureToken    graphql.String
			Is2FaInProgress graphql.Boolean `graphql:"is2FaInProgress"`
			Account         struct {
				Avatar struct {
					Id         graphql.ID
					ContentUrl graphql.String
					CreatedAt  graphql.String
					FileName   graphql.String
					FilePath   graphql.String
					UpdatedAt  graphql.String
				}
				Id            graphql.ID
				FirstName     graphql.String
				MiddleName    graphql.String
				LastName      graphql.String
				CreatedAt     graphql.String
				Email         graphql.String
				Roles         []graphql.String
				ClientCompany struct {
					Id        graphql.ID
					LegalName graphql.String
				}
			}
		}
	} `graphql:"startAuthentication(input: $input)"`
}
