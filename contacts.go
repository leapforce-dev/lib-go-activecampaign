package activecampaign

import (
	"fmt"
	"net/url"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Contacts struct {
	Contacts []Contact `json:"contacts"`
	//Meta     FieldValuesMeta `json:"meta"`
}

type Contact struct {
	CreateDate       string       `json:"cdate"`
	Email            string       `json:"email"`
	Phone            string       `json:"phone"`
	FirstName        string       `json:"firstName,omitempty"`
	LastName         string       `json:"lastName,omitempty"`
	ID               string       `json:"id"`
	UpdateDate       string       `json:"udate"`
	Links            ContactLinks `json:"links"`
	CreatedTimestamp string       `json:"created_timestamp"`
	UpdatedTimestamp string       `json:"updated_timestamp"`
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

type GetContactsConfig struct {
	Limit        *uint
	Email        *string
	ListID       *string
	CreatedAfter *time.Time
	UpdatedAfter *time.Time
}

func (service *Service) GetContacts(filter *GetContactsConfig) (*Contacts, *errortools.Error) {
	params := url.Values{}

	if filter != nil {
		if filter.Email != nil {
			params.Add("email", *filter.Email)
		}
		if filter.ListID != nil {
			params.Add("listid", *filter.ListID)
		}
		if filter.CreatedAfter != nil {
			params.Add("filters[created_after]", (*filter.CreatedAfter).Format(TimestampFormat))
		}
		if filter.UpdatedAfter != nil {
			params.Add("filters[updated_after]", (*filter.UpdatedAfter).Format(TimestampFormat))
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

		contactsBatch := Contacts{}

		requestConfig := go_http.RequestConfig{
			URL:           service.url(fmt.Sprintf("contacts?%s", params.Encode())),
			ResponseModel: &contactsBatch,
		}

		_, _, e := service.get(&requestConfig)
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

func (service *Service) SyncContact(contactCreate ContactSync) (*ContactSynced, *errortools.Error) {
	d := struct {
		Contact ContactSync `json:"contact"`
	}{
		Contact: contactCreate,
	}

	var contactCreated struct {
		Contact ContactSynced `json:"contact"`
	}

	requestConfig := go_http.RequestConfig{
		URL:           service.url("contact/sync"),
		BodyModel:     d,
		ResponseModel: &contactCreated,
	}

	_, _, e := service.post(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &contactCreated.Contact, nil
}

func (service *Service) UpdateContact(contactID string, contactCreate ContactSync) (*ContactSynced, *errortools.Error) {
	d := struct {
		Contact ContactSync `json:"contact"`
	}{
		Contact: contactCreate,
	}

	var contactUpdated struct {
		Contact ContactSynced `json:"contact"`
	}

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("contacts/%s", contactID)),
		BodyModel:     d,
		ResponseModel: &contactUpdated,
	}

	_, _, e := service.post(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &contactUpdated.Contact, nil
}

func (service *Service) DeleteContact(contactID string) *errortools.Error {
	requestConfig := go_http.RequestConfig{
		URL: service.url(fmt.Sprintf("contacts/%s", contactID)),
	}

	_, _, e := service.delete(&requestConfig)
	if e != nil {
		return e
	}

	return nil
}

func (service *Service) GetContactFieldValues(contactID string) (*[]CustomField, *errortools.Error) {
	fieldValues := struct {
		FieldValues []CustomField `json:"fieldValues"`
	}{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("contacts/%s/fieldValues?%s", contactID)),
		ResponseModel: &fieldValues,
	}

	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &fieldValues.FieldValues, nil
}
