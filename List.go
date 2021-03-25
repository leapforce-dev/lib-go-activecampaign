package activecampaign

import (
	"encoding/json"
	"fmt"
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
	StringID             string                         `json:"stringid"`
	UserID               go_types.Int64String           `json:"userid"`
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
	AnalyticsSource      *string                        `json:"analytics_source"`
	AnalyticsUA          *string                        `json:"analytics_ua"`
	TwitterToken         *string                        `json:"twitter_token"`
	TwitterTokenSecret   *string                        `json:"twitter_token_secret"`
	FacebookSession      json.RawMessage                `json:"facebook_session"`
	CarbonCopy           json.RawMessage                `json:"carboncopy"`
	SubscriptionNotify   json.RawMessage                `json:"subscription_notify"`
	UnsubscriptionNotify json.RawMessage                `json:"unsubscription_notify"`
	RequireName          go_types.BoolString            `json:"require_name"`
	GetUnsubscribeReason go_types.Int64String           `json:"get_unsubscribe_reason"`
	ToName               *string                        `json:"to_name"`
	OptInOptOut          go_types.BoolString            `json:"optinoptout"`
	SenderName           *string                        `json:"sender_name"`
	SenderAddr1          *string                        `json:"sender_addr1"`
	SenderAddr2          *string                        `json:"sender_addr2"`
	SenderCity           *string                        `json:"sender_city"`
	SenderState          *string                        `json:"sender_state"`
	SenderZip            *string                        `json:"sender_zip"`
	SenderCountry        *string                        `json:"sender_country"`
	SenderPhone          *string                        `json:"sender_phone"`
	SenderURL            *string                        `json:"sender_url"`
	SenderReminder       *string                        `json:"sender_reminder"`
	FullAddress          *string                        `json:"fulladdress"`
	OptInMessageID       *go_types.Int64String          `json:"optinmessageid"`
	OptOutConf           *go_types.Int64String          `json:"optoutconf"`
	DeleteStamp          json.RawMessage                `json:"deletestamp"`
	UpdatedDate          *a_types.DateTimeString        `json:"udate"`
	CreatedTimestamp     a_types.DateTimeString         `json:"created_timestamp"`
	UpdatedTimestamp     a_types.DateTimeString         `json:"updated_timestamp"`
	CreatedBy            *go_types.Int64String          `json:"created_by"`
	UpdatedBy            *go_types.Int64String          `json:"updated_by"`
	Links                *Links                         `json:"links"`
	ID                   go_types.Int64String           `json:"id"`
	User                 go_types.Int64String           `json:"user"`
}

type GetListsConfig struct {
	Limit *uint
	Name  *string
}

func (service *Service) GetLists(getListsConfig *GetListsConfig) (*Lists, *errortools.Error) {
	params := url.Values{}

	lists := Lists{}
	offset := uint(0)
	limit := defaultLimit

	if getListsConfig != nil {
		if getListsConfig.Name != nil {
			params.Add("filters[name]", *getListsConfig.Name)
		}
		if getListsConfig.Limit != nil {
			limit = *getListsConfig.Limit
		}
	}
	params.Set("limit", fmt.Sprintf("%v", limit))

	for true {
		params.Set("offset", fmt.Sprintf("%v", offset))

		listsBatch := Lists{}

		requestConfig := go_http.RequestConfig{
			URL:           service.url(fmt.Sprintf("lists?%s", params.Encode())),
			ResponseModel: &listsBatch,
		}

		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		lists.Lists = append(lists.Lists, listsBatch.Lists...)

		if len(listsBatch.Lists) < int(limit) {
			break
		}
		offset += limit
	}

	return &lists, nil
}
