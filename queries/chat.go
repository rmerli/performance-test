package queries

import "github.com/hasura/go-graphql-client"

var GetChatsRooms struct {
	ChatRooms struct {
		TotalCount graphql.Int `graphql:"totalCount"`
		Edges      []struct {
			Node struct {
				Id                 graphql.ID     `graphql:"id"`
				Name               graphql.String `graphql:"name"`
				UnreadCount        graphql.Int    `graphql:"unreadCount"`
				LastMessageContent graphql.String `graphql:"lastMessageContent"`
				LastMessageTime    graphql.String `graphql:"lastMessageTime"`
				ClientCompany      struct {
					Id        graphql.ID     `graphql:"id"`
					LegalName graphql.String `graphql:"legalName"`
				} `graphql:"clientCompany"`
				Product struct {
					Id                     graphql.ID     `graphql:"id"`
					Abbreviation           graphql.String `graphql:"abbreviation"`
					Name                   graphql.String `graphql:"name"`
					SustainabilityAdvisors struct {
						Edges []struct {
							Node struct {
								Id               graphql.ID      `graphql:"id"`
								IsProductManager graphql.Boolean `graphql:"isProductManager"`
								Account          struct {
									Id graphql.ID `graphql:"id"`
								} `graphql:"account"`
								Product struct {
									Id   graphql.ID     `graphql:"id"`
									Name graphql.String `graphql:"name"`
								} `graphql:"product"`
							}
						}
					} `graphql:"sustainabilityAdvisors"`
				} `graphql:"product"`
				Participants struct {
					Collection []struct {
						Avatar                        graphql.ID     `graphql:"avatar"`
						Id                            graphql.ID     `graphql:"id"`
						FirstName                     graphql.String `graphql:"firstName"`
						LastName                      graphql.String `graphql:"lastName"`
						MiddleName                    graphql.String `graphql:"middleName"`
						Email                         graphql.String `graphql:"email"`
						Roles                         []string       `graphql:"roles"`
						SustainabilityAdvisorProducts struct {
							Edges []struct {
								Node struct {
									Id      graphql.ID `graphql:"id"`
									Account struct {
										Id         graphql.ID     `graphql:"id"`
										FirstName  graphql.String `graphql:"firstName"`
										LastName   graphql.String `graphql:"lastName"`
										MiddleName graphql.String `graphql:"middleName"`
									} `graphql:"account"`
								}
							}
						} `graphql:"sustainabilityAdvisorProducts"`
						ClientCompany struct {
							Id        graphql.ID     `graphql:"id"`
							LegalName graphql.String `graphql:"legalName"`
						} `graphql:"clientCompany"`
					}
				} `graphql:"participants(itemsPerPage: 200)"`
			}
		}
	} `graphql:"chatRooms(participants_firstName: $participants_firstName, participants_lastName: $participants_lastName, last: $last)"`
}

var GetChatsMessages struct {
	ChatMessages struct {
		PaginationInfo struct {
			ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
			LastPage     graphql.Int `graphql:"lastPage"`
			TotalCount   graphql.Int `graphql:"totalCount"`
		}
		Collection []struct {
			Id      graphql.ID     `graphql:"id"`
			Content graphql.String `graphql:"content"`
			Sender  struct {
				Email      graphql.String `graphql:"email"`
				Id         graphql.ID     `graphql:"id"`
				FirstName  graphql.String `graphql:"firstName"`
				LastName   graphql.String `graphql:"lastName"`
				MiddleName graphql.String `graphql:"middleName"`
				Avatar     struct {
					Id         graphql.ID     `graphql:"id"`
					ContentUrl graphql.String `graphql:"contentUrl"`
					CreatedAt  graphql.String `graphql:"createdAt"`
					FileName   graphql.String `graphql:"fileName"`
					FilePath   graphql.String `graphql:"filePath"`
					UpdatedAt  graphql.String `graphql:"updatedAt"`
				}
			} `graphql:"sender"`
			SentAt   graphql.String `graphql:"sentAt"`
			Document struct {
				Id         graphql.ID     `graphql:"id"`
				Title      graphql.String `graphql:"title"`
				Note       graphql.String `graphql:"note"`
				ContentUrl graphql.String `graphql:"contentUrl"`
			} `graphql:"document"`
		}
	} `graphql:"chatMessages(chatRoom: $chatRoom, page: $page, itemsPerPage: $itemsPerPage)"`
}

var GetUsersChat struct {
	Accounts struct {
		PaginationInfo struct {
			ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
			LastPage     graphql.Int `graphql:"lastPage"`
			TotalCount   graphql.Int `graphql:"totalCount"`
		}
		Collection []struct {
			Id                            graphql.ID     `graphql:"id"`
			FirstName                     graphql.String `graphql:"firstName"`
			LastName                      graphql.String `graphql:"lastName"`
			MiddleName                    graphql.String `graphql:"middleName"`
			Email                         graphql.String `graphql:"email"`
			DeletedAt                     graphql.String `graphql:"deletedAt"`
			Status                        graphql.String `graphql:"status"`
			Roles                         []string       `graphql:"roles"`
			SustainabilityAdvisorProducts struct {
				Edges []struct {
					Node struct {
						Id               graphql.ID      `graphql:"id"`
						IsProductManager graphql.Boolean `graphql:"isProductManager"`
					}
				}
			} `graphql:"sustainabilityAdvisorProducts"`
		}
	} `graphql:"accounts(status: $status, type: $type, page: 1, itemsPerPage: 500, search: $search, sustainabilityAdvisorProducts_product_id: $sustainabilityAdvisorProducts_product_id)"`
}
