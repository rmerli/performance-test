package requests

import "perfomance/test/queries"

func (c Client) SendNotificationRequests() {
	c.SendAsync(&queries.GetNotifyModules, "GetNotifyModules", nil)
}
