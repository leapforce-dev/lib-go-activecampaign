package activecampaign

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

type Contacts struct {
	Contacts []Contact `json:"contacts"`
	//Meta     FieldValuesMeta `json:"meta"`
}

type Contact struct {
	CreateDate string       `json:"cdate"`
	Email      string       `json:"email"`
	Phone      string       `json:"phone"`
	FirstName  string       `json:"firstName,omitempty"`
	LastName   string       `json:"lastName,omitempty"`
	ID         string       `json:"id"`
	UpdateDate string       `json:"udate"`
	Links      ContactLinks `json:"links"`
}

type ContactLinks struct {
	BounceLogs            string `json:"bounceLogs"`
	ContactAutomations    string `json:"contactAutomations"`
	ContactData           string `json:"contactData"`
	ContactGoals          string `json:"contactGoals"`
	ContactLists          string `json:"contactLists"`
	ContactLogs           string `json:"contactLogs"`
	ContactTags           string `json:"contactTags"`
	ContactDeals          string `json:"contactDeals"`
	Deals                 string `json:"deals"`
	FieldValues           string `json:"fieldValues"`
	GeoIps                string `json:"geoIps"`
	Notes                 string `json:"notes"`
	Organization          string `json:"organization"`
	PlusAppend            string `json:"plusAppend"`
	TrackingLogs          string `json:"trackingLogs"`
	ScoreValues           string `json:"scoreValues"`
	AccountContacts       string `json:"accountContacts"`
	AutomationEntryCounts string `json:"automationEntryCounts"`
}

type ContactSync struct {
	Email       string        `json:"email"`
	FirstName   string        `json:"firstName,omitempty"`
	LastName    string        `json:"lastName,omitempty"`
	Phone       string        `json:"phone,omitempty"`
	FieldValues []CustomField `json:"fieldValues,omitempty"`
}

type ContactSynced struct {
	Email      string       `json:"email"`
	FirstName  string       `json:"firstName"`
	LastName   string       `json:"lastName"`
	Phone      string       `json:"phone"`
	CreateDate string       `json:"cdate"`
	UpdateDate string       `json:"udate"`
	Links      ContactLinks `json:"links"`
	ID         string       `json:"id"`
}

type GetContactsFilter struct {
	Email *string
}

func (ac *ActiveCampaign) GetContacts(filter *GetContactsFilter) (*Contacts, error) {
	urlStr := fmt.Sprintf("%s/contacts", ac.baseURL())

	if filter != nil {
		params := url.Values{}

		if filter.Email != nil {
			params.Add("email", *filter.Email)
		}

		urlStr = fmt.Sprintf("%s?%s", urlStr, params.Encode())
	}

	contacts := Contacts{}

	err := ac.get(urlStr, &contacts)
	if err != nil {
		return nil, err
	}

	return &contacts, nil
}

func (ac *ActiveCampaign) SyncContact(contactCreate ContactSync) (*ContactSynced, error) {
	urlStr := fmt.Sprintf("%s/contact/sync", ac.baseURL())

	b, err := json.Marshal(struct {
		Contact ContactSync `json:"contact"`
	}{
		Contact: contactCreate,
	})
	if err != nil {
		return nil, err
	}

	var contactCreated struct {
		Contact ContactSynced `json:"contact"`
	}

	err = ac.post(urlStr, bytes.NewBuffer(b), &contactCreated)
	if err != nil {
		return nil, err
	}

	return &contactCreated.Contact, nil
}
