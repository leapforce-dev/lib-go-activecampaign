package activecampaign

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
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

func (ac *ActiveCampaign) GetFieldValues(contactID string) (*FieldValues, *errortools.Error) {
	urlStr := fmt.Sprintf("%s/contacts/%s/fieldValues", ac.baseURL(), contactID)
	//fmt.Println(urlStr)

	fieldValues := FieldValues{}

	e := ac.get(urlStr, &fieldValues)
	if e != nil {
		return nil, e
	}

	return &fieldValues, nil
}
