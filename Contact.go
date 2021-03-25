package activecampaign

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type Contacts struct {
	ContactLists *[]ContactList       `json:"contactLists"`
	FieldValues  *[]FieldValueContact `json:"fieldValues"`
	Contacts     []Contact            `json:"contacts"`
	Meta         Meta                 `json:"meta"`
}

type Contact struct {
	CreatedDate         a_types.DateTimeTimezoneString  `json:"cdate"`
	Email               string                          `json:"email"`
	Phone               *string                         `json:"phone"`
	FirstName           *string                         `json:"firstName"`
	LastName            *string                         `json:"lastName"`
	SegmentIOID         go_types.Int64String            `json:"segmentio_id"`
	BouncedHard         go_types.Int64String            `json:"bounced_hard"`
	BouncedSoft         go_types.Int64String            `json:"bounced_soft"`
	BouncedDate         *a_types.DateString             `json:"bounced_date"`
	IP                  *string                         `json:"ip"`
	UA                  *string                         `json:"ua"`
	Hash                string                          `json:"hash"`
	SocialdataLastcheck a_types.DateTimeString          `json:"socialdata_lastcheck"`
	EmailLocal          *string                         `json:"email_local"`
	EmailDomain         *string                         `json:"email_domain"`
	SentCount           go_types.Int64String            `json:"sentcnt"`
	RatingDate          *a_types.DateString             `json:"rating_tstamp"`
	Gravatar            go_types.Int64String            `json:"gravatar"`
	Deleted             go_types.BoolString             `json:"deleted"`
	Anonymized          go_types.BoolString             `json:"anonymized"`
	ADate               a_types.DateTimeTimezoneString  `json:"adate"`
	UpdatedDate         a_types.DateTimeTimezoneString  `json:"udate"`
	EDate               *a_types.DateTimeTimezoneString `json:"edate"`
	DeletedDate         *a_types.DateTimeString         `json:"deleted_at"`
	CreatedUTCTimestamp a_types.DateTimeString          `json:"created_utc_timestamp"`
	UpdatedUTCTimestamp a_types.DateTimeString          `json:"updated_utc_timestamp"`
	CreatedTimestamp    a_types.DateTimeString          `json:"created_timestamp"`
	UpdatedTimestamp    a_types.DateTimeString          `json:"updated_timestamp"`
	CreatedBy           *go_types.Int64String           `json:"created_by"`
	UpdatedBy           *go_types.Int64String           `json:"updated_by"`
	EmailEmpty          bool                            `json:"email_empty"`
	ScoreValues         *go_types.Int64Strings          `json:"scoreValues"`
	AccountContacts     *go_types.Int64Strings          `json:"accountContacts"`
	ContactListIDs      *go_types.Int64Strings          `json:"contactLists"`
	FieldValueIDs       *go_types.Int64Strings          `json:"fieldValues"`
	Links               *Links                          `json:"links"`
	ID                  go_types.Int64String            `json:"id"`
	Organization        *go_types.Int64String           `json:"organization"`
	ContactLists        *[]ContactList                  `json:"-"`
	FieldValues         *[]FieldValueContact            `json:"-"`
}

type ContactSync struct {
	Email       string        `json:"email"`
	FirstName   string        `json:"firstName,omitempty"`
	LastName    string        `json:"lastName,omitempty"`
	Phone       string        `json:"phone,omitempty"`
	FieldValues []CustomField `json:"fieldValues,omitempty"`
}

type ContactSynced struct {
	Email      string            `json:"email"`
	FirstName  string            `json:"firstName"`
	LastName   string            `json:"lastName"`
	Phone      string            `json:"phone"`
	CreateDate string            `json:"cdate"`
	UpdateDate string            `json:"udate"`
	Links      map[string]string `json:"links"`
	ID         string            `json:"id"`
}

type ContactInclude string

const (
	ContactIncludeFieldValues  ContactInclude = "fieldValues"
	ContactIncludeContactLists ContactInclude = "contactLists"
)

type GetContactsConfig struct {
	Limit        *uint
	Email        *string
	ListID       *string
	CreatedAfter *time.Time
	UpdatedAfter *time.Time
	Include      *[]ContactInclude
}

func (service *Service) GetContacts(getContactsConfig *GetContactsConfig) (*Contacts, *errortools.Error) {
	params := url.Values{}

	contacts := Contacts{}
	offset := uint(0)
	limit := defaultLimit

	if getContactsConfig != nil {
		if getContactsConfig.Limit != nil {
			limit = *getContactsConfig.Limit
		}
		if getContactsConfig.Email != nil {
			params.Add("email", *getContactsConfig.Email)
		}
		if getContactsConfig.ListID != nil {
			params.Add("listid", *getContactsConfig.ListID)
		}
		if getContactsConfig.CreatedAfter != nil {
			params.Add("filters[created_after]", (*getContactsConfig.CreatedAfter).Format(TimestampFormat))
		}
		if getContactsConfig.UpdatedAfter != nil {
			params.Add("filters[updated_after]", (*getContactsConfig.UpdatedAfter).Format(TimestampFormat))
		}
		if getContactsConfig.Include != nil {
			includes := []string{}
			for _, include := range *getContactsConfig.Include {
				includes = append(includes, string(include))
			}

			params.Add("include", strings.Join(includes, ","))
		}
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

		if contactsBatch.ContactLists != nil {
			for i, contact := range contactsBatch.Contacts {
				var contactLists []ContactList
				for _, contactList := range *contactsBatch.ContactLists {
					if contact.ID == contactList.ContactID {
						contactLists = append(contactLists, contactList)
					}
				}
				contactsBatch.Contacts[i].ContactLists = &contactLists
			}
		}
		if contactsBatch.FieldValues != nil {
			for i, contact := range contactsBatch.Contacts {
				var fieldValues []FieldValueContact
				for _, fieldValue := range *contactsBatch.FieldValues {
					if contact.ID == fieldValue.ContactID {
						fieldValues = append(fieldValues, fieldValue)
					}
				}
				contactsBatch.Contacts[i].FieldValues = &fieldValues
			}
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

func (service *Service) GetContactFieldValues(contactID int64) (*[]CustomField, *errortools.Error) {
	fieldValues := struct {
		FieldValues []CustomField `json:"fieldValues"`
	}{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("contacts/%v/fieldValues", contactID)),
		ResponseModel: &fieldValues,
	}

	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &fieldValues.FieldValues, nil
}
