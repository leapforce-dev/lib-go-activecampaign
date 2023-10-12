package activecampaign

import (
	"fmt"
	"net/http"
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
	ContactId   go_types.Int64String           `json:"contact"`
	FieldId     go_types.Int64String           `json:"field"`
	Value       string                         `json:"value"`
	CreatedDate a_types.DateTimeTimezoneString `json:"cdate,omitempty"`
	UpdatedDate a_types.DateTimeTimezoneString `json:"udate,omitempty"`
	CreatedBy   *go_types.String               `json:"created_by,omitempty"`
	UpdatedBy   *go_types.String               `json:"updated_by,omitempty"`
	Id          go_types.Int64String           `json:"id,omitempty"`
	OwnerId     go_types.Int64String           `json:"owner,omitempty"`
	Links       *Links                         `json:"links,omitempty"`
}

type GetContactFieldValuesConfig struct {
	ContactId *int64
	FieldId   *int64
	Value     *string
}

func (service *Service) GetContactFieldValues(getFieldValuesConfig *GetContactFieldValuesConfig) (*ContactFieldValues, bool, *errortools.Error) {
	params := url.Values{}

	fieldValues := ContactFieldValues{}

	path := "fieldValues"

	if getFieldValuesConfig != nil {
		if getFieldValuesConfig.ContactId != nil {
			path = fmt.Sprintf("contacts/%v/fieldValues", *getFieldValuesConfig.ContactId)
		}
		if getFieldValuesConfig.FieldId != nil {
			params.Add("filters[fieldid]", fmt.Sprintf("%v", *getFieldValuesConfig.FieldId))
		}
		if getFieldValuesConfig.Value != nil {
			params.Add("filters[val]", *getFieldValuesConfig.Value)
		}
	}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("%s?%s", path, params.Encode())),
		ResponseModel: &fieldValues,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, true, e
	}

	return &fieldValues, false, nil
}
