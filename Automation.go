package activecampaign

import (
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Automations struct {
	Automations []Automation `json:"automations"`
	//Meta     FieldValuesMeta `json:"meta"`
}

type Automation struct {
	Name              string          `json:"name"`
	CreateDate        string          `json:"cdate"`
	ModifiedDate      string          `json:"mdate"`
	UserID            string          `json:"userid"`
	Status            string          `json:"status"`
	Entered           string          `json:"entered"`
	Exited            string          `json:"exited"`
	Hidden            string          `json:"hidden"`
	DefaultScreenshot string          `json:"defaultscreenshot"`
	ID                string          `json:"id"`
	Links             AutomationLinks `json:"links"`
}

type AutomationLinks struct {
	Campaigns          string `json:"campaigns"`
	ContactGoals       string `json:"contactGoals"`
	ContactAutomations string `json:"contactAutomations"`
	Blocks             string `json:"blocks"`
	Goals              string `json:"goals"`
	SMS                string `json:"sms"`
	Sitemessages       string `json:"sitemessages"`
}

type GetAutomationsFilter struct {
	Email *string
}

func (service *Service) GetAutomations() (*Automations, *errortools.Error) {
	automations := Automations{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url("automations"),
		ResponseModel: &automations,
	}

	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &automations, nil
}
