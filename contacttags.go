package activecampaign

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type ContactTags struct {
	ContactTags []ContactTag `json:"contactTags"`
	//Meta     FieldValuesMeta `json:"meta"`
}

type ContactTag struct {
	Contact          string          `json:"contact"`
	Tag              string          `json:"tag"`
	CreatedDate      string          `json:"cdate"`
	CreatedTimestamp string          `json:"created_timestamp"`
	UpdatedTimestamp string          `json:"updated_timestamp"`
	CreatedBy        string          `json:"created_by"`
	UpdatedBy        string          `json:"updated_by"`
	ID               string          `json:"id"`
	Links            ContactTagLinks `json:"links"`
}

type ContactTagLinks struct {
	Tag     string `json:"tag"`
	Contact string `json:"contact"`
}

func (service *Service) GetContactTags(contactID string) (*ContactTags, *errortools.Error) {
	contactTags := ContactTags{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("contacts/%s/contactTags", contactID)),
		ResponseModel: &contactTags,
	}

	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &contactTags, nil
}
