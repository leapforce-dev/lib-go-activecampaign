package activecampaign

import (
	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type ContactList struct {
	ContactID               go_types.Int64String            `json:"contact"`
	ListID                  go_types.Int64String            `json:"list"`
	FormID                  *go_types.Int64String           `json:"form"`
	SeriesID                *go_types.Int64String           `json:"seriesid"`
	SubscribedDate          a_types.DateTimeTimezoneString  `json:"sdate"`
	UnsubscribedDate        *a_types.DateTimeTimezoneString `json:"udate"`
	Status                  go_types.Int64String            `json:"status"`
	Responder               go_types.BoolString             `json:"responder"`
	Sync                    go_types.Int64String            `json:"sync"`
	UnsubscribeReason       *go_types.String                `json:"unsubreason"`
	CampaignID              *go_types.Int64String           `json:"campaign"`
	MessageID               *go_types.Int64String           `json:"message"`
	FirstName               *go_types.String                `json:"first_name"`
	LastName                *go_types.String                `json:"last_name"`
	IPSubcribe              *go_types.String                `json:"ip4sub"`
	SourceID                *go_types.Int64String           `json:"sourceid"`
	AutosyncLog             *go_types.String                `json:"autosyncLog"`
	IPLast                  *go_types.Int64String           `json:"ip4_last"`
	IPUnsubscribe           *go_types.Int64String           `json:"ip4Unsub"`
	CreatedTimestamp        a_types.DateTimeString          `json:"created_timestamp"`
	UpdatedTimestamp        a_types.DateTimeString          `json:"updated_timestamp"`
	CreatedBy               *go_types.String                `json:"created_by"`
	UpdatedBy               *go_types.String                `json:"updated_by"`
	UnsubscribeAutomationID *go_types.Int64String           `json:"unsubscribeAutomation"`
	Links                   Links                           `json:"links"`
	ID                      go_types.Int64String            `json:"id"`
	AutomationID            *go_types.Int64String           `json:"automation"`
}

func (service *Service) Subscribe(contactID int64, listID int64) *errortools.Error {
	return service.setContactLists(contactID, listID, 1)
}

func (service *Service) Unsubscribe(contactID int64, listID int64) *errortools.Error {
	return service.setContactLists(contactID, listID, 2)
}

func (service *Service) setContactLists(listID int64, contactID int64, status int64) *errortools.Error {
	type contactList struct {
		ListID    int64 `json:"list"`
		ContactID int64 `json:"contact"`
		Status    int64 `json:"status"`
	}

	type data struct {
		ContactList contactList `json:"contactList"`
	}

	d := data{
		contactList{
			ListID:    listID,
			ContactID: contactID,
			Status:    status,
		},
	}

	requestConfig := go_http.RequestConfig{
		URL:       service.url("contactLists"),
		BodyModel: d,
	}

	_, _, e := service.post(&requestConfig)
	if e != nil {
		return e
	}

	return nil
}
