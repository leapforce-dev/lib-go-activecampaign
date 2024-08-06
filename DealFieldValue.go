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

type DealFieldValues struct {
	FieldValues []DealFieldValue `json:"dealCustomFieldData"`
	Meta        Meta             `json:"meta"`
}

type DealFieldValue struct {
	DealId     go_types.Int64String `json:"dealId"`
	FieldId    go_types.Int64String `json:"customFieldId"`
	FieldValue *go_types.String     `json:"fieldValue"`
	/*TextValue              *go_types.String        `json:"custom_field_text_value"`
	TextBlob               *go_types.String        `json:"custom_field_text_blob"`
	DateValue              *a_types.DateTimeString `json:"custom_field_date_value"`
	NumberValue            *go_types.Float64String `json:"custom_field_number_value"`
	CurrencyValue          *go_types.Float64String `json:"custom_field_currency_value"`
	CurrencyType           *go_types.String        `json:"custom_field_currency_type"`*/
	CreatedTimestamp a_types.DateTimeString `json:"createdTimestamp"`
	UpdatedTimestamp a_types.DateTimeString `json:"updatedTimestamp"`
	//CreatedBy              *go_types.String        `json:"created_by"`
	//UpdatedBy              *go_types.String        `json:"updated_by"`
	Links                  *Links               `json:"links"`
	Id                     go_types.Int64String `json:"id"`
	DealCustomFieldMetumId go_types.Int64String `json:"dealCustomFieldMetumId"`
}

type GetDealFieldValuesConfig struct {
	Limit   *uint64
	Offset  *uint64
	DealId  *int64
	FieldId *int64
	//Value     *string
}

func (service *Service) GetDealFieldValues(getFieldValuesConfig *GetDealFieldValuesConfig) (*DealFieldValues, bool, *errortools.Error) {
	params := url.Values{}

	fieldValues := DealFieldValues{}
	rowCount := uint64(0)
	limit := defaultLimit

	path := "dealCustomFieldData"

	if getFieldValuesConfig != nil {
		if getFieldValuesConfig.Limit != nil {
			limit = getLimit(*getFieldValuesConfig.Limit)
		}
		if getFieldValuesConfig.Offset != nil {
			service.nextOffsets.Deal = *getFieldValuesConfig.Offset
		}
		if getFieldValuesConfig.DealId != nil {
			params.Add("filters[dealId]", fmt.Sprintf("%v", *getFieldValuesConfig.DealId))
		}
	}
	params.Add("limit", fmt.Sprintf("%v", limit))

	for {
		params.Set("offset", fmt.Sprintf("%v", service.nextOffsets.Deal))

		fmt.Println(service.nextOffsets.Deal)

		fieldValuesBatch := DealFieldValues{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("%s?%s", path, params.Encode())),
			ResponseModel: &fieldValuesBatch,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, true, e
		}

		if getFieldValuesConfig.FieldId == nil {
			fieldValues.FieldValues = append(fieldValues.FieldValues, fieldValuesBatch.FieldValues...)
		} else {
			for _, v := range fieldValuesBatch.FieldValues {
				if v.FieldId.Value() == *getFieldValuesConfig.FieldId {
					fieldValues.FieldValues = append(fieldValues.FieldValues, v)
				}
			}
		}

		if len(fieldValuesBatch.FieldValues) < int(limit) {
			service.nextOffsets.Deal = 0
			break
		}

		service.nextOffsets.Deal += limit
		rowCount += limit

		if rowCount >= service.maxRowCount {
			return &fieldValues, true, nil
		}
	}

	return &fieldValues, false, nil
}
