package activecampaign

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type Deals struct {
	FieldValues *[]DealFieldValue `json:"dealCustomFieldData"`
	Deals       []Deal            `json:"deals"`
	Meta        MetaNew           `json:"meta"`
}

type Deal struct {
	Hash                       string                          `json:"hash"`
	OwnerId                    go_types.Int64String            `json:"owner"`
	ContactId                  go_types.Int64String            `json:"contact"`
	OrganizationId             *go_types.Int64String           `json:"organization"`
	GroupId                    go_types.Int64String            `json:"group"`
	StageId                    go_types.Int64String            `json:"stage"`
	Title                      string                          `json:"title"`
	Description                string                          `json:"description"`
	Percent                    go_types.Int64String            `json:"percent"`
	CreatedDate                a_types.DateTimeTimezoneString  `json:"cdate"`
	ModifiedDate               a_types.DateTimeTimezoneString  `json:"mdate"`
	NextDate                   *a_types.DateTimeTimezoneString `json:"nextdate"`
	NextTaskId                 *go_types.Int64String           `json:"nexttaskid"`
	Value                      go_types.Int64String            `json:"value"`
	Currency                   string                          `json:"currency"`
	WinProbability             *int64                          `json:"winProbability"`
	WinProbabilityModifiedDate *a_types.DateTimeTimezoneString `json:"winProbabilityMdate"`
	Status                     go_types.Int64String            `json:"status"`
	ActivityCount              go_types.Int64String            `json:"activitycount"`
	NextDealId                 *go_types.Int64String           `json:"nextdealid"`
	EDate                      *a_types.DateTimeString         `json:"edate"`
	FieldValueIds              *go_types.Int64Strings          `json:"dealCustomFieldData"`
	Links                      *Links                          `json:"links"`
	Id                         go_types.Int64String            `json:"id"`
	IsDisabled                 bool                            `json:"isDisabled"`
	AccountId                  *go_types.Int64String           `json:"account"`
	CustomerAccountId          *go_types.Int64String           `json:"customerAccount"`
	FieldValues                *[]DealFieldValue               `json:"-"`
}

type DealInclude string

const (
	DealIncludeFieldValues DealInclude = "dealCustomFieldData"
)

type GetDealsConfig struct {
	Limit        *uint64
	Offset       *uint64
	CreatedAfter *time.Time
	UpdatedAfter *time.Time
	Include      *[]DealInclude
}

func (service *Service) GetDeals(getDealsConfig *GetDealsConfig) (*Deals, bool, *errortools.Error) {
	params := url.Values{}

	deals := Deals{}
	rowCount := uint64(0)
	limit := defaultLimit

	if getDealsConfig != nil {
		if getDealsConfig.Limit != nil {
			limit = getLimit(*getDealsConfig.Limit)
		}
		if getDealsConfig.Offset != nil {
			service.nextOffsets.Deal = *getDealsConfig.Offset
		}
		if getDealsConfig.CreatedAfter != nil {
			params.Add("filters[created_after]", (*getDealsConfig.CreatedAfter).Format(timestampLayout))
		}
		if getDealsConfig.UpdatedAfter != nil {
			params.Add("filters[updated_after]", (*getDealsConfig.UpdatedAfter).Format(timestampLayout))
		}
		if getDealsConfig.Include != nil {
			includes := []string{}
			for _, include := range *getDealsConfig.Include {
				includes = append(includes, string(include))
			}

			params.Add("include", strings.Join(includes, ","))
		}
	}
	params.Add("limit", fmt.Sprintf("%v", limit))

	for {
		params.Set("offset", fmt.Sprintf("%v", service.nextOffsets.Deal))

		dealsBatch := Deals{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("deals?%s", params.Encode())),
			ResponseModel: &dealsBatch,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, false, e
		}

		if dealsBatch.FieldValues != nil {
			for i, deal := range dealsBatch.Deals {
				var fieldValues []DealFieldValue
				for _, fieldValue := range *dealsBatch.FieldValues {
					if deal.Id == fieldValue.DealId {
						fieldValues = append(fieldValues, fieldValue)
					}
				}
				dealsBatch.Deals[i].FieldValues = &fieldValues
			}
		}
		deals.Deals = append(deals.Deals, dealsBatch.Deals...)

		if len(dealsBatch.Deals) < int(limit) {
			service.nextOffsets.Deal = 0
			break
		}

		service.nextOffsets.Deal += limit
		rowCount += limit

		if rowCount >= service.maxRowCount {
			return &deals, true, nil
		}
	}

	return &deals, false, nil
}

func (service *Service) GetContactDeals(contactId int64) (*Deals, bool, *errortools.Error) {
	deals := Deals{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("contacts/%v/deals", contactId)),
		ResponseModel: &deals,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, false, e
	}

	return &deals, false, nil
}

type DealCreate struct {
	Id             *string                 `json:"id,omitempty"`
	OwnerId        *string                 `json:"owner,omitempty"`
	ContactId      *string                 `json:"contact,omitempty"`
	GroupId        *string                 `json:"group,omitempty"`
	StageId        *string                 `json:"stage,omitempty"`
	Title          *string                 `json:"title,omitempty"`
	Description    *string                 `json:"description,omitempty"`
	Value          *int64                  `json:"value,omitempty"`
	Currency       *string                 `json:"currency,omitempty"`
	WinProbability *int64                  `json:"winProbability,omitempty"`
	Status         *int64                  `json:"status"`
	Fields         *[]DealFieldValueCreate `json:"fields,omitempty"`
}

type DealFieldValueCreate struct {
	CustomFieldId int64           `json:"customFieldId"`
	FieldValue    json.RawMessage `json:"fieldValue"`
	FieldCurrency *string         `json:"fieldCurrency,omitempty"`
}

func (service *Service) CreateDeal(deal *DealCreate) (*DealCreate, *errortools.Error) {
	if deal == nil {
		return nil, nil
	}

	d := struct {
		Deal DealCreate `json:"deal"`
	}{
		Deal: *deal,
	}

	var dealCreated struct {
		Deal DealCreate `json:"deal"`
	}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url("deals"),
		BodyModel:     d,
		ResponseModel: &dealCreated,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &dealCreated.Deal, nil
}

func (service *Service) UpdateDeal(dealId int64, dealUpdate DealCreate) (*Deal, *errortools.Error) {
	d := struct {
		Deal DealCreate `json:"deal"`
	}{
		Deal: dealUpdate,
	}

	var dealUpdated struct {
		Deal Deal `json:"deal"`
	}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPut,
		Url:           service.url(fmt.Sprintf("deals/%v", dealId)),
		BodyModel:     d,
		ResponseModel: &dealUpdated,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &dealUpdated.Deal, nil
}
