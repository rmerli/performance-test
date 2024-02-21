package queries

import "github.com/hasura/go-graphql-client"

var GetCalendar struct {
	Account struct {
		Id                        graphql.ID      `graphql:"id"`
		CalendarId                graphql.String  `graphql:"calendarId"`
		CalendlyId                graphql.String  `graphql:"calendlyId"`
		IsConnectedGoogleCalendar graphql.Boolean `graphql:"isConnectedGoogleCalendar"`
	} `graphql:"account(id: $id)"`
}
