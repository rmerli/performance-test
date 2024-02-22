package queries

import "github.com/hasura/go-graphql-client"

var GetProducts struct {
	Products struct {
		Edges []struct {
			Node struct {
				Id           graphql.ID
				Name         graphql.String
				Abbreviation graphql.String
				Modules      struct {
					Edges []struct {
						Node struct {
							Id   graphql.ID
							Name graphql.String
							Slug graphql.String
						}
					}
				}
			}
		}
	}
}

var GetProductCategories struct {
	ProductDocumentCategories struct {
		Edges []struct {
			Node struct {
				Id           graphql.ID      `graphql:"id"`
				Name         graphql.String  `graphql:"name"`
				CanBeDeleted graphql.Boolean `graphql:"canBeDeleted"`
				CanBeUpdated graphql.Boolean `graphql:"canBeUpdated"`
				Product      struct {
					Id           graphql.ID     `graphql:"id"`
					Abbreviation graphql.String `graphql:"abbreviation"`
				} `graphql:"product"`
				ProductDocuments struct {
					PaginationInfo struct {
						ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
						LastPage     graphql.Int `graphql:"lastPage"`
						TotalCount   graphql.Int `graphql:"totalCount"`
					} `graphql:"paginationInfo"`
					Collection []struct {
						Id           graphql.ID      `graphql:"id"`
						Note         graphql.String  `graphql:"note"`
						Status       graphql.String  `graphql:"status"`
						Title        graphql.String  `graphql:"title"`
						CreatedAt    graphql.String  `graphql:"createdAt"`
						UpdatedAt    graphql.String  `graphql:"updatedAt"`
						DeletedAt    graphql.String  `graphql:"deletedAt"`
						CanBeDeleted graphql.Boolean `graphql:"canBeDeleted"`
						CanBeUpdated graphql.Boolean `graphql:"canBeUpdated"`
						CreatedBy    struct {
							Id         graphql.ID     `graphql:"id"`
							FirstName  graphql.String `graphql:"firstName"`
							MiddleName graphql.String `graphql:"middleName"`
							LastName   graphql.String `graphql:"lastName"`
						} `graphql:"createdBy"`
						ProductDocumentCategory struct {
							Id   graphql.ID     `graphql:"id"`
							Name graphql.String `graphql:"name"`
						} `graphql:"productDocumentCategory"`
						SharedWith struct {
							Collection []struct {
								Id        graphql.ID     `graphql:"id"`
								LegalName graphql.String `graphql:"legalName"`
							} `graphql:"collection"`
						} `graphql:"sharedWith"`
					} `graphql:"collection"`
				} `graphql:"productDocuments(page: $page, itemsPerPage: $itemsPerPage, order: { updatedAt: $updatedAtOrder, createdAt: $createdAtOrder })"`
			} `graphql:"node"`
		} `graphql:"edges"`
	} `graphql:"productDocumentCategories"`
}

var GetAllProductDocuments struct {
	ProductDocuments struct {
		PaginationInfo struct {
			ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
			LastPage     graphql.Int `graphql:"lastPage"`
			TotalCount   graphql.Int `graphql:"totalCount"`
		} `graphql:"paginationInfo"`
		Collection []struct {
			Id           graphql.ID      `graphql:"id"`
			Note         graphql.String  `graphql:"note"`
			Status       graphql.String  `graphql:"status"`
			Title        graphql.String  `graphql:"title"`
			CreatedAt    graphql.String  `graphql:"createdAt"`
			CanBeDeleted graphql.Boolean `graphql:"canBeDeleted"`
			CanBeUpdated graphql.Boolean `graphql:"canBeUpdated"`
			CreatedBy    struct {
				Id         graphql.ID     `graphql:"id"`
				FirstName  graphql.String `graphql:"firstName"`
				MiddleName graphql.String `graphql:"middleName"`
				LastName   graphql.String `graphql:"lastName"`
			} `graphql:"createdBy"`
			ProductDocumentCategory struct {
				Id   graphql.ID     `graphql:"id"`
				Name graphql.String `graphql:"name"`
			} `graphql:"productDocumentCategory"`
			Product struct {
				Id           graphql.ID     `graphql:"id"`
				Abbreviation graphql.String `graphql:"abbreviation"`
			} `graphql:"product"`
			SharedWith struct {
				Collection []struct {
					Id        graphql.ID     `graphql:"id"`
					LegalName graphql.String `graphql:"legalName"`
				} `graphql:"collection"`
				PaginationInfo struct {
					ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
					LastPage     graphql.Int `graphql:"lastPage"`
					TotalCount   graphql.Int `graphql:"totalCount"`
				} `graphql:"paginationInfo"`
			} `graphql:"sharedWith"`
		} `graphql:"collection"`
	} `graphql:"productDocuments(status: $documentStatus, page: $page, order: { createdAt: $createdAtOrder }, itemsPerPage: $itemsPerPage)"`
}

var GetProductCategoriesForCR struct {
	ProductDocumentCategories struct {
		Edges []struct {
			Node struct {
				Id           graphql.ID      `graphql:"id"`
				Name         graphql.String  `graphql:"name"`
				CanBeDeleted graphql.Boolean `graphql:"canBeDeleted"`
				CanBeUpdated graphql.Boolean `graphql:"canBeUpdated"`
				Product      struct {
					Id           graphql.ID     `graphql:"id"`
					Abbreviation graphql.String `graphql:"abbreviation"`
				} `graphql:"product"`
				ProductDocuments struct {
					PaginationInfo struct {
						ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
						LastPage     graphql.Int `graphql:"lastPage"`
						TotalCount   graphql.Int `graphql:"totalCount"`
					} `graphql:"paginationInfo"`
					Collection []struct {
						Id      graphql.ID `graphql:"id"`
						Product struct {
							Id           graphql.ID     `graphql:"id"`
							Abbreviation graphql.String `graphql:"abbreviation"`
						} `graphql:"product"`
						Note      graphql.String `graphql:"note"`
						Status    graphql.String `graphql:"status"`
						Title     graphql.String `graphql:"title"`
						CreatedAt graphql.String `graphql:"createdAt"`
						UpdatedAt graphql.String `graphql:"updatedAt"`
						CreatedBy struct {
							Id         graphql.ID     `graphql:"id"`
							FirstName  graphql.String `graphql:"firstName"`
							MiddleName graphql.String `graphql:"middleName"`
							LastName   graphql.String `graphql:"lastName"`
						} `graphql:"createdBy"`
						DeletedAt               graphql.String  `graphql:"deletedAt"`
						CanBeDeleted            graphql.Boolean `graphql:"canBeDeleted"`
						CanBeUpdated            graphql.Boolean `graphql:"canBeUpdated"`
						ProductDocumentCategory struct {
							Id   graphql.ID     `graphql:"id"`
							Name graphql.String `graphql:"name"`
						} `graphql:"productDocumentCategory"`
					} `graphql:"collection"`
				} `graphql:"productDocuments(status: $documentStatus, page: $page, itemsPerPage: $itemsPerPage, order: { updatedAt: $updatedAtOrder })"`
			} `graphql:"node"`
		} `graphql:"edges"`
		TotalCount graphql.Int `graphql:"totalCount"`
	} `graphql:"productDocumentCategories(product: $productId)"`
}

var GetDeletedProduct struct {
	ProductDocuments struct {
		PaginationInfo struct {
			ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
			LastPage     graphql.Int `graphql:"lastPage"`
			TotalCount   graphql.Int `graphql:"totalCount"`
		} `graphql:"paginationInfo"`
		Collection []struct {
			Id           graphql.ID      `graphql:"id"`
			ContentUrl   graphql.String  `graphql:"contentUrl"`
			Note         graphql.String  `graphql:"note"`
			Status       graphql.String  `graphql:"status"`
			Title        graphql.String  `graphql:"title"`
			CreatedAt    graphql.String  `graphql:"createdAt"`
			DeletedAt    graphql.String  `graphql:"deletedAt"`
			CanBeDeleted graphql.Boolean `graphql:"canBeDeleted"`
			CanBeUpdated graphql.Boolean `graphql:"canBeUpdated"`
			CreatedBy    struct {
				Id         graphql.ID     `graphql:"id"`
				FirstName  graphql.String `graphql:"firstName"`
				MiddleName graphql.String `graphql:"middleName"`
				LastName   graphql.String `graphql:"lastName"`
			} `graphql:"createdBy"`
			ProductDocumentCategory struct {
				Id   graphql.ID     `graphql:"id"`
				Name graphql.String `graphql:"name"`
			} `graphql:"productDocumentCategory"`
			SharedWith struct {
				Collection []struct {
					Id        graphql.ID     `graphql:"id"`
					LegalName graphql.String `graphql:"legalName"`
				} `graphql:"collection"`
			} `graphql:"sharedWith"`
		} `graphql:"collection"`
	} `graphql:"productDocuments(status: $documentStatus, page: $page, order: { deletedAt: $deletedAtOrder })"`
}

var GetProductAccessRequests struct {
	ProductAccessRequests struct {
		PaginationInfo struct {
			ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
			LastPage     graphql.Int `graphql:"lastPage"`
			TotalCount   graphql.Int `graphql:"totalCount"`
		} `graphql:"paginationInfo"`
		Collection []struct {
			Id      graphql.ID `graphql:"id"`
			Account struct {
				Id            graphql.ID     `graphql:"id"`
				FirstName     graphql.String `graphql:"firstName"`
				MiddleName    graphql.String `graphql:"middleName"`
				LastName      graphql.String `graphql:"lastName"`
				ClientCompany struct {
					Id        graphql.ID     `graphql:"id"`
					LegalName graphql.String `graphql:"legalName"`
				} `graphql:"clientCompany"`
			} `graphql:"account"`
			Product struct {
				Id           graphql.ID     `graphql:"id"`
				Abbreviation graphql.String `graphql:"abbreviation"`
				Name         graphql.String `graphql:"name"`
			} `graphql:"product"`
			Status    graphql.String `graphql:"status"`
			CreatedAt graphql.String `graphql:"createdAt"`
		} `graphql:"collection"`
	} `graphql:"productAccessRequests(page: $page, itemsPerPage: $itemsPerPage, status: $status, order: { createdAt: \"DESC\" })"`
}
