package activecampaign

import (
	"fmt"
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

func (ac *ActiveCampaign) GetContactTags(contactID string) (*ContactTags, error) {
	urlStr := fmt.Sprintf("%s/contacts/%s/contactTags", ac.baseURL(), contactID)

	contactTags := ContactTags{}

	err := ac.get(urlStr, &contactTags)
	if err != nil {
		return nil, err
	}

	return &contactTags, nil
}
