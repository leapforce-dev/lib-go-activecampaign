package activecampaign

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type Lists struct {
	Lists []List `json:"lists"`
	Meta  Meta   `json:"meta"`
}

type List struct {
	StringId             string                         `json:"stringid"`
	UserId               go_types.Int64String           `json:"userid"`
	Name                 string                         `json:"name"`
	CreatedDate          a_types.DateTimeTimezoneString `json:"cdate"`
	UseTracking          go_types.BoolString            `json:"p_use_tracking"`
	UseAnalyticsRead     go_types.BoolString            `json:"p_use_analytics_read"`
	UseAnalyticsLink     go_types.BoolString            `json:"p_use_analytics_link"`
	UseTwitter           go_types.BoolString            `json:"p_use_twitter"`
	UseFacebook          go_types.BoolString            `json:"p_use_facebook"`
	EmbedImage           go_types.BoolString            `json:"p_embed_image"`
	UseCaptcha           go_types.BoolString            `json:"p_use_captcha"`
	SendLastBroadcast    go_types.BoolString            `json:"send_last_broadcast"`
	Private              go_types.BoolString            `json:"private"`
	AnalyticsDomains     json.RawMessage                `json:"analytics_domains"`
	AnalyticsSource      *go_types.String               `json:"analytics_source"`
	AnalyticsUA          *go_types.String               `json:"analytics_ua"`
	TwitterToken         *go_types.String               `json:"twitter_token"`
	TwitterTokenSecret   *go_types.String               `json:"twitter_token_secret"`
	FacebookSession      json.RawMessage                `json:"facebook_session"`
	CarbonCopy           json.RawMessage                `json:"carboncopy"`
	SubscriptionNotify   json.RawMessage                `json:"subscription_notify"`
	UnsubscriptionNotify json.RawMessage                `json:"unsubscription_notify"`
	RequireName          go_types.BoolString            `json:"require_name"`
	GetUnsubscribeReason go_types.Int64String           `json:"get_unsubscribe_reason"`
	ToName               *go_types.String               `json:"to_name"`
	OptInOptOut          go_types.BoolString            `json:"optinoptout"`
	SenderName           *go_types.String               `json:"sender_name"`
	SenderAddr1          *go_types.String               `json:"sender_addr1"`
	SenderAddr2          *go_types.String               `json:"sender_addr2"`
	SenderCity           *go_types.String               `json:"sender_city"`
	SenderState          *go_types.String               `json:"sender_state"`
	SenderZip            *go_types.String               `json:"sender_zip"`
	SenderCountry        *go_types.String               `json:"sender_country"`
	SenderPhone          *go_types.String               `json:"sender_phone"`
	SenderUrl            *go_types.String               `json:"sender_url"`
	SenderReminder       *go_types.String               `json:"sender_reminder"`
	FullAddress          *go_types.String               `json:"fulladdress"`
	OptInMessageId       *go_types.Int64String          `json:"optinmessageid"`
	OptOutConf           *go_types.Int64String          `json:"optoutconf"`
	DeleteStamp          json.RawMessage                `json:"deletestamp"`
	UpdatedDate          *a_types.DateTimeString        `json:"udate"`
	CreatedTimestamp     a_types.DateTimeString         `json:"created_timestamp"`
	UpdatedTimestamp     a_types.DateTimeString         `json:"updated_timestamp"`
	CreatedBy            *go_types.Int64String          `json:"created_by"`
	UpdatedBy            *go_types.Int64String          `json:"updated_by"`
	Links                *Links                         `json:"links"`
	Id                   go_types.Int64String           `json:"id"`
}

type GetListsConfig struct {
	Limit  *uint64
	Offset *uint64
	Name   *string
}

func (service *Service) GetLists(getListsConfig *GetListsConfig) (*Lists, bool, *errortools.Error) {
	params := url.Values{}

	lists := Lists{}
	rowCount := uint64(0)
	limit := defaultLimit

	if getListsConfig != nil {
		if getListsConfig.Name != nil {
			params.Add("filters[name]", *getListsConfig.Name)
		}
		if getListsConfig.Limit != nil {
			limit = getLimit(*getListsConfig.Limit)
		}
		if getListsConfig.Offset != nil {
			service.nextOffsets.List = *getListsConfig.Offset
		}
	}
	params.Set("limit", fmt.Sprintf("%v", limit))

	for {
		params.Set("offset", fmt.Sprintf("%v", service.nextOffsets.List))

		listsBatch := Lists{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("lists?%s", params.Encode())),
			ResponseModel: &listsBatch,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, false, e
		}

		lists.Lists = append(lists.Lists, listsBatch.Lists...)

		if len(listsBatch.Lists) < int(limit) {
			service.nextOffsets.List = 0
			break
		}

		service.nextOffsets.List += limit
		rowCount += limit

		if rowCount >= service.maxRowCount {
			return &lists, true, nil
		}
	}

	return &lists, false, nil
}
