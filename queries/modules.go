package queries

import "github.com/hasura/go-graphql-client"

var GetModules struct {
	Modules struct {
		Edges []struct {
			Node struct {
				Id   graphql.ID     `graphql:"id"`
				Name graphql.String `graphql:"name"`
				Slug graphql.String `graphql:"slug"`
			}
		}
	}
}

var GetModulesWithGroups struct {
	Modules struct {
		Edges []struct {
			Node struct {
				Id              graphql.ID     `graphql:"id"`
				Name            graphql.String `graphql:"name"`
				Slug            graphql.String `graphql:"slug"`
				DashboardGroups struct {
					Edges []struct {
						Node struct {
							Id   graphql.ID     `graphql:"id"`
							Name graphql.String `graphql:"name"`
							Slug graphql.String `graphql:"slug"`
						}
					}
				} `graphql:"dashboardGroups"`
			}
		}
	}
}
