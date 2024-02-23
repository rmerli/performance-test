package requests

import (
	"perfomance/test/queries"

	"github.com/hasura/go-graphql-client"
)

func (c Client) SendModuleRequests() {
	c.SendAsync(&queries.GetModules, "GetModules", nil)
	c.SendAsync(&queries.GetModulesWithGroups, "GetModulesWithGroups", nil)

	variables := map[string]interface{}{
		"id": graphql.ID(queries.GetClientCompanies.ClientCompanies.Collection[len(queries.GetClientCompanies.ClientCompanies.Collection)-1].Id),
	}
	c.SendAsync(&queries.GetClientDashboardGroups, "GetClientDashboardGroups", variables)
	c.SendAsync(&queries.GetDashboardGroups, "GetDashboardGroups", nil)
}
