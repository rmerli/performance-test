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

var CalendarEvents struct {
	CalendarEvents struct {
		PaginationInfo struct {
			ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
			LastPage     graphql.Int `graphql:"lastPage"`
			TotalCount   graphql.Int `graphql:"totalCount"`
		}
		Collection []struct {
			Id                        graphql.ID     `graphql:"id"`
			Summary                   graphql.String `graphql:"summary"`
			Description               graphql.String `graphql:"description"`
			StartDateTime             graphql.Int    `graphql:"startDateTime"`
			EndDateTime               graphql.Int    `graphql:"endDateTime"`
			GoogleEventLink           graphql.String `graphql:"googleEventLink"`
			GoogleCancelEventLink     graphql.String `graphql:"googleCancelEventLink"`
			GoogleRescheduleEventLink graphql.String `graphql:"googleRescheduleEventLink"`
			Organizer                 struct {
				Id    graphql.ID `graphql:"id"`
				Roles []string   `graphql:"roles"`
			} `graphql:"organizer"`
			Attendees struct {
				Collection []struct {
					Id         graphql.ID     `graphql:"id"`
					FirstName  graphql.String `graphql:"firstName"`
					MiddleName graphql.String `graphql:"middleName"`
					LastName   graphql.String `graphql:"lastName"`
				}
			} `graphql:"attendees(itemsPerPage: 50)"`
		}
	} `graphql:"calendarEvents(page: $page, itemsPerPage: $perPage)"`
}

var CalendarEventsForCR struct {
	CalendarEvents struct {
		PaginationInfo struct {
			ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
			LastPage     graphql.Int `graphql:"lastPage"`
			TotalCount   graphql.Int `graphql:"totalCount"`
		}
		Collection []struct {
			Id                        graphql.ID     `graphql:"id"`
			Summary                   graphql.String `graphql:"summary"`
			Description               graphql.String `graphql:"description"`
			StartDateTime             graphql.String `graphql:"startDateTime"`
			EndDateTime               graphql.String `graphql:"endDateTime"`
			GoogleEventLink           graphql.String `graphql:"googleEventLink"`
			GoogleCancelEventLink     graphql.String `graphql:"googleCancelEventLink"`
			GoogleRescheduleEventLink graphql.String `graphql:"googleRescheduleEventLink"`
			Organizer                 struct {
				Id graphql.ID `graphql:"id"`
			} `graphql:"organizer"`
			Attendees struct {
				Collection []struct {
					Id         graphql.ID     `graphql:"id"`
					FirstName  graphql.String `graphql:"firstName"`
					MiddleName graphql.String `graphql:"middleName"`
					LastName   graphql.String `graphql:"lastName"`
				}
			} `graphql:"attendees(itemsPerPage: 50)"`
		}
	} `graphql:"calendarEvents(page: $page, eventDate: $eventDate, itemsPerPage: $perPage, dateRange: $dateRange)"`
}

var AppEvents struct {
	InAppEvents struct {
		PaginationInfo struct {
			ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
			LastPage     graphql.Int `graphql:"lastPage"`
			TotalCount   graphql.Int `graphql:"totalCount"`
		}
		Collection []struct {
			Id            graphql.ID     `graphql:"id"`
			Title         graphql.String `graphql:"title"`
			Description   graphql.String `graphql:"description"`
			StartDateTime graphql.Int    `graphql:"startDateTime"`
			EndDateTime   graphql.Int    `graphql:"endDateTime"`
			CreatedBy     struct {
				Id         graphql.ID     `graphql:"id"`
				FirstName  graphql.String `graphql:"firstName"`
				MiddleName graphql.String `graphql:"middleName"`
				LastName   graphql.String `graphql:"lastName"`
			} `graphql:"createdBy"`
			Product struct {
				Id           graphql.ID     `graphql:"id"`
				Name         graphql.String `graphql:"name"`
				Abbreviation graphql.String `graphql:"abbreviation"`
			}
			CanBeDeleted graphql.Boolean `graphql:"canBeDeleted"`
			CanBeEdited  graphql.Boolean `graphql:"canBeEdited"`
			SharedWith   struct {
				Collection []struct {
					Id        graphql.ID     `graphql:"id"`
					LegalName graphql.String `graphql:"legalName"`
				}
			} `graphql:"sharedWith"`
		}
	} `graphql:"inAppEvents(page: $page)"`
}

var AppEventsForCR struct {
	InAppEvents struct {
		PaginationInfo struct {
			ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
			LastPage     graphql.Int `graphql:"lastPage"`
			TotalCount   graphql.Int `graphql:"totalCount"`
		}
		Collection []struct {
			Id            graphql.ID     `graphql:"id"`
			Title         graphql.String `graphql:"title"`
			Description   graphql.String `graphql:"description"`
			StartDateTime graphql.String `graphql:"startDateTime"`
			EndDateTime   graphql.String `graphql:"endDateTime"`
			CreatedBy     struct {
				Id         graphql.ID     `graphql:"id"`
				FirstName  graphql.String `graphql:"firstName"`
				MiddleName graphql.String `graphql:"middleName"`
				LastName   graphql.String `graphql:"lastName"`
			} `graphql:"createdBy"`
			Product struct {
				Id           graphql.ID     `graphql:"id"`
				Name         graphql.String `graphql:"name"`
				Abbreviation graphql.String `graphql:"abbreviation"`
			}
			CanBeDeleted graphql.Boolean `graphql:"canBeDeleted"`
			CanBeEdited  graphql.Boolean `graphql:"canBeEdited"`
		}
	} `graphql:"inAppEvents(eventDate: $eventDate, page: $page)"`
}
