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

type Deals struct {
	FieldValues *[]DealFieldValue `json:"dealCustomFieldData"`
	Deals       []Deal            `json:"deals"`
	Meta        MetaNew           `json:"meta"`
}

type Deal struct {
	Hash                       string                          `json:"hash"`
	OwnerID                    go_types.Int64String            `json:"owner"`
	ContactID                  go_types.Int64String            `json:"contact"`
	OrganizationID             *go_types.Int64String           `json:"organization"`
	GroupID                    go_types.Int64String            `json:"group"`
	StageID                    go_types.Int64String            `json:"stage"`
	Title                      string                          `json:"title"`
	Description                string                          `json:"description"`
	Percent                    go_types.Int64String            `json:"percent"`
	CreatedDate                a_types.DateTimeTimezoneString  `json:"cdate"`
	ModifiedDate               a_types.DateTimeTimezoneString  `json:"mdate"`
	NextDate                   *a_types.DateTimeTimezoneString `json:"nextdate"`
	NextTaskID                 *go_types.Int64String           `json:"nexttaskid"`
	Value                      go_types.Int64String            `json:"value"`
	Currency                   string                          `json:"currency"`
	WinProbability             *int64                          `json:"winProbability"`
	WinProbabilityModifiedDate *a_types.DateTimeTimezoneString `json:"winProbabilityMdate"`
	Status                     go_types.Int64String            `json:"status"`
	ActivityCount              go_types.Int64String            `json:"activitycount"`
	NextDealID                 *go_types.Int64String           `json:"nextdealid"`
	EDate                      *a_types.DateTimeString         `json:"edate"`
	FieldValueIDs              *go_types.Int64Strings          `json:"dealCustomFieldData"`
	Links                      *Links                          `json:"links"`
	ID                         go_types.Int64String            `json:"id"`
	IsDisabled                 bool                            `json:"isDisabled"`
	AccountID                  *go_types.Int64String           `json:"account"`
	CustomerAccountID          *go_types.Int64String           `json:"customerAccount"`
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

func (service *Service) GetDeals(getDealsConfig *GetDealsConfig) (*Deals, *errortools.Error) {
	params := url.Values{}

	deals := Deals{}
	rowCount := uint64(0)
	limit := defaultLimit

	if getDealsConfig != nil {
		if getDealsConfig.Limit != nil {
			limit = *getDealsConfig.Limit
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

	for true {
		params.Set("offset", fmt.Sprintf("%v", service.nextOffsets.Deal))

		dealsBatch := Deals{}

		requestConfig := go_http.RequestConfig{
			URL:           service.url(fmt.Sprintf("deals?%s", params.Encode())),
			ResponseModel: &dealsBatch,
		}

		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		if dealsBatch.FieldValues != nil {
			for i, deal := range dealsBatch.Deals {
				var fieldValues []DealFieldValue
				for _, fieldValue := range *dealsBatch.FieldValues {
					if deal.ID == fieldValue.DealID {
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
			return &deals, nil
		}
	}

	return &deals, nil
}
