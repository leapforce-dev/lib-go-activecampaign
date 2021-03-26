package activecampaign

import (
	"fmt"
	"net/url"

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
	DefaultScreenshot *go_types.String               `json:"defaultscreenshot"`
	ID                go_types.Int64String           `json:"id"`
	Links             *Links                         `json:"links"`
}

type GetAutomationsConfig struct {
	Limit *uint
}

func (service *Service) GetAutomations(getAutomationsConfig *GetAutomationsConfig) (*Automations, *errortools.Error) {
	params := url.Values{}

	automations := Automations{}
	offset := uint(0)
	limit := defaultLimit

	if getAutomationsConfig != nil {
		if getAutomationsConfig.Limit != nil {
			limit = *getAutomationsConfig.Limit
		}
	}

	params.Add("limit", fmt.Sprintf("%v", limit))

	for true {
		params.Set("offset", fmt.Sprintf("%v", offset))

		automationsBatch := Automations{}

		requestConfig := go_http.RequestConfig{
			URL:           service.url(fmt.Sprintf("automations?%s", params.Encode())),
			ResponseModel: &automationsBatch,
		}

		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		automations.Automations = append(automations.Automations, automationsBatch.Automations...)

		if len(automationsBatch.Automations) < int(limit) {
			break
		}
		offset += limit
	}

	return &automations, nil
}
