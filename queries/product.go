package queries

import "github.com/hasura/go-graphql-client"

var GetProduct struct {
	Products struct {
		Edges []struct {
			Node struct {
				Id           graphql.ID
				Name         graphql.String
				Abbreviation graphql.String
				Modules      struct {
					Edges []struct {
						Node struct {
							Id   graphql.ID
							Name graphql.String
							Slug graphql.String
						}
					}
				}
			}
		}
	}
}
