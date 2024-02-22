package requests

import (
	"perfomance/test/queries"

	"github.com/hasura/go-graphql-client"
)

func (c Client) SendChatRequests() {
	variables := map[string]interface{}{
		"participants_firstName": graphql.String("first_name_here"),
		"participants_lastName":  graphql.String("last_name_here"),
		"last":                   graphql.Int(10), // or any desired value
	}
	c.SendAsync(&queries.GetChatsRooms, "GetChatsRooms", variables)

	variables = map[string]interface{}{
		"chatRoom":     graphql.String("your_chat_room_id_here"),
		"page":         graphql.Int(1),
		"itemsPerPage": graphql.Int(10),
	}
	c.SendAsync(&queries.GetChatsMessages, "GetChatsMessages", variables)

	variables = map[string]interface{}{
		"status": graphql.String("your_status_here"),
		"type":   graphql.String("your_type_here"),
		"search": graphql.String("your_search_here"),
		"sustainabilityAdvisorProducts_product_id": graphql.String("your_product_id_here"),
	}
	c.SendAsync(&queries.GetUsersChat, "GetUsersChat", variables)

}
