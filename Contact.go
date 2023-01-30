package activecampaign

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type ContactsResponse struct {
	ContactAutomations *[]ContactAutomation `json:"contactAutomations"`
	ContactLists       *[]ContactList       `json:"contactLists"`
	ContactTags        *[]ContactTag        `json:"contactTags"`
	FieldValues        *[]ContactFieldValue `json:"fieldValues"`
	Contacts           []Contact            `json:"contacts"`
	Meta               Meta                 `json:"meta"`
}

type Contact struct {
	CreatedDate          a_types.DateTimeTimezoneString  `json:"cdate"`
	Email                string                          `json:"email"`
	Phone                *go_types.String                `json:"phone"`
	FirstName            *go_types.String                `json:"firstName"`
	LastName             *go_types.String                `json:"lastName"`
	SegmentIoId          go_types.Int64String            `json:"segmentio_id"`
	BouncedHard          go_types.Int64String            `json:"bounced_hard"`
	BouncedSoft          go_types.Int64String            `json:"bounced_soft"`
	BouncedDate          *a_types.DateString             `json:"bounced_date"`
	Ip                   *go_types.String                `json:"ip"`
	Ua                   *go_types.String                `json:"ua"`
	Hash                 string                          `json:"hash"`
	SocialdataLastcheck  a_types.DateTimeString          `json:"socialdata_lastcheck"`
	EmailLocal           *go_types.String                `json:"email_local"`
	EmailDomain          *go_types.String                `json:"email_domain"`
	SentCount            go_types.Int64String            `json:"sentcnt"`
	RatingDate           *a_types.DateString             `json:"rating_tstamp"`
	Gravatar             go_types.Int64String            `json:"gravatar"`
	Deleted              go_types.BoolString             `json:"deleted"`
	Anonymized           go_types.BoolString             `json:"anonymized"`
	ADate                a_types.DateTimeTimezoneString  `json:"adate"`
	UpdatedDate          a_types.DateTimeTimezoneString  `json:"udate"`
	EDate                *a_types.DateTimeTimezoneString `json:"edate"`
	DeletedDate          *a_types.DateTimeString         `json:"deleted_at"`
	CreatedUtcTimestamp  a_types.DateTimeString          `json:"created_utc_timestamp"`
	UpdatedUtcTimestamp  a_types.DateTimeString          `json:"updated_utc_timestamp"`
	CreatedTimestamp     a_types.DateTimeString          `json:"created_timestamp"`
	UpdatedTimestamp     a_types.DateTimeString          `json:"updated_timestamp"`
	CreatedBy            *go_types.Int64String           `json:"created_by"`
	UpdatedBy            *go_types.Int64String           `json:"updated_by"`
	EmailEmpty           bool                            `json:"email_empty"`
	ScoreValues          *go_types.Int64Strings          `json:"scoreValues"`
	AccountContacts      *go_types.Int64Strings          `json:"accountContacts"`
	ContactAutomationIds *go_types.Int64Strings          `json:"contactAutomations"`
	ContactListIds       *go_types.Int64Strings          `json:"contactLists"`
	ContactTagIds        *go_types.Int64Strings          `json:"contactTags"`
	FieldValueIds        *go_types.Int64Strings          `json:"fieldValues"`
	Links                *Links                          `json:"links"`
	Id                   go_types.Int64String            `json:"id"`
	OrganizationId       *go_types.Int64String           `json:"organization"`
	ContactAutomations   *[]ContactAutomation            `json:"-"`
	ContactLists         *[]ContactList                  `json:"-"`
	ContactTags          *[]ContactTag                   `json:"-"`
	FieldValues          *[]ContactFieldValue            `json:"-"`
}

type ContactSync struct {
	Email       string              `json:"email"`
	FirstName   string              `json:"firstName,omitempty"`
	LastName    string              `json:"lastName,omitempty"`
	Phone       string              `json:"phone,omitempty"`
	FieldValues []ContactFieldValue `json:"fieldValues,omitempty"`
}

type ContactInclude string

const (
	ContactIncludeContactAutomations ContactInclude = "contactAutomations"
	ContactIncludeContactLists       ContactInclude = "contactLists"
	ContactIncludeContactTags        ContactInclude = "contactTags"
	ContactIncludeFieldValues        ContactInclude = "fieldValues"
)

type GetContactsConfig struct {
	Limit        *uint64
	Offset       *uint64
	Email        *string
	ListId       *int64
	CreatedAfter *time.Time
	UpdatedAfter *time.Time
	Include      *[]ContactInclude
}

func (service *Service) GetContacts(getContactsConfig *GetContactsConfig) (*ContactsResponse, *errortools.Error) {
	params := url.Values{}

	contacts := ContactsResponse{}
	rowCount := uint64(0)
	limit := defaultLimit

	if getContactsConfig != nil {
		if getContactsConfig.Limit != nil {
			limit = *getContactsConfig.Limit
		}
		if getContactsConfig.Offset != nil {
			service.nextOffsets.Contact = *getContactsConfig.Offset
		}
		if getContactsConfig.Email != nil {
			params.Add("email", *getContactsConfig.Email)
		}
		if getContactsConfig.ListId != nil {
			params.Add("listid", fmt.Sprintf("%v", *getContactsConfig.ListId))
		}
		if getContactsConfig.CreatedAfter != nil {
			params.Add("filters[created_after]", (*getContactsConfig.CreatedAfter).Format(timestampLayout))
		}
		if getContactsConfig.UpdatedAfter != nil {
			params.Add("filters[updated_after]", (*getContactsConfig.UpdatedAfter).Format(timestampLayout))
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

	for {
		params.Set("offset", fmt.Sprintf("%v", service.nextOffsets.Contact))

		contactsBatch := ContactsResponse{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("contacts?%s", params.Encode())),
			ResponseModel: &contactsBatch,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		if contactsBatch.ContactAutomations != nil {
			for i, contact := range contactsBatch.Contacts {
				var contactAutomations []ContactAutomation
				for _, contactAutomation := range *contactsBatch.ContactAutomations {
					if contact.Id == contactAutomation.ContactId {
						contactAutomations = append(contactAutomations, contactAutomation)
					}
				}
				contactsBatch.Contacts[i].ContactAutomations = &contactAutomations
			}
		}

		if contactsBatch.ContactLists != nil {
			for i, contact := range contactsBatch.Contacts {
				var contactLists []ContactList
				for _, contactList := range *contactsBatch.ContactLists {
					if contact.Id == contactList.ContactId {
						contactLists = append(contactLists, contactList)
					}
				}
				contactsBatch.Contacts[i].ContactLists = &contactLists
			}
		}

		if contactsBatch.ContactTags != nil {
			for i, contact := range contactsBatch.Contacts {
				var contactTags []ContactTag
				for _, contactTag := range *contactsBatch.ContactTags {
					if contact.Id == contactTag.ContactId {
						contactTags = append(contactTags, contactTag)
					}
				}
				contactsBatch.Contacts[i].ContactTags = &contactTags
			}
		}

		if contactsBatch.FieldValues != nil {
			for i, contact := range contactsBatch.Contacts {
				var fieldValues []ContactFieldValue
				for _, fieldValue := range *contactsBatch.FieldValues {
					if contact.Id == fieldValue.ContactId {
						fieldValues = append(fieldValues, fieldValue)
					}
				}
				contactsBatch.Contacts[i].FieldValues = &fieldValues
			}
		}
		contacts.Contacts = append(contacts.Contacts, contactsBatch.Contacts...)

		if len(contactsBatch.Contacts) < int(limit) {
			service.nextOffsets.Contact = 0
			break
		}

		service.nextOffsets.Contact += limit
		rowCount += limit

		if rowCount >= service.maxRowCount {
			return &contacts, nil
		}
	}

	return &contacts, nil
}

type ContactResponse struct {
	ContactAutomations *[]ContactAutomation         `json:"contactAutomations"`
	ContactLists       *[]ContactList               `json:"contactLists"`
	Deals              *[]Deal                      `json:"deals"`
	FieldValues        *[]ContactFieldValue         `json:"fieldValues"`
	GeoIps             *GeoIps                      `json:"geoIps"`
	Accounts           *[]Account                   `json:"accounts"`
	AccountContacts    *[]AccountContactAssociation `json:"accountContacts"`
	Contact            Contact                      `json:"contact"`
}

type GeoIps []GeoIp

type GeoIp struct {
	Contact    string    `json:"contact"`
	CampaignId string    `json:"campaignid"`
	MessageId  string    `json:"messageid"`
	GeoAddrId  string    `json:"geoaddrid"`
	Ip4        string    `json:"ip4"`
	Tstamp     time.Time `json:"tstamp"`
	GeoAddress string    `json:"geoAddress"`
	Links      struct {
		GeoAddress string `json:"geoAddress"`
	} `json:"links"`
	Id string `json:"id"`
}

func (g *GeoIp) UnmarshalJSON(b []byte) error {
	var returnError = func() error {
		errortools.CaptureError(fmt.Sprintf("Cannot parse '%s' to GeoIp", string(b)))
		return nil
	}

	type geoIp GeoIp

	var g1 geoIp

	err := json.Unmarshal(b, &g1)
	if err != nil {
		var s string
		err = json.Unmarshal(b, &s)
		if err != nil {
			return returnError()
		}
		_, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return returnError()
		}

		*g = GeoIp{
			Id: s,
		}
	} else {
		*g = GeoIp(g1)
	}

	return nil
}

func (service *Service) GetContact(contactId int64) (*ContactResponse, *errortools.Error) {
	contactResponse := ContactResponse{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("contacts/%v", contactId)),
		ResponseModel: &contactResponse,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &contactResponse, nil
}

func (service *Service) SyncContact(contactCreate ContactSync) (*Contact, *errortools.Error) {
	d := struct {
		Contact ContactSync `json:"contact"`
	}{
		Contact: contactCreate,
	}

	var contactCreated struct {
		ContactAutomations *[]ContactAutomation `json:"contactAutomations"`
		ContactLists       *[]ContactList       `json:"contactLists"`
		ContactTags        *[]ContactTag        `json:"contactTags"`
		FieldValues        *[]ContactFieldValue `json:"fieldValues"`
		Contact            Contact              `json:"contact"`
	}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url("contact/sync"),
		BodyModel:     d,
		ResponseModel: &contactCreated,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	contactCreated.Contact.ContactAutomations = contactCreated.ContactAutomations
	contactCreated.Contact.ContactLists = contactCreated.ContactLists
	contactCreated.Contact.ContactTags = contactCreated.ContactTags
	contactCreated.Contact.FieldValues = contactCreated.FieldValues

	return &contactCreated.Contact, nil
}

func (service *Service) CreateContact(contactCreate ContactSync) (*Contact, *errortools.Error) {
	d := struct {
		Contact ContactSync `json:"contact"`
	}{
		Contact: contactCreate,
	}

	var contactCreated struct {
		Contact Contact `json:"contact"`
	}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url("contacts"),
		BodyModel:     d,
		ResponseModel: &contactCreated,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &contactCreated.Contact, nil
}

func (service *Service) UpdateContact(contactId int64, contactUpdate ContactSync) (*Contact, *errortools.Error) {
	d := struct {
		Contact ContactSync `json:"contact"`
	}{
		Contact: contactUpdate,
	}

	var contactUpdated struct {
		Contact Contact `json:"contact"`
	}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPut,
		Url:           service.url(fmt.Sprintf("contacts/%v", contactId)),
		BodyModel:     d,
		ResponseModel: &contactUpdated,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &contactUpdated.Contact, nil
}

func (service *Service) DeleteContact(contactId int64) *errortools.Error {
	requestConfig := go_http.RequestConfig{
		Method: http.MethodDelete,
		Url:    service.url(fmt.Sprintf("contacts/%v", contactId)),
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return e
	}

	return nil
}
