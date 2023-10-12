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

type DealGroups struct {
	DealStages *[]DealStage `json:"dealStages"`
	DealGroups []DealGroup  `json:"dealGroups"`
	Meta       Meta         `json:"meta"`
}

type DealGroup struct {
	Title                        string                         `json:"title"`
	Currency                     string                         `json:"currency"`
	AutoAssign                   go_types.BoolString            `json:"autoassign"`
	AllUsers                     go_types.BoolString            `json:"allusers"`
	AllGroups                    go_types.BoolString            `json:"allgroups"`
	WinProbabilityInitializeDate *a_types.DateTimeString        `json:"win_probability_initialize_date"`
	CreatedDate                  a_types.DateTimeTimezoneString `json:"cdate"`
	UpdatedDate                  a_types.DateTimeTimezoneString `json:"udate"`
	StageIds                     *go_types.Int64Strings         `json:"stages"`
	Links                        *Links                         `json:"links"`
	Id                           go_types.Int64String           `json:"id"`
	Stages                       *[]DealStage                   `json:"-"`
}

type GetDealGroupsConfig struct {
	Limit          *uint64
	Offset         *uint64
	Title          *string
	HaveStages     *bool
	OrderByTitle   *OrderByDirection
	OrderByPopular *OrderByDirection
}

func (service *Service) GetDealGroups(getDealGroupsConfig *GetDealGroupsConfig) (*DealGroups, bool, *errortools.Error) {
	params := url.Values{}

	dealGroups := DealGroups{}
	rowCount := uint64(0)
	limit := defaultLimit

	if getDealGroupsConfig != nil {
		if getDealGroupsConfig.Limit != nil {
			limit = *getDealGroupsConfig.Limit
		}
		if getDealGroupsConfig.Offset != nil {
			service.nextOffsets.DealGroup = *getDealGroupsConfig.Offset
		}
		if getDealGroupsConfig.Title != nil {
			params.Add("filters[title]", *getDealGroupsConfig.Title)
		}
		if getDealGroupsConfig.HaveStages != nil {
			haveStages := 0
			if *getDealGroupsConfig.HaveStages {
				haveStages = 1
			}
			params.Add("filters[have_stages]", fmt.Sprintf("%v", haveStages))
		}
		if getDealGroupsConfig.OrderByTitle != nil {
			params.Add("orders[title]", string(*getDealGroupsConfig.OrderByTitle))
		}
		if getDealGroupsConfig.OrderByPopular != nil {
			params.Add("orders[popular]", string(*getDealGroupsConfig.OrderByPopular))
		}
	}
	params.Add("limit", fmt.Sprintf("%v", limit))

	for {
		params.Set("offset", fmt.Sprintf("%v", service.nextOffsets.DealGroup))

		dealGroupsBatch := DealGroups{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("dealGroups?%s", params.Encode())),
			ResponseModel: &dealGroupsBatch,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, false, e
		}

		if dealGroupsBatch.DealStages != nil {
			for i, dealGroup := range dealGroupsBatch.DealGroups {
				var dealStages []DealStage
				for _, dealStage := range *dealGroupsBatch.DealStages {
					if dealGroup.Id == dealStage.GroupId {
						dealStages = append(dealStages, dealStage)
					}
				}
				dealGroupsBatch.DealGroups[i].Stages = &dealStages
			}
		}
		dealGroups.DealGroups = append(dealGroups.DealGroups, dealGroupsBatch.DealGroups...)

		if len(dealGroupsBatch.DealGroups) < int(limit) {
			service.nextOffsets.DealGroup = 0
			break
		}

		service.nextOffsets.DealGroup += limit
		rowCount += limit

		if rowCount >= service.maxRowCount {
			return &dealGroups, true, nil
		}
	}

	return &dealGroups, false, nil
}
