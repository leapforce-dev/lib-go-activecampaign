package activecampaign

import (
	"encoding/json"
	"fmt"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Lists struct {
	Lists []List `json:"lists"`
	//Meta     FieldValuesMeta `json:"meta"`
}

type List struct {
	ID                   string          `json:"id"`
	StringID             string          `json:"stringid"`
	UserID               string          `json:"userid"`
	User                 string          `json:"user"`
	Name                 string          `json:"name"`
	CreateDate           string          `json:"cdate"`
	UseTracking          string          `json:"p_use_tracking"`
	UseAnalyticsRead     string          `json:"p_use_analytics_read"`
	UseAnalyticsLink     string          `json:"p_use_analytics_link"`
	UseTwitter           string          `json:"p_use_twitter"`
	UseFacebook          string          `json:"p_use_facebook"`
	EmbedImage           string          `json:"p_embed_image"`
	UseCaptcha           string          `json:"p_use_captcha"`
	SendLastBroadcast    string          `json:"send_last_broadcast"`
	Private              string          `json:"private"`
	AnalyticsDomains     json.RawMessage `json:"analytics_domains"`
	AnalyticsSource      string          `json:"analytics_source"`
	AnalyticsUA          string          `json:"analytics_ua"`
	TwitterToken         string          `json:"twitter_token"`
	TwitterTokenSecret   string          `json:"twitter_token_secret"`
	FacebookSession      json.RawMessage `json:"facebook_session"`
	CarbonCopy           json.RawMessage `json:"carboncopy"`
	SubscriptionNotify   json.RawMessage `json:"subscription_notify"`
	UnsubscriptionNotify json.RawMessage `json:"unsubscription_notify"`
	RequireName          string          `json:"require_name"`
	GetUnsubscribeReason string          `json:"get_unsubscribe_reason"`
	ToName               string          `json:"to_name"`
	OptInOptOut          string          `json:"optinoptout"`
	SenderName           string          `json:"sender_name"`
	SenderAddr1          string          `json:"sender_addr1"`
	SenderAddr2          string          `json:"sender_addr2"`
	SenderCity           string          `json:"sender_city"`
	SenderState          string          `json:"sender_state"`
	SenderZip            string          `json:"sender_zip"`
	SenderCountry        string          `json:"sender_country"`
	SenderPhone          string          `json:"sender_phone"`
	SenderURL            string          `json:"sender_url"`
	SenderReminder       string          `json:"sender_reminder"`
	FullAddress          string          `json:"fulladdress"`
	OptInMessageID       string          `json:"optinmessageid"`
	OptOutConf           string          `json:"optoutconf"`
	DeleteStamp          json.RawMessage `json:"deletestamp"`
	UDate                json.RawMessage `json:"udate"`
	CreatedTimestamp     string          `json:"created_timestamp"`
	UpdatedTimestamp     string          `json:"updated_timestamp"`
	CreatedBy            string          `json:"created_by"`
	UpdatedBy            string          `json:"updated_by"`
	Links                ListLinks       `json:"links"`
}

type ListLinks struct {
	Campaigns    string `json:"contactGoalLists"`
	ContactGoals string `json:"user"`
	ContactLists string `json:"addressLists"`
	Blocks       string `json:"blocks"`
	Goals        string `json:"goals"`
	SMS          string `json:"sms"`
	Sitemessages string `json:"sitemessages"`
}

type GetListsConfig struct {
	Limit *uint
	Name  *string
}

func (service *Service) GetLists(getListsConfig *GetListsConfig) (*Lists, *errortools.Error) {
	params := url.Values{}

	lists := Lists{}
	offset := uint(0)
	limit := uint(100)

	if getListsConfig != nil {
		if getListsConfig.Name != nil {
			params.Add("filters[name]", *getListsConfig.Name)
		}
		if getListsConfig.Limit != nil {
			limit = *getListsConfig.Limit
		}
	}
	params.Add("limit", fmt.Sprintf("%v", limit))

	for true {
		params.Set("offset", fmt.Sprintf("%v", offset))

		listsBatch := Lists{}

		requestConfig := go_http.RequestConfig{
			URL:           service.url("lists"),
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
