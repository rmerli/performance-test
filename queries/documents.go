package queries

import "github.com/hasura/go-graphql-client"

var GetProductDocumentContentUrl struct {
	ProductDocument struct {
		ContentUrl graphql.String `graphql:"contentUrl"`
	} `graphql:"productDocument(id: $id)"`
}

var GetClientDocumentContentUrl struct {
	ClientCompanyDocument struct {
		ContentUrl graphql.String `graphql:"contentUrl"`
	} `graphql:"clientCompanyDocument(id: $id)"`
}

var GetClientDocumentCategories struct {
	ClientDocumentCategories struct {
		Edges []struct {
			Node struct {
				Id            graphql.ID     `graphql:"id"`
				Name          graphql.String `graphql:"name"`
				ClientCompany struct {
					Id graphql.ID `graphql:"id"`
				} `graphql:"clientCompany"`
				ClientCompanyDocuments struct {
					Collection []struct {
						Id                     graphql.ID      `graphql:"id"`
						Status                 graphql.String  `graphql:"status"`
						Title                  graphql.String  `graphql:"title"`
						Note                   graphql.String  `graphql:"note"`
						CreatedAt              graphql.String  `graphql:"createdAt"`
						UpdatedAt              graphql.String  `graphql:"updatedAt"`
						DeletedAt              graphql.String  `graphql:"deletedAt"`
						CanBeDeleted           graphql.Boolean `graphql:"canBeDeleted"`
						CanBeUpdated           graphql.Boolean `graphql:"canBeUpdated"`
						ClientDocumentCategory struct {
							Id   graphql.ID     `graphql:"id"`
							Name graphql.String `graphql:"name"`
						} `graphql:"clientDocumentCategory"`
					} `graphql:"collection"`
				} `graphql:"clientCompanyDocuments"`
			} `graphql:"node"`
		} `graphql:"edges"`
	} `graphql:"clientDocumentCategories(clientCompany: $clientCompany)"`
}

var GetClientCategoryDocuments struct {
	ClientDocumentCategory struct {
		ClientCompanyDocuments struct {
			PaginationInfo struct {
				ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
				LastPage     graphql.Int `graphql:"lastPage"`
				TotalCount   graphql.Int `graphql:"totalCount"`
			} `graphql:"paginationInfo"`
			Collection []struct {
				Id                     graphql.ID      `graphql:"id"`
				Status                 graphql.String  `graphql:"status"`
				Title                  graphql.String  `graphql:"title"`
				Note                   graphql.String  `graphql:"note"`
				CreatedAt              graphql.String  `graphql:"createdAt"`
				UpdatedAt              graphql.String  `graphql:"updatedAt"`
				CanBeDeleted           graphql.Boolean `graphql:"canBeDeleted"`
				CanBeUpdated           graphql.Boolean `graphql:"canBeUpdated"`
				ClientDocumentCategory struct {
					Id   graphql.ID     `graphql:"id"`
					Name graphql.String `graphql:"name"`
				} `graphql:"clientDocumentCategory"`
				ClientCompany struct {
					Id        graphql.ID     `graphql:"id"`
					LegalName graphql.String `graphql:"legalName"`
				} `graphql:"clientCompany"`
			} `graphql:"collection"`
		} `graphql:"clientCompanyDocuments(clientCompany: $clientCompany, itemsPerPage: $itemsPerPage, page: $page, order: { updatedAt: $updatedAtOrder })"`
	} `graphql:"clientDocumentCategory(id: $id)"`
}

var GetProductCategoryDocuments struct {
	ProductDocumentCategory struct {
		ProductDocuments struct {
			PaginationInfo struct {
				ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
				LastPage     graphql.Int `graphql:"lastPage"`
				TotalCount   graphql.Int `graphql:"totalCount"`
			} `graphql:"paginationInfo"`
			Collection []struct {
				Id                      graphql.ID      `graphql:"id"`
				Status                  graphql.String  `graphql:"status"`
				Title                   graphql.String  `graphql:"title"`
				Note                    graphql.String  `graphql:"note"`
				CreatedAt               graphql.String  `graphql:"createdAt"`
				CanBeDeleted            graphql.Boolean `graphql:"canBeDeleted"`
				CanBeUpdated            graphql.Boolean `graphql:"canBeUpdated"`
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
		} `graphql:"productDocuments"`
	} `graphql:"productDocumentCategory(id: $id)"`
}

var GetProductCategoryDocumentsForCR struct {
	ProductDocumentCategory struct {
		ProductDocuments struct {
			PaginationInfo struct {
				ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
				LastPage     graphql.Int `graphql:"lastPage"`
				TotalCount   graphql.Int `graphql:"totalCount"`
			} `graphql:"paginationInfo"`
			Collection []struct {
				Id                      graphql.ID      `graphql:"id"`
				Status                  graphql.String  `graphql:"status"`
				Title                   graphql.String  `graphql:"title"`
				Note                    graphql.String  `graphql:"note"`
				CreatedAt               graphql.String  `graphql:"createdAt"`
				CanBeDeleted            graphql.Boolean `graphql:"canBeDeleted"`
				CanBeUpdated            graphql.Boolean `graphql:"canBeUpdated"`
				ProductDocumentCategory struct {
					Id   graphql.ID     `graphql:"id"`
					Name graphql.String `graphql:"name"`
				} `graphql:"productDocumentCategory"`
			} `graphql:"collection"`
		} `graphql:"productDocuments(status: $documentStatus, page: $page)"`
	} `graphql:"productDocumentCategory(id: $id)"`
}

var GetAllClientsDocuments struct {
	ClientCompanyDocuments struct {
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
			ClientCompany struct {
				Id        graphql.ID     `graphql:"id"`
				LegalName graphql.String `graphql:"legalName"`
			} `graphql:"clientCompany"`
			ClientDocumentCategory struct {
				Id   graphql.ID     `graphql:"id"`
				Name graphql.String `graphql:"name"`
			} `graphql:"clientDocumentCategory"`
		} `graphql:"collection"`
	} `graphql:"clientCompanyDocuments(page: $page, order: { createdAt: $createdAtOrder }, itemsPerPage: $itemsPerPage)"`
}

var GetClientDocumentAllCategories struct {
	ClientDocumentCategories struct {
		Edges []struct {
			Node struct {
				Id            graphql.ID      `graphql:"id"`
				Name          graphql.String  `graphql:"name"`
				CanBeDeleted  graphql.Boolean `graphql:"canBeDeleted"`
				CanBeUpdated  graphql.Boolean `graphql:"canBeUpdated"`
				ClientCompany struct {
					Id graphql.ID `graphql:"id"`
				} `graphql:"clientCompany"`
				ClientCompanyDocuments struct {
					Collection []struct {
						Id                     graphql.ID      `graphql:"id"`
						Status                 graphql.String  `graphql:"status"`
						Title                  graphql.String  `graphql:"title"`
						Note                   graphql.String  `graphql:"note"`
						CreatedAt              graphql.String  `graphql:"createdAt"`
						UpdatedAt              graphql.String  `graphql:"updatedAt"`
						DeletedAt              graphql.String  `graphql:"deletedAt"`
						CanBeDeleted           graphql.Boolean `graphql:"canBeDeleted"`
						CanBeUpdated           graphql.Boolean `graphql:"canBeUpdated"`
						ClientDocumentCategory struct {
							Id   graphql.ID     `graphql:"id"`
							Name graphql.String `graphql:"name"`
						} `graphql:"clientDocumentCategory"`
					} `graphql:"collection"`
				} `graphql:"clientCompanyDocuments(itemsPerPage: $itemsPerPage, order: { updatedAt: $updatedAtOrder })"`
			} `graphql:"node"`
		} `graphql:"edges"`
	} `graphql:"clientDocumentCategories(clientCompany: $clientCompany)"`
}

var GetAllMyDocuments struct {
	ClientDocumentCategories struct {
		Edges []struct {
			Node struct {
				Id            graphql.ID      `graphql:"id"`
				Name          graphql.String  `graphql:"name"`
				CanBeDeleted  graphql.Boolean `graphql:"canBeDeleted"`
				CanBeUpdated  graphql.Boolean `graphql:"canBeUpdated"`
				ClientCompany struct {
					Id graphql.ID `graphql:"id"`
				} `graphql:"clientCompany"`
				ClientCompanyDocuments struct {
					Collection []struct {
						Id                     graphql.ID      `graphql:"id"`
						Status                 graphql.String  `graphql:"status"`
						Title                  graphql.String  `graphql:"title"`
						Note                   graphql.String  `graphql:"note"`
						CreatedAt              graphql.String  `graphql:"createdAt"`
						UpdatedAt              graphql.String  `graphql:"updatedAt"`
						DeletedAt              graphql.String  `graphql:"deletedAt"`
						CanBeDeleted           graphql.Boolean `graphql:"canBeDeleted"`
						CanBeUpdated           graphql.Boolean `graphql:"canBeUpdated"`
						ClientDocumentCategory struct {
							Id   graphql.ID     `graphql:"id"`
							Name graphql.String `graphql:"name"`
						} `graphql:"clientDocumentCategory"`
					} `graphql:"collection"`
				} `graphql:"clientCompanyDocuments(itemsPerPage: $itemsPerPage, order: { updatedAt: $updatedAtOrder })"`
			} `graphql:"node"`
		} `graphql:"edges"`
	} `graphql:"clientDocumentCategories(clientCompany: $clientCompany)"`
}

var GetAllClientsDeletedDocuments struct {
	ClientCompanyDocuments struct {
		PaginationInfo struct {
			ItemsPerPage graphql.Int `graphql:"itemsPerPage"`
			LastPage     graphql.Int `graphql:"lastPage"`
			TotalCount   graphql.Int `graphql:"totalCount"`
		} `graphql:"paginationInfo"`
		Collection []struct {
			Id            graphql.ID     `graphql:"id"`
			Note          graphql.String `graphql:"note"`
			Status        graphql.String `graphql:"status"`
			Title         graphql.String `graphql:"title"`
			ClientCompany struct {
				Id        graphql.ID     `graphql:"id"`
				LegalName graphql.String `graphql:"legalName"`
			} `graphql:"clientCompany"`
			ClientDocumentCategory struct {
				Id   graphql.ID     `graphql:"id"`
				Name graphql.String `graphql:"name"`
			} `graphql:"clientDocumentCategory"`
			DeletedAt    graphql.String  `graphql:"deletedAt"`
			CanBeDeleted graphql.Boolean `graphql:"canBeDeleted"`
			CanBeUpdated graphql.Boolean `graphql:"canBeUpdated"`
			ContentUrl   graphql.String  `graphql:"contentUrl"`
		} `graphql:"collection"`
	} `graphql:"clientCompanyDocuments(page: $page, order: { deletedAt: $deletedAtOrder })"`
}
