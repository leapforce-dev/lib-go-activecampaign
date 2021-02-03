package activecampaign

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
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
	Limit        *uint
	Email        *string
	CreatedAfter *time.Time
}

func (ac *ActiveCampaign) GetContacts(filter *GetContactsFilter) (*Contacts, *errortools.Error) {
	urlStr := fmt.Sprintf("%s/contacts", ac.baseURL())
	params := url.Values{}

	if filter != nil {
		if filter.Email != nil {
			params.Add("email", *filter.Email)
		}
		if filter.CreatedAfter != nil {
			params.Add("filters[created_after]", (*filter.CreatedAfter).Format(time.RFC3339))
		}
	}

	contacts := Contacts{}
	offset := uint(0)
	limit := uint(100)
	if filter.Limit != nil {
		limit = *filter.Limit
	}
	params.Add("limit", fmt.Sprintf("%v", limit))

	for true {
		params.Set("offset", fmt.Sprintf("%v", offset))

		urlStrBatch := fmt.Sprintf("%s?%s", urlStr, params.Encode())
		//fmt.Println(urlStrBatch)

		contactsBatch := Contacts{}

		e := ac.get(urlStrBatch, &contactsBatch)
		if e != nil {
			return nil, e
		}

		contacts.Contacts = append(contacts.Contacts, contactsBatch.Contacts...)

		if len(contactsBatch.Contacts) < int(limit) {
			break
		}
		offset += limit
	}

	return &contacts, nil
}

func (ac *ActiveCampaign) SyncContact(contactCreate ContactSync) (*ContactSynced, *errortools.Error) {
	urlStr := fmt.Sprintf("%s/contact/sync", ac.baseURL())

	b, err := json.Marshal(struct {
		Contact ContactSync `json:"contact"`
	}{
		Contact: contactCreate,
	})
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	var contactCreated struct {
		Contact ContactSynced `json:"contact"`
	}

	e := ac.post(urlStr, bytes.NewBuffer(b), &contactCreated)
	if e != nil {
		return nil, e
	}

	return &contactCreated.Contact, nil
}

func (ac *ActiveCampaign) UpdateContact(contactID string, contactCreate ContactSync) (*ContactSynced, *errortools.Error) {
	urlStr := fmt.Sprintf("%s/contacts/%s", ac.baseURL(), contactID)

	b, err := json.Marshal(struct {
		Contact ContactSync `json:"contact"`
	}{
		Contact: contactCreate,
	})
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	var contactUpdated struct {
		Contact ContactSynced `json:"contact"`
	}

	e := ac.put(urlStr, bytes.NewBuffer(b), &contactUpdated)
	if e != nil {
		return nil, e
	}

	return &contactUpdated.Contact, nil
}

func (ac *ActiveCampaign) DeleteContact(contactID string) *errortools.Error {
	urlStr := fmt.Sprintf("%s/contacts/%s", ac.baseURL(), contactID)

	e := ac.delete(urlStr)
	if e != nil {
		return e
	}

	return nil
}
