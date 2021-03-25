package activecampaign

import (
	"fmt"

	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type FieldValues struct {
	FieldValues []FieldValue `json:"fieldValues"`
	Meta        Meta         `json:"meta"`
}

type FieldValue struct {
	ContactID   go_types.Int64String           `json:"contact"`
	FieldID     go_types.Int64String           `json:"field"`
	Value       string                         `json:"value"`
	CreatedDate a_types.DateTimeTimezoneString `json:"cdate"`
	UpdatedDate a_types.DateTimeTimezoneString `json:"udate"`
	CreatedBy   *go_types.String               `json:"created_by"`
	UpdatedBy   *go_types.String               `json:"updated_by"`
	ID          go_types.Int64String           `json:"id"`
	OwnerID     go_types.Int64String           `json:"owner"`
	Links       Links                          `json:"links"`
}

func (service *Service) GetFieldValues(contactID int64) (*FieldValues, *errortools.Error) {
	fieldValues := FieldValues{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("contacts/%v/fieldValues", contactID)),
		ResponseModel: &fieldValues,
	}

	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &fieldValues, nil
}
