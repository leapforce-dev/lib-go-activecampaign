package activecampaign

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
	"net/http"
	"net/url"
)

type AccountContactAssociations struct {
	AccountContactAssociations []AccountContactAssociation `json:"accountContacts"`
	Meta                       Meta                        `json:"meta"`
}

type AccountContactAssociation struct {
	Account  go_types.Int64String `json:"account"`
	Contact  go_types.Int64String `json:"contact"`
	JobTitle string               `json:"jobTitle"`
	Links    *Links               `json:"links"`
	Id       go_types.Int64String `json:"id"`
}

type GetAccountContactAssociationsConfig struct {
	Limit     *uint64
	Offset    *uint64
	AccountId *int64
	ContactId *int64
}

func (service *Service) GetAccountContactAssociations(getAccountContactAssociationsConfig *GetAccountContactAssociationsConfig) (*AccountContactAssociations, *errortools.Error) {
	params := url.Values{}

	accountContactAssociations := AccountContactAssociations{}
	rowCount := uint64(0)
	limit := defaultLimit

	if getAccountContactAssociationsConfig != nil {
		if getAccountContactAssociationsConfig.Limit != nil {
			limit = *getAccountContactAssociationsConfig.Limit
		}
		if getAccountContactAssociationsConfig.AccountId != nil {
			params.Add("filters[account]", fmt.Sprintf("%v", *getAccountContactAssociationsConfig.AccountId))
		}
		if getAccountContactAssociationsConfig.ContactId != nil {
			params.Add("filters[contact]", fmt.Sprintf("%v", *getAccountContactAssociationsConfig.ContactId))
		}
	}
	params.Add("limit", fmt.Sprintf("%v", limit))

	for {
		params.Set("offset", fmt.Sprintf("%v", service.nextOffsets.AccountContactAssociation))

		accountContactAssociationsBatch := AccountContactAssociations{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("accountContacts?%s", params.Encode())),
			ResponseModel: &accountContactAssociationsBatch,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		accountContactAssociations.AccountContactAssociations = append(accountContactAssociations.AccountContactAssociations, accountContactAssociationsBatch.AccountContactAssociations...)

		if len(accountContactAssociationsBatch.AccountContactAssociations) < int(limit) {
			service.nextOffsets.AccountContactAssociation = 0
			break
		}

		service.nextOffsets.AccountContactAssociation += limit
		rowCount += limit

		if rowCount >= service.maxRowCount {
			return &accountContactAssociations, nil
		}
	}

	return &accountContactAssociations, nil
}
