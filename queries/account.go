package queries

import "github.com/hasura/go-graphql-client"

var GetUsers struct {
	Accounts struct {
		PaginationInfo struct {
			ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
			LastPage     graphql.Int `graphql:"lastPage"`
			TotalCount   graphql.Int `graphql:"totalCount"`
		} `graphql:"paginationInfo"`
		Collection []struct {
			Id                            graphql.ID       `graphql:"id"`
			Email                         graphql.String   `graphql:"email"`
			DeletedAt                     graphql.String   `graphql:"deletedAt"`
			Status                        graphql.String   `graphql:"status"`
			Roles                         []graphql.String `graphql:"roles"`
			PhoneNumber                   graphql.String   `graphql:"phoneNumber"`
			FirstName                     graphql.String   `graphql:"firstName"`
			LastName                      graphql.String   `graphql:"lastName"`
			MiddleName                    graphql.String   `graphql:"middleName"`
			SustainabilityAdvisorProducts struct {
				Edges []struct {
					Node struct {
						Id               graphql.ID      `graphql:"id"`
						IsProductManager graphql.Boolean `graphql:"isProductManager"`
						Account          struct {
							Id         graphql.ID     `graphql:"id"`
							FirstName  graphql.String `graphql:"firstName"`
							MiddleName graphql.String `graphql:"middleName"`
							LastName   graphql.String `graphql:"lastName"`
						} `graphql:"account"`
						Product struct {
							Name         graphql.String `graphql:"name"`
							Id           graphql.ID     `graphql:"id"`
							Abbreviation graphql.String `graphql:"abbreviation"`
						} `graphql:"product"`
					} `graphql:"node"`
				} `graphql:"edges"`
			} `graphql:"sustainabilityAdvisorProducts"`
			AppointedSustainabilityAdvisorProducts struct {
				Edges []struct {
					Node struct {
						ClientCompany struct {
							Id        graphql.ID     `graphql:"id"`
							LegalName graphql.String `graphql:"legalName"`
						} `graphql:"clientCompany"`
						Product struct {
							Id           graphql.ID     `graphql:"id"`
							Abbreviation graphql.String `graphql:"abbreviation"`
						} `graphql:"product"`
						AppointedSustainabilityAdvisors struct {
							Collection []struct {
								Id         graphql.ID     `graphql:"id"`
								FirstName  graphql.String `graphql:"firstName"`
								MiddleName graphql.String `graphql:"middleName"`
								LastName   graphql.String `graphql:"lastName"`
							} `graphql:"collection"`
						} `graphql:"appointedSustainabilityAdvisors"`
					} `graphql:"node"`
				} `graphql:"edges"`
			} `graphql:"appointedSustainabilityAdvisorProducts"`
			CompanyRepresentativeProducts struct {
				Edges []struct {
					Node struct {
						Id           graphql.ID     `graphql:"id"`
						Name         graphql.String `graphql:"name"`
						Abbreviation graphql.String `graphql:"abbreviation"`
					} `graphql:"node"`
				} `graphql:"edges"`
			} `graphql:"companyRepresentativeProducts"`
			Inviter struct {
				Id         graphql.ID     `graphql:"id"`
				FirstName  graphql.String `graphql:"firstName"`
				MiddleName graphql.String `graphql:"middleName"`
				LastName   graphql.String `graphql:"lastName"`
			} `graphql:"inviter"`
			SubmittedForApprovalBy struct {
				Id         graphql.ID     `graphql:"id"`
				FirstName  graphql.String `graphql:"firstName"`
				MiddleName graphql.String `graphql:"middleName"`
				LastName   graphql.String `graphql:"lastName"`
			} `graphql:"submittedForApprovalBy"`
			ClientCompany struct {
				Id               graphql.ID     `graphql:"id"`
				Status           graphql.String `graphql:"status"`
				PrimarySubEntity struct {
					LegalName graphql.String `graphql:"legalName"`
				} `graphql:"primarySubEntity"`
				Products struct {
					Edges []struct {
						Node struct {
							Product struct {
								Name graphql.String `graphql:"name"`
							} `graphql:"product"`
							AppointedSustainabilityAdvisors struct {
								Collection []struct {
									Id         graphql.ID     `graphql:"id"`
									FirstName  graphql.String `graphql:"firstName"`
									LastName   graphql.String `graphql:"lastName"`
									MiddleName graphql.String `graphql:"middleName"`
								} `graphql:"collection"`
							} `graphql:"appointedSustainabilityAdvisors"`
						} `graphql:"node"`
					} `graphql:"edges"`
				} `graphql:"products"`
			} `graphql:"clientCompany"`
		} `graphql:"collection"`
	} `graphql:"accounts(status: $status, type: $type, page: $page, orderByFullName: $orderByFullName)"`
}

var GetAccount struct {
	Account struct {
		Id                            graphql.ID
		Email                         graphql.String
		CreatedAt                     graphql.String
		Status                        graphql.String
		Type                          graphql.String
		Roles                         []graphql.String
		CalendarId                    graphql.String
		CalendlyId                    graphql.String
		IsConnectedGoogleCalendar     graphql.Boolean
		IsCopywriter                  graphql.Boolean
		FirstName                     graphql.String
		LastName                      graphql.String
		MiddleName                    graphql.String
		PhoneNumber                   graphql.String
		Bio                           graphql.String
		Timezone                      graphql.String
		SustainabilityAdvisorProducts struct {
			Edges []struct {
				Node struct {
					Id               graphql.ID
					IsProductManager graphql.Boolean
					Product          struct {
						Id           graphql.ID
						Name         graphql.String
						Abbreviation graphql.String
					}
				}
			}
		}
		AppointedSustainabilityAdvisorProducts struct {
			Edges []struct {
				Node struct {
					Id      graphql.ID
					Product struct {
						Id           graphql.ID
						Abbreviation graphql.String
					}
					ClientCompany struct {
						Id        graphql.ID
						LegalName graphql.String
					}
				}
			}
		}
		CompanyRepresentativeProducts struct {
			Edges []struct {
				Node struct {
					Id           graphql.ID
					Name         graphql.String
					Abbreviation graphql.String
				}
			}
		}
		ProductAccessRequests struct {
			Collection []struct {
				Id      graphql.ID
				Product struct {
					Id           graphql.ID
					Abbreviation graphql.String
					Name         graphql.String
				}
				Status graphql.String
			}
		} `graphql:"productAccessRequests(status: $productStatus)"`
		Avatar struct {
			Id         graphql.ID
			ContentUrl graphql.String
			CreatedAt  graphql.String
			FileName   graphql.String
			FilePath   graphql.String
			UpdatedAt  graphql.String
		}
		Inviter struct {
			Id          graphql.ID
			FirstName   graphql.String
			LastName    graphql.String
			MiddleName  graphql.String
			PhoneNumber graphql.String
		}
		ClientCompany struct {
			Id               graphql.ID
			Status           graphql.String
			PrimarySubEntity struct {
				Id          graphql.ID
				LegalName   graphql.String
				ZohoAccount struct {
					Id   graphql.ID
					Name graphql.String
				}
			}
			Products struct {
				Edges []struct {
					Node struct {
						Id                    graphql.ID
						DefaultDashboardGroup struct {
							Id        graphql.String
							Name      graphql.String
							Slug      graphql.String
							DeletedAt graphql.String
						}
						Product struct {
							Id           graphql.String
							Name         graphql.String
							Abbreviation graphql.String
						}
						AllowedModules struct {
							Edges []struct {
								Node struct {
									Id   graphql.String
									Name graphql.String
									Slug graphql.String
								}
							}
						}
						ServiceType                     graphql.String
						AppointedSustainabilityAdvisors struct {
							Collection []struct {
								Id         graphql.String
								CalendlyId graphql.String
								Email      graphql.String
								FirstName  graphql.String
								LastName   graphql.String
								MiddleName graphql.String
								Avatar     struct {
									ContentUrl graphql.String
								}
							}
						}
					}
				}
			}
		}
	} `graphql:"account(id: $id)"`
}
