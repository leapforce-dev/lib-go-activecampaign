package activecampaign

import (
	"encoding/json"
	"fmt"
	"net/url"

	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type DealFields struct {
	DealFields []DealField `json:"dealCustomFieldMeta"`
	Meta       Meta        `json:"meta"`
}

type DealField struct {
	ID               go_types.Int64String           `json:"id"`
	FieldLabel       string                         `json:"fieldLabel"`
	FieldType        string                         `json:"fieldType"`
	Type             string                         `json:"type"`
	FieldOptions     json.RawMessage                `json:"fieldOptions"`
	FieldDefault     *go_types.String               `json:"fieldDefault"`
	IsFormVisible    go_types.BoolInt               `json:"isFormVisible"`
	IsRequired       go_types.BoolInt               `json:"isRequired"`
	DisplayOrder     int64                          `json:"displayOrder"`
	CreatedTimestamp a_types.DateTimeTimezoneString `json:"createdTimestamp"`
	UpdatedTimestamp a_types.DateTimeTimezoneString `json:"updatedTimestamp"`
	Links            *Links                         `json:"links"`
}

type GetDealFieldsConfig struct {
	Limit *uint
}

func (service *Service) GetDealFields(getDealFieldsConfig *GetDealFieldsConfig) (*DealFields, *errortools.Error) {
	params := url.Values{}

	dealFields := DealFields{}
	offset := uint(0)
	limit := defaultLimit

	if getDealFieldsConfig != nil {
		if getDealFieldsConfig.Limit != nil {
			limit = *getDealFieldsConfig.Limit
		}
	}

	params.Add("limit", fmt.Sprintf("%v", limit))

	for true {
		params.Set("offset", fmt.Sprintf("%v", offset))

		dealFieldsBatch := DealFields{}

		requestConfig := go_http.RequestConfig{
			URL:           service.url(fmt.Sprintf("dealCustomFieldMeta?%s", params.Encode())),
			ResponseModel: &dealFieldsBatch,
		}

		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		dealFields.DealFields = append(dealFields.DealFields, dealFieldsBatch.DealFields...)

		if len(dealFieldsBatch.DealFields) < int(limit) {
			break
		}
		offset += limit
	}

	return &dealFields, nil
}
