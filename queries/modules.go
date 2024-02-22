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

var GetClientDashboardGroups struct {
	ClientCompany struct {
		Id        graphql.ID     `graphql:"id"`
		LegalName graphql.String `graphql:"legalName"`
		Status    graphql.String `graphql:"status"`
		Products  struct {
			Edges []struct {
				Node struct {
					DefaultDashboardGroup struct {
						Id      graphql.ID     `graphql:"id"`
						Name    graphql.String `graphql:"name"`
						Modules struct {
							Edges []struct {
								Node struct {
									Name graphql.String `graphql:"name"`
								} `graphql:"node"`
							} `graphql:"edges"`
						} `graphql:"modules"`
					} `graphql:"defaultDashboardGroup"`
				} `graphql:"node"`
			} `graphql:"edges"`
		} `graphql:"products"`
	} `graphql:"clientCompany(id: $id)"`
}

var GetDashboardGroups struct {
	DashboardGroups struct {
		Edges []struct {
			Node struct {
				Id        graphql.ID     `graphql:"id"`
				Name      graphql.String `graphql:"name"`
				Slug      graphql.String `graphql:"slug"`
				DeletedAt graphql.String `graphql:"deletedAt"`
				Modules   struct {
					Edges []struct {
						Node struct {
							Id   graphql.ID     `graphql:"id"`
							Name graphql.String `graphql:"name"`
						} `graphql:"node"`
					} `graphql:"edges"`
				} `graphql:"modules"`
			} `graphql:"node"`
		} `graphql:"edges"`
	} `graphql:"dashboardGroups"`
}
