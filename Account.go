package activecampaign

import (
	"fmt"
	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
	"net/http"
	"net/url"
)

type Accounts struct {
	Accounts []Account `json:"accounts"`
	Meta     Meta      `json:"meta"`
}

type Account struct {
	Name                string                 `json:"name"`
	AccountUrl          *go_types.String       `json:"accountUrl"`
	CreatedUtcTimestamp a_types.DateTimeString `json:"created_utc_timestamp"`
	UpdatedUtcTimestamp a_types.DateTimeString `json:"updated_utc_timestamp"`
	ContactCount        go_types.Int64String   `json:"contactCount"`
	DealCount           go_types.Int64String   `json:"dealCount"`
	Owner               *go_types.Int64String  `json:"owner"`
	Id                  go_types.Int64String   `json:"id"`
	Links               *Links                 `json:"links"`
}

type GetAccountsConfig struct {
	Limit      *uint64
	Offset     *uint64
	Search     *string
	CountDeals *bool
}

func (service *Service) GetAccounts(getAccountsConfig *GetAccountsConfig) (*Accounts, *errortools.Error) {
	params := url.Values{}

	accounts := Accounts{}
	rowCount := uint64(0)
	limit := defaultLimit

	if getAccountsConfig != nil {
		if getAccountsConfig.Limit != nil {
			limit = *getAccountsConfig.Limit
		}
		if getAccountsConfig.Offset != nil {
			service.nextOffsets.Account = *getAccountsConfig.Offset
		}
		if getAccountsConfig.Search != nil {
			params.Add("search", *getAccountsConfig.Search)
		}
		if getAccountsConfig.CountDeals != nil {
			params.Add("count_deals", fmt.Sprintf("%v", *getAccountsConfig.CountDeals))
		}
	}
	params.Add("limit", fmt.Sprintf("%v", limit))

	for {
		params.Set("offset", fmt.Sprintf("%v", service.nextOffsets.Account))

		accountsBatch := Accounts{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("accounts?%s", params.Encode())),
			ResponseModel: &accountsBatch,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		accounts.Accounts = append(accounts.Accounts, accountsBatch.Accounts...)

		if len(accountsBatch.Accounts) < int(limit) {
			service.nextOffsets.Account = 0
			break
		}

		service.nextOffsets.Account += limit
		rowCount += limit

		if rowCount >= service.maxRowCount {
			return &accounts, nil
		}
	}

	return &accounts, nil
}

type AccountSync struct {
	Name        *string              `json:"name,omitempty"`
	AccountUrl  *string              `json:"accountUrl,omitempty"`
	Owner       *int64               `json:"owner,omitempty"`
	FieldValues *[]AccountFieldValue `json:"fields,omitempty"`
}

func (service *Service) UpdateAccount(accountId string, accountCreate AccountSync) (*Account, *errortools.Error) {
	d := struct {
		Account AccountSync `json:"account"`
	}{
		Account: accountCreate,
	}

	var accountUpdated struct {
		Account Account `json:"account"`
	}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url(fmt.Sprintf("accounts/%s", accountId)),
		BodyModel:     d,
		ResponseModel: &accountUpdated,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &accountUpdated.Account, nil
}

func (service *Service) DeleteAccount(accountId int64) *errortools.Error {
	requestConfig := go_http.RequestConfig{
		Method: http.MethodDelete,
		Url:    service.url(fmt.Sprintf("accounts/%v", accountId)),
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return e
	}

	return nil
}
