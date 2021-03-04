package activecampaign

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type ContactAutomations struct {
	ContactAutomations []ContactAutomation `json:"contactAutomations"`
	//Meta     FieldValuesMeta `json:"meta"`
}

type ContactAutomation struct {
	Contact           string                 `json:"contact"`
	SeriesID          string                 `json:"seriesid"`
	StartID           string                 `json:"startid"`
	BatchID           string                 `json:"batchid"`
	AddDate           string                 `json:"adddate"`
	RemDate           string                 `json:"remdate"`
	Timespan          string                 `json:"timespan"`
	LastBlock         string                 `json:"lastblock"`
	LastLogID         string                 `json:"lastlogid"`
	LastDate          string                 `json:"lastdate"`
	InAls             string                 `json:"in_als"`
	CompletedElements int                    `json:"completedElements"`
	TotalElements     int                    `json:"totalElements"`
	Completed         int                    `json:"completed"`
	CompleteValue     int                    `json:"completeValue"`
	ID                string                 `json:"id"`
	Automation        string                 `json:"automation"`
	Links             ContactAutomationLinks `json:"links"`
}

type ContactAutomationLinks struct {
	Automation     string `json:"automation"`
	Contact        string `json:"contact"`
	ContactGoals   string `json:"contactGoals"`
	AutomationLogs string `json:"automationLogs"`
}

func (service *Service) GetContactAutomations(automationID string) (*ContactAutomations, *errortools.Error) {
	contactAutomations := ContactAutomations{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("automations/%s/contactAutomations", automationID)),
		ResponseModel: &contactAutomations,
	}

	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &contactAutomations, nil
}
