package activecampaign

import (
	"net/http"

	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type ContactList struct {
	ContactId               go_types.Int64String            `json:"contact"`
	ListId                  go_types.Int64String            `json:"list"`
	FormId                  *go_types.Int64String           `json:"form"`
	SeriesId                *go_types.Int64String           `json:"seriesid"`
	SubscribedDate          a_types.DateTimeTimezoneString  `json:"sdate"`
	UnsubscribedDate        *a_types.DateTimeTimezoneString `json:"udate"`
	Status                  go_types.Int64String            `json:"status"`
	Responder               go_types.BoolString             `json:"responder"`
	Sync                    go_types.Int64String            `json:"sync"`
	UnsubscribeReason       *go_types.String                `json:"unsubreason"`
	CampaignId              *go_types.Int64String           `json:"campaign"`
	MessageId               *go_types.Int64String           `json:"message"`
	FirstName               *go_types.String                `json:"first_name"`
	LastName                *go_types.String                `json:"last_name"`
	IPSubcribe              *go_types.String                `json:"ip4sub"`
	SourceId                *go_types.Int64String           `json:"sourceid"`
	AutosyncLog             *go_types.String                `json:"autosyncLog"`
	IPLast                  *go_types.Int64String           `json:"ip4_last"`
	IPUnsubscribe           *go_types.Int64String           `json:"ip4Unsub"`
	CreatedTimestamp        a_types.DateTimeString          `json:"created_timestamp"`
	UpdatedTimestamp        a_types.DateTimeString          `json:"updated_timestamp"`
	CreatedBy               *go_types.String                `json:"created_by"`
	UpdatedBy               *go_types.String                `json:"updated_by"`
	UnsubscribeAutomationId *go_types.Int64String           `json:"unsubscribeAutomation"`
	Links                   *Links                          `json:"links"`
	Id                      go_types.Int64String            `json:"id"`
	AutomationId            *go_types.Int64String           `json:"automation"`
}

func (service *Service) Subscribe(listId int64, contactId int64) *errortools.Error {
	return service.setContactLists(listId, contactId, 1)
}

func (service *Service) Unsubscribe(listId int64, contactId int64) *errortools.Error {
	return service.setContactLists(listId, contactId, 2)
}

func (service *Service) setContactLists(listId int64, contactId int64, status int64) *errortools.Error {
	type contactList struct {
		ListId    int64 `json:"list"`
		ContactId int64 `json:"contact"`
		Status    int64 `json:"status"`
	}

	type data struct {
		ContactList contactList `json:"contactList"`
	}

	d := data{
		contactList{
			ListId:    listId,
			ContactId: contactId,
			Status:    status,
		},
	}

	requestConfig := go_http.RequestConfig{
		Method:    http.MethodPost,
		Url:       service.url("contactLists"),
		BodyModel: d,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return e
	}

	return nil
}
