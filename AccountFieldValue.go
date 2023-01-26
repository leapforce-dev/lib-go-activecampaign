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

type AccountFieldValues struct {
	AccountFieldValues []AccountFieldValue `json:"accountCustomFieldData"`
	Meta               struct {
		Total int64 `json:"total"`
	} `json:"meta"`
}

type AccountFieldValue struct {
	Id                  go_types.Int64String   `json:"id"`
	AccountFieldMetumId int64                  `json:"accountCustomFieldMetumId"`
	FieldId             int64                  `json:"customFieldId"`
	CreatedTimestamp    a_types.DateTimeString `json:"created_timestamp"`
	UpdatedTimestamp    a_types.DateTimeString `json:"updated_timestamp"`
	FieldValue          string                 `json:"fieldValue"`
	AccountId           int64                  `json:"accountId"`
	Links               *Links                 `json:"links"`
}

type GetAccountFieldValuesConfig struct {
	Limit     *uint64
	Offset    *uint64
	AccountId *int64
}

func (service *Service) GetAccountFieldValues(getAccountFieldValuesConfig *GetAccountFieldValuesConfig) (*AccountFieldValues, *errortools.Error) {
	params := url.Values{}

	accountFieldValues := AccountFieldValues{}
	rowCount := uint64(0)
	limit := defaultLimit

	if getAccountFieldValuesConfig != nil {
		if getAccountFieldValuesConfig.Limit != nil {
			limit = *getAccountFieldValuesConfig.Limit
		}
		if getAccountFieldValuesConfig.AccountId != nil {
			params.Add("filters[customerAccountId]", fmt.Sprintf("%v", *getAccountFieldValuesConfig.AccountId))
		}
	}
	params.Add("limit", fmt.Sprintf("%v", limit))

	for {
		params.Set("offset", fmt.Sprintf("%v", service.nextOffsets.AccountFieldValue))

		accountFieldValuesBatch := AccountFieldValues{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("accountCustomFieldData?%s", params.Encode())),
			ResponseModel: &accountFieldValuesBatch,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		accountFieldValues.AccountFieldValues = append(accountFieldValues.AccountFieldValues, accountFieldValuesBatch.AccountFieldValues...)

		if len(accountFieldValuesBatch.AccountFieldValues) < int(limit) {
			service.nextOffsets.AccountFieldValue = 0
			break
		}

		service.nextOffsets.AccountFieldValue += limit
		rowCount += limit

		if rowCount >= service.maxRowCount {
			return &accountFieldValues, nil
		}
	}

	return &accountFieldValues, nil
}
