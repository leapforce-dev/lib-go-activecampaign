package activecampaign

import (
	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type Automations struct {
	Automations []Automation `json:"automations"`
	Meta        Meta         `json:"meta"`
}

type Automation struct {
	Name              string                         `json:"name"`
	CreatedDate       a_types.DateTimeTimezoneString `json:"cdate"`
	ModifiedDate      a_types.DateTimeTimezoneString `json:"mdate"`
	UserID            go_types.Int64String           `json:"userid"`
	Status            go_types.Int64String           `json:"status"`
	Entered           go_types.Int64String           `json:"entered"`
	Exited            go_types.Int64String           `json:"exited"`
	Hidden            go_types.Int64String           `json:"hidden"`
	DefaultScreenshot *string                        `json:"defaultscreenshot"`
	ID                go_types.Int64String           `json:"id"`
	Links             AutomationLinks                `json:"links"`
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
