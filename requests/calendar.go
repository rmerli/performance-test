package requests

import (
	"perfomance/test/mutations"
	"perfomance/test/queries"

	"github.com/hasura/go-graphql-client"
)

func (c Client) SendCalendarRequests() {
	variables := map[string]interface{}{
		"id": mutations.StartAuth.StartAuthentication.Authentication.Account.Id,
	}
	c.SendAsync(&queries.GetCalendar, "GetCalendar", variables)

	variables = map[string]interface{}{
		"page":    graphql.Int(1),
		"perPage": graphql.Int(10),
	}
	c.SendAsync(&queries.CalendarEvents, "CalendarEvents", variables)

	variables = map[string]interface{}{
		"page": graphql.Int(1),
	}
	c.SendAsync(&queries.AppEvents, "AppEvents ", variables)
}
