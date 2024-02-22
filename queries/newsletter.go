package queries

import "github.com/hasura/go-graphql-client"

var GetNewsletters struct {
	Newsletters struct {
		PaginationInfo struct {
			LastPage graphql.Int `graphql:"lastPage"`
		}
		Collection []struct {
			Id          graphql.ID     `graphql:"id"`
			Title       graphql.String `graphql:"title"`
			Body        graphql.String `graphql:"body"`
			Description graphql.String `graphql:"description"`
			Products    struct {
				Edges []struct {
					Node struct {
						Id           graphql.ID     `graphql:"id"`
						Abbreviation graphql.String `graphql:"abbreviation"`
					}
				}
			} `graphql:"products"`
			SharedWith struct {
				Collection []struct {
					Id        graphql.ID     `graphql:"id"`
					LegalName graphql.String `graphql:"legalName"`
				}
			} `graphql:"sharedWith"`
			Cover struct {
				ContentUrl graphql.String `graphql:"contentUrl"`
			}
			CreatedAt graphql.String `graphql:"createdAt"`
			CreatedBy struct {
				Id graphql.ID `graphql:"id"`
			} `graphql:"createdBy"`
			PublishedAt graphql.String `graphql:"publishedAt"`
			UpdatedAt   graphql.String `graphql:"updatedAt"`
			Status      graphql.String `graphql:"status"`
		}
	} `graphql:"newsletters(products_id: $products_id, status_list: $status_list, itemsPerPage: $itemsPerPage, status: $status, page: $page, order: { publishedAt: $publishedAtOrder, createdAt: $createdAtOrder })"`
}
