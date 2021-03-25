package activecampaign

import (
	"fmt"
	"net/url"

	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type Campaigns struct {
	Campaigns []Campaign `json:"campaigns"`
	Meta      Meta       `json:"meta"`
}

type Campaign struct {
	Type                  string                          `json:"type"`
	UserID                go_types.Int64String            `json:"userid"`
	SegmentID             go_types.Int64String            `json:"segmentid"`
	BounceID              go_types.Int64String            `json:"bounceid"`
	RealcID               go_types.Int64String            `json:"realcid"`
	SendID                go_types.Int64String            `json:"sendid"`
	ThreadID              go_types.Int64String            `json:"threadid"`
	SeriesID              go_types.Int64String            `json:"seriesid"`
	FormID                go_types.Int64String            `json:"formid"`
	BaseTemplateID        *string                         `json:"basetemplateid"`
	BaseMessageID         go_types.Int64String            `json:"basemessageid"`
	AddressID             go_types.Int64String            `json:"addressid"`
	Source                string                          `json:"source"`
	Name                  string                          `json:"name"`
	CreatedDate           a_types.DateTimeTimezoneString  `json:"cdate"`
	ModifiedDate          a_types.DateTimeTimezoneString  `json:"mdate"`
	SendDate              *a_types.DateTimeTimezoneString `json:"sdate"`
	LDate                 *a_types.DateTimeTimezoneString `json:"ldate"`
	SendAmount            go_types.Int64String            `json:"send_amount"`
	TotalAmount           go_types.Int64String            `json:"total_amt"`
	Opens                 go_types.Int64String            `json:"opens"`
	UniqueOpens           go_types.Int64String            `json:"uniqueopens"`
	LinkClicks            go_types.Int64String            `json:"linkclicks"`
	UniqueLinkClicks      go_types.Int64String            `json:"uniquelinkclicks"`
	SubscriberClicks      go_types.Int64String            `json:"subscriberclicks"`
	Forwards              go_types.Int64String            `json:"forwards"`
	UniqueForwards        go_types.Int64String            `json:"uniqueforwards"`
	HardBounces           go_types.Int64String            `json:"hardbounces"`
	SoftBounces           go_types.Int64String            `json:"softbounces"`
	Unsubscribes          go_types.Int64String            `json:"unsubscribes"`
	UnsubscribeReasons    go_types.Int64String            `json:"unsubreasons"`
	Updates               go_types.Int64String            `json:"updates"`
	SocialShares          go_types.Int64String            `json:"socialshares"`
	Replies               go_types.Int64String            `json:"replies"`
	UniqueReplies         go_types.Int64String            `json:"uniquereplies"`
	Status                go_types.Int64String            `json:"status"`
	Public                go_types.BoolString             `json:"public"`
	MailTransfer          go_types.BoolString             `json:"mail_transfer"`
	MailSend              go_types.BoolString             `json:"mail_send"`
	MailCleanup           go_types.BoolString             `json:"mail_cleanup"`
	MailerLogFile         go_types.BoolString             `json:"mailer_log_file"`
	TrackLinks            *string                         `json:"tracklinks"`
	TrackLinksAnalytics   go_types.BoolString             `json:"tracklinksanalytics"`
	TrackReads            go_types.BoolString             `json:"trackreads"`
	TrackReadsAnalytics   go_types.BoolString             `json:"trackreadsanalytics"`
	AnalyticsCampaignName *string                         `json:"analytics_campaign_name"`
	Tweet                 go_types.BoolString             `json:"tweet"`
	Facebook              go_types.BoolString             `json:"facebook"`
	Survey                *string                         `json:"survey"`
	EmbedImages           go_types.BoolString             `json:"embed_images"`
	HTMLUnsubscibe        go_types.BoolString             `json:"htmlunsub"`
	TextUnsubscribe       go_types.BoolString             `json:"textunsub"`
	HTMLUnsubscibeData    *string                         `json:"htmlunsubdata"`
	TextUnsubscribeData   *string                         `json:"textunsubdata"`
	Recurring             *string                         `json:"recurring"`
	WillRecur             go_types.BoolString             `json:"willrecur"`
	SplitType             *string                         `json:"split_type"`
	SplitContent          go_types.BoolString             `json:"split_content"`
	SplitOffset           go_types.Int64String            `json:"split_offset"`
	SplitOffsetType       *string                         `json:"split_offset_type"`
	SplitWinnerMessageID  go_types.Int64String            `json:"split_winner_messageid"`
	SplitWinnerAwaiting   go_types.BoolString             `json:"split_winner_awaiting"`
	ResponderOffset       go_types.Int64String            `json:"responder_offset"`
	ResponderType         *string                         `json:"responder_type"`
	ResponderExisting     go_types.BoolString             `json:"responder_existing"`
	ReminderField         *string                         `json:"reminder_field"`
	ReminderFormat        *string                         `json:"reminder_format"`
	ReminderType          *string                         `json:"reminder_type"`
	ReminderOffset        go_types.Int64String            `json:"reminder_offset"`
	ReminderOffsetType    *string                         `json:"reminder_offset_type"`
	ReminderOffsetSign    *string                         `json:"reminder_offset_sign"`
	ReminderLastCronRun   *a_types.DateTimeString         `json:"reminder_last_cron_run"`
	ActiveRSSInterval     *string                         `json:"activerss_interval"`
	ActiveRSSURL          *string                         `json:"activerss_url"`
	ActiveRSSItems        go_types.Int64String            `json:"activerss_items"`
	IP4                   *go_types.Int64String           `json:"ip4"`
	LastStep              *string                         `json:"laststep"`
	ManageText            go_types.BoolString             `json:"managetext"`
	Schedule              go_types.BoolString             `json:"schedule"`
	ScheduleDate          *a_types.DateTimeString         `json:"scheduleddate"`
	WaitPreview           go_types.BoolString             `json:"waitpreview"`
	DeleteStamp           *a_types.DateTimeString         `json:"deletestamp"`
	ReplySys              go_types.BoolString             `json:"replysys"`
	Links                 *Links                          `json:"links"`
	ID                    go_types.Int64String            `json:"id"`
	User                  *go_types.Int64String           `json:"user"`
	Automation            *go_types.Int64String           `json:"automation"`
}

type OrderByDirection string

const (
	OrderByDirectionAscending  OrderByDirection = "ASC"
	OrderByDirectionDescending OrderByDirection = "DESC"
)

type GetCampaignsConfig struct {
	Limit           *uint
	OrderBySendDate *OrderByDirection
}

func (service *Service) GetCampaigns(getCampaignsConfig *GetCampaignsConfig) (*Campaigns, *errortools.Error) {
	params := url.Values{}

	campaigns := Campaigns{}
	offset := uint(0)
	limit := defaultLimit

	if getCampaignsConfig != nil {
		if getCampaignsConfig.OrderBySendDate != nil {
			params.Add("orders[sdate]", string(*getCampaignsConfig.OrderBySendDate))
		}
		if getCampaignsConfig.Limit != nil {
			limit = *getCampaignsConfig.Limit
		}
	}

	params.Add("limit", fmt.Sprintf("%v", limit))

	for true {
		params.Set("offset", fmt.Sprintf("%v", offset))

		campaignsBatch := Campaigns{}

		requestConfig := go_http.RequestConfig{
			URL:           service.url(fmt.Sprintf("campaigns?%s", params.Encode())),
			ResponseModel: &campaignsBatch,
		}

		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		campaigns.Campaigns = append(campaigns.Campaigns, campaignsBatch.Campaigns...)

		if len(campaignsBatch.Campaigns) < int(limit) {
			break
		}
		offset += limit
	}

	return &campaigns, nil
}
