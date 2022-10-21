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

type AccountCustomFields struct {
	AccountCustomFields []AccountCustomField `json:"accountCustomFieldMeta"`
	Meta                Meta                 `json:"meta"`
}

type AccountCustomField struct {
	Id                   go_types.Int64String   `json:"id"`
	FieldLabel           string                 `json:"fieldLabel"`
	FieldType            string                 `json:"fieldType"`
	FieldOptions         *[]string              `json:"fieldOptions"`
	FieldDefault         *string                `json:"fieldDefault"`
	FieldDefaultCurrency *string                `json:"fieldDefaultCurrency"`
	IsFormVisible        go_types.BoolInt       `json:"isFormVisible"`
	IsRequired           go_types.BoolInt       `json:"isRequired"`
	DisplayOrder         int64                  `json:"displayOrder"`
	Personalization      string                 `json:"personalization"`
	KnownFieldId         *int64                 `json:"knownFieldId"`
	HideFieldFlag        go_types.BoolInt       `json:"hideFieldFlag"`
	CreatedTimestamp     a_types.DateTimeString `json:"created_timestamp"`
	UpdatedTimestamp     a_types.DateTimeString `json:"updated_timestamp"`
	Links                *Links                 `json:"links"`
}

type GetAccountCustomFieldsConfig struct {
	Limit  *uint64
	Offset *uint64
}

func (service *Service) GetAccountCustomFields(getAccountCustomFieldsConfig *GetAccountCustomFieldsConfig) (*AccountCustomFields, *errortools.Error) {
	params := url.Values{}

	accountCustomFields := AccountCustomFields{}
	rowCount := uint64(0)
	limit := defaultLimit

	if getAccountCustomFieldsConfig != nil {
		if getAccountCustomFieldsConfig.Limit != nil {
			limit = *getAccountCustomFieldsConfig.Limit
		}
	}
	params.Add("limit", fmt.Sprintf("%v", limit))

	for {
		params.Set("offset", fmt.Sprintf("%v", service.nextOffsets.AccountCustomField))

		accountCustomFieldsBatch := AccountCustomFields{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("accountCustomFieldMeta?%s", params.Encode())),
			ResponseModel: &accountCustomFieldsBatch,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		accountCustomFields.AccountCustomFields = append(accountCustomFields.AccountCustomFields, accountCustomFieldsBatch.AccountCustomFields...)

		if len(accountCustomFieldsBatch.AccountCustomFields) < int(limit) {
			service.nextOffsets.AccountCustomField = 0
			break
		}

		service.nextOffsets.AccountCustomField += limit
		rowCount += limit

		if rowCount >= service.maxRowCount {
			return &accountCustomFields, nil
		}
	}

	return &accountCustomFields, nil
}
