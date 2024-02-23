package queries

import "github.com/hasura/go-graphql-client"

var GetClientCompanies struct {
	ClientCompanies struct {
		PaginationInfo struct {
			ItemsPerPage graphql.Int
			LastPage     graphql.Int
			TotalCount   graphql.Int
		}
		Collection []struct {
			Id                      graphql.ID
			IsProductManagerRelated graphql.Boolean
			Status                  graphql.String
			LegalName               graphql.String
			PrimarySubEntity        struct {
				LegalName graphql.String
			}
			CompanyRepresentatives struct {
				Collection []struct {
					Id                            graphql.ID
					FirstName                     graphql.String
					LastName                      graphql.String
					CompanyRepresentativeProducts struct {
						Edges []struct {
							Node struct {
								Abbreviation graphql.String
							}
						}
					}
				}
			}
			Products struct {
				Edges []struct {
					Node struct {
						Product struct {
							Name         graphql.String
							Abbreviation graphql.String
							Id           graphql.ID
						}
						AppointedSustainabilityAdvisors struct {
							Collection []struct {
								Id                            graphql.ID
								Email                         graphql.String
								SustainabilityAdvisorProducts struct {
									Edges []struct {
										Node struct {
											Product struct {
												Name graphql.String
											}
										}
									}
								}
								FirstName graphql.String
								LastName  graphql.String
							}
						}
					}
				}
			}
		}
	}
}

var GetClientCompaniesFilteredByProduct struct {
	ClientCompanies struct {
		PaginationInfo struct {
			ItemsPerPage graphql.Int
			LastPage     graphql.Int
			TotalCount   graphql.Int
		} `graphql:"paginationInfo"`
		Collection []struct {
			Id        graphql.ID     `graphql:"id"`
			LegalName graphql.String `graphql:"legalName"`
			Products  struct {
				Edges []struct {
					Node struct {
						Product struct {
							Id           graphql.ID     `graphql:"id"`
							Abbreviation graphql.String `graphql:"abbreviation"`
						} `graphql:"product"`
					}
				}
			} `graphql:"products"`
		}
	} `graphql:"clientCompanies(products_product_id_list: $products_product_id_list, status_list: $status_list, itemsPerPage: $itemsPerPage, page: $page)"`
}

var GetClientCompanyProducts struct {
	ClientCompany struct {
		Id        graphql.ID     `graphql:"id"`
		LegalName graphql.String `graphql:"legalName"`
		Products  struct {
			Edges []struct {
				Node struct {
					Product struct {
						Id           graphql.ID     `graphql:"id"`
						Abbreviation graphql.String `graphql:"abbreviation"`
					} `graphql:"product"`
				}
			} `graphql:"edges"`
		} `graphql:"products"`
	} `graphql:"clientCompany(id: $id)"`
}

var GetCompanyOnboarding struct {
	ClientCompany struct {
		Id               graphql.ID `graphql:"id"`
		PrimarySubEntity struct {
			Addition          graphql.String `graphql:"addition"`
			LegalEmail        graphql.String `graphql:"legalEmail"`
			LegalName         graphql.String `graphql:"legalName"`
			LegalFullName     graphql.String `graphql:"legalFullName"`
			Country           graphql.String `graphql:"country"`
			City              graphql.String `graphql:"city"`
			Address           graphql.String `graphql:"address"`
			PostalCode        graphql.String `graphql:"postalCode"`
			ChamberOfCommerce graphql.String `graphql:"chamberOfCommerce"`
			LEI               graphql.String `graphql:"lei"`
			VAT               graphql.String `graphql:"vat"`
			Id                graphql.ID     `graphql:"id"`
			ZohoAccount       struct {
				Id   graphql.ID     `graphql:"id"`
				Name graphql.String `graphql:"name"`
			} `graphql:"zohoAccount"`
		} `graphql:"primarySubEntity"`
		CompanyRepresentatives struct {
			PaginationInfo struct {
				ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
				LastPage     graphql.Int `graphql:"lastPage"`
				TotalCount   graphql.Int `graphql:"totalCount"`
			} `graphql:"paginationInfo"`
		} `graphql:"companyRepresentatives"`
	} `graphql:"clientCompany(id: $id)"`
}

var GetCompany struct {
	ClientCompany struct {
		Id          graphql.ID     `graphql:"id"`
		Status      graphql.String `graphql:"status"`
		CreatedAt   graphql.String `graphql:"createdAt"`
		SubEntities struct {
			Edges []struct {
				Node struct {
					Id            graphql.ID     `graphql:"id"`
					LegalName     graphql.String `graphql:"legalName"`
					LegalFullName graphql.String `graphql:"legalFullName"`
					ZohoAccount   struct {
						Name graphql.String `graphql:"name"`
					} `graphql:"zohoAccount"`
				}
			} `graphql:"edges"`
		} `graphql:"subEntities"`
		PrimarySubEntity struct {
			Addition          graphql.String `graphql:"addition"`
			LegalEmail        graphql.String `graphql:"legalEmail"`
			LegalName         graphql.String `graphql:"legalName"`
			LegalFullName     graphql.String `graphql:"legalFullName"`
			Country           graphql.String `graphql:"country"`
			City              graphql.String `graphql:"city"`
			Address           graphql.String `graphql:"address"`
			PostalCode        graphql.String `graphql:"postalCode"`
			ChamberOfCommerce graphql.String `graphql:"chamberOfCommerce"`
			LEI               graphql.String `graphql:"lei"`
			VAT               graphql.String `graphql:"vat"`
			Id                graphql.ID     `graphql:"id"`
			ZohoAccount       struct {
				Id   graphql.ID     `graphql:"id"`
				Name graphql.String `graphql:"name"`
			} `graphql:"zohoAccount"`
		} `graphql:"primarySubEntity"`
		Products struct {
			Edges []struct {
				Node struct {
					Id                    graphql.ID     `graphql:"id"`
					ServiceType           graphql.String `graphql:"serviceType"`
					DefaultDashboardGroup struct {
						Id        graphql.ID     `graphql:"id"`
						Name      graphql.String `graphql:"name"`
						Slug      graphql.String `graphql:"slug"`
						DeletedAt graphql.String `graphql:"deletedAt"`
					} `graphql:"defaultDashboardGroup"`
					Product struct {
						Id           graphql.ID     `graphql:"id"`
						Abbreviation graphql.String `graphql:"abbreviation"`
						Name         graphql.String `graphql:"name"`
					} `graphql:"product"`
					AppointedSustainabilityAdvisors struct {
						Collection []struct {
							Id         graphql.ID     `graphql:"id"`
							FirstName  graphql.String `graphql:"firstName"`
							MiddleName graphql.String `graphql:"middleName"`
							LastName   graphql.String `graphql:"lastName"`
						} `graphql:"collection"`
					} `graphql:"appointedSustainabilityAdvisors"`
					AllowedModules struct {
						Edges []struct {
							Node struct {
								Id   graphql.ID     `graphql:"id"`
								Name graphql.String `graphql:"name"`
								Slug graphql.String `graphql:"slug"`
							} `graphql:"node"`
						} `graphql:"edges"`
					} `graphql:"allowedModules"`
				} `graphql:"node"`
			} `graphql:"edges"`
		} `graphql:"products"`
		CompanyRepresentatives struct {
			Collection []struct {
				Id                            graphql.ID `graphql:"id"`
				CompanyRepresentativeProducts struct {
					Edges []struct {
						Node struct {
							Name         graphql.String `graphql:"name"`
							Id           graphql.ID     `graphql:"id"`
							Abbreviation graphql.String `graphql:"abbreviation"`
						} `graphql:"node"`
					} `graphql:"edges"`
				} `graphql:"companyRepresentativeProducts"`
				FirstName   graphql.String `graphql:"firstName"`
				MiddleName  graphql.String `graphql:"middleName"`
				LastName    graphql.String `graphql:"lastName"`
				PhoneNumber graphql.String `graphql:"phoneNumber"`
				Email       graphql.String `graphql:"email"`
				Status      graphql.String `graphql:"status"`
				Inviter     struct {
					Id          graphql.ID     `graphql:"id"`
					FirstName   graphql.String `graphql:"firstName"`
					LastName    graphql.String `graphql:"lastName"`
					MiddleName  graphql.String `graphql:"middleName"`
					PhoneNumber graphql.String `graphql:"phoneNumber"`
					Avatar      struct {
						Id         graphql.ID     `graphql:"id"`
						ContentUrl graphql.String `graphql:"contentUrl"`
					} `graphql:"avatar"`
				} `graphql:"inviter"`
				SubmittedForApprovalBy struct {
					Id          graphql.ID     `graphql:"id"`
					FirstName   graphql.String `graphql:"firstName"`
					LastName    graphql.String `graphql:"lastName"`
					MiddleName  graphql.String `graphql:"middleName"`
					PhoneNumber graphql.String `graphql:"phoneNumber"`
					Avatar      struct {
						Id         graphql.ID     `graphql:"id"`
						ContentUrl graphql.String `graphql:"contentUrl"`
					} `graphql:"avatar"`
				} `graphql:"submittedForApprovalBy"`
			} `graphql:"collection"`
		} `graphql:"companyRepresentatives(page: $page, itemsPerPage: $itemsPerPage)"`
		UpdatedAt graphql.String `graphql:"updatedAt"`
		CreatedBy struct {
			Id graphql.ID `graphql:"id"`
		} `graphql:"createdBy"`
	} `graphql:"clientCompany(id: $id)"`
}

var GetCompanyClientView struct {
	ClientCompany struct {
		Id          graphql.ID     `graphql:"id"`
		Status      graphql.String `graphql:"status"`
		CreatedAt   graphql.String `graphql:"createdAt"`
		SubEntities struct {
			Edges []struct {
				Node struct {
					Id            graphql.ID     `graphql:"id"`
					LegalName     graphql.String `graphql:"legalName"`
					LegalFullName graphql.String `graphql:"legalFullName"`
				}
			} `graphql:"edges"`
		} `graphql:"subEntities"`
		PrimarySubEntity struct {
			Addition          graphql.String `graphql:"addition"`
			LegalEmail        graphql.String `graphql:"legalEmail"`
			LegalName         graphql.String `graphql:"legalName"`
			LegalFullName     graphql.String `graphql:"legalFullName"`
			Country           graphql.String `graphql:"country"`
			City              graphql.String `graphql:"city"`
			Address           graphql.String `graphql:"address"`
			PostalCode        graphql.String `graphql:"postalCode"`
			ChamberOfCommerce graphql.String `graphql:"chamberOfCommerce"`
			LEI               graphql.String `graphql:"lei"`
			VAT               graphql.String `graphql:"vat"`
			Id                graphql.ID     `graphql:"id"`
			ZohoAccount       struct {
				Id   graphql.ID     `graphql:"id"`
				Name graphql.String `graphql:"name"`
			} `graphql:"zohoAccount"`
		} `graphql:"primarySubEntity"`
		Products struct {
			Edges []struct {
				Node struct {
					Id                    graphql.ID     `graphql:"id"`
					ServiceType           graphql.String `graphql:"serviceType"`
					DefaultDashboardGroup struct {
						Id        graphql.ID     `graphql:"id"`
						Name      graphql.String `graphql:"name"`
						Slug      graphql.String `graphql:"slug"`
						DeletedAt graphql.String `graphql:"deletedAt"`
					} `graphql:"defaultDashboardGroup"`
					Product struct {
						Id           graphql.ID     `graphql:"id"`
						Abbreviation graphql.String `graphql:"abbreviation"`
						Name         graphql.String `graphql:"name"`
					} `graphql:"product"`
					AppointedSustainabilityAdvisors struct {
						Collection []struct {
							Id         graphql.ID     `graphql:"id"`
							FirstName  graphql.String `graphql:"firstName"`
							MiddleName graphql.String `graphql:"middleName"`
							LastName   graphql.String `graphql:"lastName"`
						} `graphql:"collection"`
					} `graphql:"appointedSustainabilityAdvisors"`
					AllowedModules struct {
						Edges []struct {
							Node struct {
								Id   graphql.ID     `graphql:"id"`
								Name graphql.String `graphql:"name"`
								Slug graphql.String `graphql:"slug"`
							} `graphql:"node"`
						} `graphql:"edges"`
					} `graphql:"allowedModules"`
				} `graphql:"node"`
			} `graphql:"edges"`
		} `graphql:"products"`
		CompanyRepresentatives struct {
			Collection []struct {
				Id                            graphql.ID     `graphql:"id"`
				FirstName                     graphql.String `graphql:"firstName"`
				MiddleName                    graphql.String `graphql:"middleName"`
				LastName                      graphql.String `graphql:"lastName"`
				Status                        graphql.String `graphql:"status"`
				CompanyRepresentativeProducts struct {
					Edges []struct {
						Node struct {
							Name         graphql.String `graphql:"name"`
							Abbreviation graphql.String `graphql:"abbreviation"`
						} `graphql:"node"`
					} `graphql:"edges"`
				} `graphql:"companyRepresentativeProducts"`
			} `graphql:"collection"`
		} `graphql:"companyRepresentatives(itemsPerPage: $itemsPerPage)"`
		UpdatedAt graphql.String `graphql:"updatedAt"`
	} `graphql:"clientCompany(id: $id)"`
}

var GetCompanyStatus struct {
	ClientCompany struct {
		Status graphql.String `graphql:"status"`
	} `graphql:"clientCompany(id: $id)"`
}

var GetSubEntity struct {
	SubEntity struct {
		Addition          graphql.String `graphql:"addition"`
		Address           graphql.String `graphql:"address"`
		LegalName         graphql.String `graphql:"legalName"`
		ChamberOfCommerce graphql.String `graphql:"chamberOfCommerce"`
		City              graphql.String `graphql:"city"`
		ClientCompany     struct {
			Id        graphql.ID     `graphql:"id"`
			LegalName graphql.String `graphql:"legalName"`
		} `graphql:"clientCompany"`
		Country       graphql.String `graphql:"country"`
		DeletedAt     graphql.String `graphql:"deletedAt"`
		Id            graphql.ID     `graphql:"id"`
		LegalEmail    graphql.String `graphql:"legalEmail"`
		LegalFullName graphql.String `graphql:"legalFullName"`
		LEI           graphql.String `graphql:"lei"`
		PostalCode    graphql.String `graphql:"postalCode"`
		Status        graphql.String `graphql:"status"`
		VAT           graphql.String `graphql:"vat"`
	} `graphql:"subEntity(id: $id)"`
}

var GetSubEntityBsupp struct {
	SubEntity struct {
		Addition          graphql.String `graphql:"addition"`
		Address           graphql.String `graphql:"address"`
		LegalName         graphql.String `graphql:"legalName"`
		ChamberOfCommerce graphql.String `graphql:"chamberOfCommerce"`
		City              graphql.String `graphql:"city"`
		ClientCompany     struct {
			Id        graphql.ID     `graphql:"id"`
			LegalName graphql.String `graphql:"legalName"`
		} `graphql:"clientCompany"`
		Country       graphql.String `graphql:"country"`
		DeletedAt     graphql.String `graphql:"deletedAt"`
		Id            graphql.ID     `graphql:"id"`
		LegalEmail    graphql.String `graphql:"legalEmail"`
		LegalFullName graphql.String `graphql:"legalFullName"`
		LEI           graphql.String `graphql:"lei"`
		PostalCode    graphql.String `graphql:"postalCode"`
		Status        graphql.String `graphql:"status"`
		VAT           graphql.String `graphql:"vat"`
		ZohoAccount   struct {
			Id   graphql.ID     `graphql:"id"`
			Name graphql.String `graphql:"name"`
		} `graphql:"zohoAccount"`
	} `graphql:"subEntity(id: $id)"`
}

var GetZohoAccounts struct {
	ZohoAccounts struct {
		PaginationInfo struct {
			ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
			LastPage     graphql.Int `graphql:"lastPage"`
			TotalCount   graphql.Int `graphql:"totalCount"`
		} `graphql:"paginationInfo"`
		Collection []struct {
			Id   graphql.ID     `graphql:"id"`
			Name graphql.String `graphql:"name"`
		} `graphql:"collection"`
	} `graphql:"zohoAccounts(page: $page, itemsPerPage: $itemsPerPage, order: { name: $order })"`
}

var GetSharedWithCompanies struct {
	ClientCompanies struct {
		PaginationInfo struct {
			TotalCount graphql.Int `graphql:"totalCount"`
		} `graphql:"paginationInfo"`
		Collection []struct {
			Id        graphql.ID     `graphql:"id"`
			LegalName graphql.String `graphql:"legalName"`
		} `graphql:"collection"`
	} `graphql:"clientCompanies(products_product_id: $product, status_list: $status_list, itemsPerPage: $itemsPerPage, client: $client)"`
}

var GetClientViewCompanies struct {
	ClientCompanies struct {
		PaginationInfo struct {
			TotalCount graphql.Int `graphql:"totalCount"`
			LastPage   graphql.Int `graphql:"lastPage"`
		} `graphql:"paginationInfo"`
		Collection []struct {
			Id        graphql.ID     `graphql:"id"`
			LegalName graphql.String `graphql:"legalName"`
			Products  struct {
				Edges []struct {
					Node struct {
						Product struct {
							Name         graphql.String `graphql:"name"`
							Abbreviation graphql.String `graphql:"abbreviation"`
							Id           graphql.ID     `graphql:"id"`
						} `graphql:"product"`
						AllowedModules struct {
							Edges []struct {
								Node struct {
									Id   graphql.ID     `graphql:"id"`
									Name graphql.String `graphql:"name"`
									Slug graphql.String `graphql:"slug"`
								} `graphql:"node"`
							} `graphql:"edges"`
						} `graphql:"allowedModules"`
					} `graphql:"node"`
				} `graphql:"edges"`
			} `graphql:"products"`
		} `graphql:"collection"`
	} `graphql:"clientCompanies(client: $client, products_product_id: $product itemsPerPage: $itemsPerPage page: $page)"`
}

var GetClientName struct {
	ClientCompanies struct {
		PaginationInfo struct {
			ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
			LastPage     graphql.Int `graphql:"lastPage"`
			TotalCount   graphql.Int `graphql:"totalCount"`
		} `graphql:"paginationInfo"`
		Collection []struct {
			Id        graphql.ID     `graphql:"id"`
			LegalName graphql.String `graphql:"legalName"`
		} `graphql:"collection"`
	} `graphql:"clientCompanies(products_product_id: $product, itemsPerPage: $itemsPerPage, page: $page, client: $client)"`
}

var GetAllClientCompanies struct {
	ClientCompanies struct {
		PaginationInfo struct {
			LastPage     graphql.Int `graphql:"lastPage"`
			TotalCount   graphql.Int `graphql:"totalCount"`
			ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
		} `graphql:"paginationInfo"`
		Collection []struct {
			Id        graphql.ID     `graphql:"id"`
			Status    graphql.String `graphql:"status"`
			LegalName graphql.String `graphql:"legalName"`
			Products  struct {
				Edges []struct {
					Node struct {
						Product struct {
							Name         graphql.String `graphql:"name"`
							Abbreviation graphql.String `graphql:"abbreviation"`
							Id           graphql.ID     `graphql:"id"`
						} `graphql:"product"`
					} `graphql:"node"`
				} `graphql:"edges"`
			} `graphql:"products"`
		} `graphql:"collection"`
	} `graphql:"clientCompanies(page: $page, itemsPerPage: $itemsPerPage)"`
}
