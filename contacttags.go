package activecampaign

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
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

func (ac *ActiveCampaign) GetContactTags(contactID string) (*ContactTags, *errortools.Error) {
	urlStr := fmt.Sprintf("%s/contacts/%s/contactTags", ac.baseURL(), contactID)

	contactTags := ContactTags{}

	e := ac.get(urlStr, &contactTags)
	if e != nil {
		return nil, e
	}

	return &contactTags, nil
}
