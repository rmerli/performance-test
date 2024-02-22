package queries

import "github.com/hasura/go-graphql-client"

var GetSingleBlogPost struct {
	BlogPost struct {
		Id          graphql.ID     `graphql:"id"`
		Title       graphql.String `graphql:"title"`
		Body        graphql.String `graphql:"body"`
		Description graphql.String `graphql:"description"`
		Product     struct {
			Id           graphql.ID     `graphql:"id"`
			Abbreviation graphql.String `graphql:"abbreviation"`
		} `graphql:"product"`
		SharedWith struct {
			Collection []struct {
				Id        graphql.ID     `graphql:"id"`
				LegalName graphql.String `graphql:"legalName"`
			} `graphql:"collection"`
		} `graphql:"sharedWith"`
		CreatedAt graphql.String `graphql:"createdAt"`
		CreatedBy struct {
			Avatar struct {
				Id         graphql.ID     `graphql:"id"`
				ContentUrl graphql.String `graphql:"contentUrl"`
				CreatedAt  graphql.String `graphql:"createdAt"`
				FileName   graphql.String `graphql:"fileName"`
				FilePath   graphql.String `graphql:"filePath"`
				UpdatedAt  graphql.String `graphql:"updatedAt"`
			} `graphql:"avatar"`
			Id          graphql.ID     `graphql:"id"`
			FirstName   graphql.String `graphql:"firstName"`
			LastName    graphql.String `graphql:"lastName"`
			MiddleName  graphql.String `graphql:"middleName"`
			PhoneNumber graphql.String `graphql:"phoneNumber"`
			Email       graphql.String `graphql:"email"`
			Type        graphql.String `graphql:"type"`
		} `graphql:"createdBy"`
		Cover struct {
			ContentUrl graphql.String `graphql:"contentUrl"`
			FileName   graphql.String `graphql:"fileName"`
		} `graphql:"cover"`
		PublishedAt graphql.String `graphql:"publishedAt"`
		UpdatedAt   graphql.String `graphql:"updatedAt"`
		Status      graphql.String `graphql:"status"`
	} `graphql:"blogPost(id: $id)"`
}

var GetBlogPosts struct {
	BlogPosts struct {
		PaginationInfo struct {
			ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
			TotalCount   graphql.Int `graphql:"totalCount"`
			LastPage     graphql.Int `graphql:"lastPage"`
		}
		Collection []struct {
			Id          graphql.ID     `graphql:"id"`
			Title       graphql.String `graphql:"title"`
			Body        graphql.String `graphql:"body"`
			Description graphql.String `graphql:"description"`
			Product     struct {
				Id           graphql.ID     `graphql:"id"`
				Abbreviation graphql.String `graphql:"abbreviation"`
			}
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
	} `graphql:"blogPosts(product_id: $product_id, status: $status, order: { publishedAt: $publishedAtOrder, createdAt: $createdAtOrder }, page: $page, itemsPerPage: $itemsPerPage)"`
}

var GetBlogPostsAsClient struct {
	BlogPosts struct {
		PaginationInfo struct {
			ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
			TotalCount   graphql.Int `graphql:"totalCount"`
			LastPage     graphql.Int `graphql:"lastPage"`
		}
		Collection []struct {
			Id          graphql.ID     `graphql:"id"`
			Title       graphql.String `graphql:"title"`
			Body        graphql.String `graphql:"body"`
			Description graphql.String `graphql:"description"`
			Product     struct {
				Id           graphql.ID     `graphql:"id"`
				Abbreviation graphql.String `graphql:"abbreviation"`
			}
			Cover struct {
				ContentUrl graphql.String `graphql:"contentUrl"`
			}
			CreatedAt   graphql.String `graphql:"createdAt"`
			PublishedAt graphql.String `graphql:"publishedAt"`
			UpdatedAt   graphql.String `graphql:"updatedAt"`
			Status      graphql.String `graphql:"status"`
		}
	} `graphql:"blogPosts(product_id: $product_id, status: $status, order: { publishedAt: \"DESC\" }, page: $page, itemsPerPage: $itemsPerPage)"`
}

var GetSingleBlogPostAsClient struct {
	BlogPost struct {
		Id          graphql.ID     `graphql:"id"`
		Title       graphql.String `graphql:"title"`
		Body        graphql.String `graphql:"body"`
		Description graphql.String `graphql:"description"`
		Product     struct {
			Id           graphql.ID     `graphql:"id"`
			Abbreviation graphql.String `graphql:"abbreviation"`
		}
		CreatedAt graphql.String `graphql:"createdAt"`
		CreatedBy struct {
			Avatar struct {
				Id         graphql.ID     `graphql:"id"`
				ContentUrl graphql.String `graphql:"contentUrl"`
				CreatedAt  graphql.String `graphql:"createdAt"`
				FileName   graphql.String `graphql:"fileName"`
				FilePath   graphql.String `graphql:"filePath"`
				UpdatedAt  graphql.String `graphql:"updatedAt"`
			} `graphql:"avatar"`
			Id          graphql.ID     `graphql:"id"`
			FirstName   graphql.String `graphql:"firstName"`
			MiddleName  graphql.String `graphql:"middleName"`
			LastName    graphql.String `graphql:"lastName"`
			PhoneNumber graphql.String `graphql:"phoneNumber"`
			Email       graphql.String `graphql:"email"`
			Type        graphql.String `graphql:"type"`
		} `graphql:"createdBy"`
		Cover struct {
			ContentUrl graphql.String `graphql:"contentUrl"`
		}
		PublishedAt graphql.String `graphql:"publishedAt"`
		UpdatedAt   graphql.String `graphql:"updatedAt"`
		Status      graphql.String `graphql:"status"`
	} `graphql:"blogPost(id: $id)"`
}
