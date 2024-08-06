package activecampaign

import (
	"encoding/json"
	"fmt"
	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
	"net/http"
	"net/url"
)

type AccountFields struct {
	AccountFields []AccountField `json:"accountCustomFieldMeta"`
	Meta          Meta           `json:"meta"`
}

type AccountField struct {
	Id                   go_types.Int64String   `json:"id"`
	FieldLabel           string                 `json:"fieldLabel"`
	FieldType            string                 `json:"fieldType"`
	FieldOptions         *[]string              `json:"fieldOptions"`
	FieldDefault         json.RawMessage        `json:"fieldDefault"`
	FieldDefaultCurrency json.RawMessage        `json:"fieldDefaultCurrency"`
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

type GetAccountFieldsConfig struct {
	Limit  *uint64
	Offset *uint64
}

func (service *Service) GetAccountFields(getAccountFieldsConfig *GetAccountFieldsConfig) (*AccountFields, bool, *errortools.Error) {
	params := url.Values{}

	accountFields := AccountFields{}
	rowCount := uint64(0)
	limit := defaultLimit

	if getAccountFieldsConfig != nil {
		if getAccountFieldsConfig.Limit != nil {
			limit = getLimit(*getAccountFieldsConfig.Limit)
		}
	}
	params.Add("limit", fmt.Sprintf("%v", limit))

	for {
		params.Set("offset", fmt.Sprintf("%v", service.nextOffsets.AccountField))

		accountFieldsBatch := AccountFields{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("accountCustomFieldMeta?%s", params.Encode())),
			ResponseModel: &accountFieldsBatch,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, false, e
		}

		accountFields.AccountFields = append(accountFields.AccountFields, accountFieldsBatch.AccountFields...)

		if len(accountFieldsBatch.AccountFields) < int(limit) {
			service.nextOffsets.AccountField = 0
			break
		}

		service.nextOffsets.AccountField += limit
		rowCount += limit

		if rowCount >= service.maxRowCount {
			return &accountFields, true, nil
		}
	}

	return &accountFields, false, nil
}
