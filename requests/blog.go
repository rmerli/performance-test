package requests

import (
	"perfomance/test/queries"

	"github.com/hasura/go-graphql-client"
)

func (c Client) SendBlogRequests() {
	variables := map[string]interface{}{
		"id": "YOUR_BLOG_POST_ID_HERE",
	}
	c.SendAsync(&queries.GetSingleBlogPost, "GetSingleBlogPost", variables)

	variables = map[string]interface{}{
		"product_id":       graphql.String("YOUR_PRODUCT_ID_HERE"),
		"status":           graphql.String("YOUR_STATUS_HERE"),
		"page":             graphql.Int(1),
		"itemsPerPage":     graphql.Int(10),
		"publishedAtOrder": graphql.String("DESC"),
		"createdAtOrder":   graphql.String("YOUR_CREATED_AT_ORDER_HERE"),
	}
	c.SendAsync(&queries.GetBlogPosts, "GetBlogPosts", variables)

	variables = map[string]interface{}{
		"product_id":   graphql.String("YOUR_PRODUCT_ID_HERE"),
		"status":       graphql.String("YOUR_STATUS_HERE"),
		"page":         graphql.Int(1),
		"itemsPerPage": graphql.Int(10),
	}
	c.SendAsync(&queries.GetBlogPostsAsClient, "GetBlogPostsAsClient", variables)

	variables = map[string]interface{}{
		"id": "YOUR_BLOG_POST_ID_HERE",
	}
	c.SendAsync(&queries.GetSingleBlogPostAsClient, "GetSingleBlogPostAsClient", variables)

	variables = map[string]interface{}{
		"products_id":      graphql.String("your_products_id_here"),
		"status_list":      []graphql.String{"your_status_list_here"},
		"status":           graphql.String("your_status_here"),
		"page":             graphql.Int(1),
		"itemsPerPage":     graphql.Int(10),
		"publishedAtOrder": graphql.String("DESC"),
		"createdAtOrder":   graphql.String(""), // Specify the order if needed, e.g., "ASC" or "DESC"
	}
	c.SendAsync(&queries.GetNewsletters, "GetNewsletters", variables)
}
