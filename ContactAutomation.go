package activecampaign

import (
	"fmt"
	"net/http"
	"net/url"

	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type ContactAutomations struct {
	ContactAutomations []ContactAutomation `json:"contactAutomations"`
	Meta               Meta                `json:"meta"`
}

type ContactAutomation struct {
	ContactID         go_types.Int64String            `json:"contact"`
	SeriesID          go_types.Int64String            `json:"seriesid"`
	StartID           go_types.Int64String            `json:"startid"`
	Status            go_types.Int64String            `json:"status"`
	BatchID           *go_types.String                `json:"batchid"`
	AddDate           a_types.DateTimeTimezoneString  `json:"adddate"`
	ReminderDate      *a_types.DateTimeTimezoneString `json:"remdate"`
	Timespan          *go_types.Int64String           `json:"timespan"`
	LastBlock         *go_types.Int64String           `json:"lastblock"`
	LastLogID         *go_types.Int64String           `json:"lastlogid"`
	LastDate          a_types.DateTimeTimezoneString  `json:"lastdate"`
	InAls             go_types.Int64String            `json:"in_als"`
	CompletedElements int64                           `json:"completedElements"`
	TotalElements     int64                           `json:"totalElements"`
	Completed         go_types.BoolInt                `json:"completed"`
	CompleteValue     int64                           `json:"completeValue"`
	Links             *Links                          `json:"links"`
	ID                go_types.Int64String            `json:"id"`
	AutomationID      go_types.Int64String            `json:"automation"`
}

type GetContactAutomationsConfig struct {
	Limit     *uint64
	Offset    *uint64
	ContactID *int64
}

func (service *Service) GetContactAutomations(getContactAutomationsConfig *GetContactAutomationsConfig) (*ContactAutomations, *errortools.Error) {
	params := url.Values{}

	contactAutomations := ContactAutomations{}
	rowCount := uint64(0)
	limit := defaultLimit

	path := "contactAutomations"

	if getContactAutomationsConfig != nil {
		if getContactAutomationsConfig.ContactID != nil {
			path = fmt.Sprintf("contacts/%v/contactAutomations", *getContactAutomationsConfig.ContactID)
		}
		if getContactAutomationsConfig.Limit != nil {
			limit = *getContactAutomationsConfig.Limit
		}
		if getContactAutomationsConfig.Offset != nil {
			service.nextOffsets.ContactAutomation = *getContactAutomationsConfig.Offset
		}
	}

	params.Add("limit", fmt.Sprintf("%v", limit))

	for {
		params.Set("offset", fmt.Sprintf("%v", service.nextOffsets.ContactAutomation))

		contactAutomationsBatch := ContactAutomations{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			URL:           service.url(fmt.Sprintf("%s?%s", path, params.Encode())),
			ResponseModel: &contactAutomations,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		contactAutomations.ContactAutomations = append(contactAutomations.ContactAutomations, contactAutomationsBatch.ContactAutomations...)

		if len(contactAutomationsBatch.ContactAutomations) < int(limit) {
			service.nextOffsets.ContactAutomation = 0
			break
		}

		service.nextOffsets.ContactAutomation += limit
		rowCount += limit

		if rowCount >= service.maxRowCount {
			return &contactAutomations, nil
		}
	}

	return &contactAutomations, nil
}
