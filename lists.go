package activecampaign

import (
	"encoding/json"
	"fmt"
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

func (ac *ActiveCampaign) GetLists() (*Lists, error) {
	urlStr := fmt.Sprintf("%s/lists", ac.baseURL())

	lists := Lists{}

	err := ac.get(urlStr, &lists)
	if err != nil {
		return nil, err
	}

	return &lists, nil
}
