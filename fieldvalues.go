package activecampaign

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type FieldValues struct {
	FieldValues []FieldValue `json:"fieldValues"`
	//Meta     FieldValuesMeta `json:"meta"`
}

type FieldValue struct {
	Contact     string          `json:"contact"`
	Field       string          `json:"field"`
	Value       string          `json:"value"`
	CreatedDate string          `json:"cdate"`
	UpdatedDate string          `json:"udate"`
	CreatedBy   *string         `json:"created_by"`
	UpdatedBy   *string         `json:"updated_by"`
	ID          string          `json:"id"`
	Owner       string          `json:"owner"`
	Links       FieldValueLinks `json:"links"`
}

type FieldValueLinks struct {
	Owner string `json:"owner"`
	Field string `json:"field"`
}

func (service *Service) GetFieldValues(contactID string) (*FieldValues, *errortools.Error) {
	fieldValues := FieldValues{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("contacts/%s/fieldValues", contactID)),
		ResponseModel: &fieldValues,
	}

	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &fieldValues, nil
}
