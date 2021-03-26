package activecampaign

import (
	"fmt"
	"net/url"

	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type ContactFieldValues struct {
	FieldValues []ContactFieldValue `json:"fieldValues"`
	Meta        Meta                `json:"meta"`
}

type ContactFieldValue struct {
	ContactID   go_types.Int64String           `json:"contact"`
	FieldID     go_types.Int64String           `json:"field"`
	Value       string                         `json:"value"`
	CreatedDate a_types.DateTimeTimezoneString `json:"cdate,omitempty"`
	UpdatedDate a_types.DateTimeTimezoneString `json:"udate,omitempty"`
	CreatedBy   *go_types.String               `json:"created_by,omitempty"`
	UpdatedBy   *go_types.String               `json:"updated_by,omitempty"`
	ID          go_types.Int64String           `json:"id,omitempty"`
	OwnerID     go_types.Int64String           `json:"owner,omitempty"`
	Links       *Links                         `json:"links,omitempty"`
}

type GetFieldValuesConfig struct {
	Limit     *uint
	ContactID *int64
	FieldID   *int64
	Value     *string
}

func (service *Service) GetContactFieldValues(getFieldValuesConfig *GetFieldValuesConfig) (*ContactFieldValues, *errortools.Error) {
	params := url.Values{}

	fieldValues := ContactFieldValues{}
	offset := uint(0)
	limit := defaultLimit

	path := "fieldValues"

	if getFieldValuesConfig != nil {
		if getFieldValuesConfig.ContactID != nil {
			path = fmt.Sprintf("contacts/%v/fieldValues", *getFieldValuesConfig.ContactID)
		}
		if getFieldValuesConfig.FieldID != nil {
			params.Add("filters[fieldid]", fmt.Sprintf("%v", *getFieldValuesConfig.FieldID))
		}
		if getFieldValuesConfig.Value != nil {
			params.Add("filters[val]", *getFieldValuesConfig.Value)
		}
		if getFieldValuesConfig.Limit != nil {
			limit = *getFieldValuesConfig.Limit
		}
	}

	params.Add("limit", fmt.Sprintf("%v", limit))

	for true {
		params.Set("offset", fmt.Sprintf("%v", offset))

		fieldValuesBatch := ContactFieldValues{}

		requestConfig := go_http.RequestConfig{
			URL:           service.url(fmt.Sprintf("%s?%s", path, params.Encode())),
			ResponseModel: &fieldValuesBatch,
		}

		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		fieldValues.FieldValues = append(fieldValues.FieldValues, fieldValuesBatch.FieldValues...)

		if len(fieldValuesBatch.FieldValues) < int(limit) {
			break
		}
		offset += limit
	}

	return &fieldValues, nil
}
