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
	BaseTemplateID        *go_types.String                `json:"basetemplateid"`
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
	MailTransfer          go_types.Int64String            `json:"mail_transfer"`
	MailSend              go_types.Int64String            `json:"mail_send"`
	MailCleanup           go_types.Int64String            `json:"mail_cleanup"`
	MailerLogFile         go_types.Int64String            `json:"mailer_log_file"`
	TrackLinks            *go_types.String                `json:"tracklinks"`
	TrackLinksAnalytics   go_types.BoolString             `json:"tracklinksanalytics"`
	TrackReads            go_types.BoolString             `json:"trackreads"`
	TrackReadsAnalytics   go_types.BoolString             `json:"trackreadsanalytics"`
	AnalyticsCampaignName *go_types.String                `json:"analytics_campaign_name"`
	Tweet                 go_types.BoolString             `json:"tweet"`
	Facebook              go_types.BoolString             `json:"facebook"`
	Survey                *go_types.String                `json:"survey"`
	EmbedImages           go_types.BoolString             `json:"embed_images"`
	HTMLUnsubscibe        go_types.BoolString             `json:"htmlunsub"`
	TextUnsubscribe       go_types.BoolString             `json:"textunsub"`
	HTMLUnsubscibeData    *go_types.String                `json:"htmlunsubdata"`
	TextUnsubscribeData   *go_types.String                `json:"textunsubdata"`
	Recurring             *go_types.String                `json:"recurring"`
	WillRecur             go_types.BoolString             `json:"willrecur"`
	SplitType             *go_types.String                `json:"split_type"`
	SplitContent          go_types.BoolString             `json:"split_content"`
	SplitOffset           go_types.Int64String            `json:"split_offset"`
	SplitOffsetType       *go_types.String                `json:"split_offset_type"`
	SplitWinnerMessageID  go_types.Int64String            `json:"split_winner_messageid"`
	SplitWinnerAwaiting   go_types.BoolString             `json:"split_winner_awaiting"`
	ResponderOffset       go_types.Int64String            `json:"responder_offset"`
	ResponderType         *go_types.String                `json:"responder_type"`
	ResponderExisting     go_types.BoolString             `json:"responder_existing"`
	ReminderField         *go_types.String                `json:"reminder_field"`
	ReminderFormat        *go_types.String                `json:"reminder_format"`
	ReminderType          *go_types.String                `json:"reminder_type"`
	ReminderOffset        go_types.Int64String            `json:"reminder_offset"`
	ReminderOffsetType    *go_types.String                `json:"reminder_offset_type"`
	ReminderOffsetSign    *go_types.String                `json:"reminder_offset_sign"`
	ReminderLastCronRun   *a_types.DateTimeString         `json:"reminder_last_cron_run"`
	ActiveRSSInterval     *go_types.String                `json:"activerss_interval"`
	ActiveRSSURL          *go_types.String                `json:"activerss_url"`
	ActiveRSSItems        go_types.Int64String            `json:"activerss_items"`
	IP4                   *go_types.Int64String           `json:"ip4"`
	LastStep              *go_types.String                `json:"laststep"`
	ManageText            go_types.BoolString             `json:"managetext"`
	Schedule              go_types.BoolString             `json:"schedule"`
	ScheduleDate          *a_types.DateTimeTimezoneString `json:"scheduleddate"`
	WaitPreview           go_types.BoolString             `json:"waitpreview"`
	DeleteStamp           *a_types.DateTimeString         `json:"deletestamp"`
	ReplySys              go_types.BoolString             `json:"replysys"`
	Links                 *Links                          `json:"links"`
	ID                    go_types.Int64String            `json:"id"`
	AutomationID          *go_types.Int64String           `json:"automation"`
}

type GetCampaignsConfig struct {
	Limit           *uint64
	Offset          *uint64
	OrderBySendDate *OrderByDirection
}

func (service *Service) GetCampaigns(getCampaignsConfig *GetCampaignsConfig) (*Campaigns, *errortools.Error) {
	params := url.Values{}

	campaigns := Campaigns{}
	rowCount := uint64(0)
	limit := defaultLimit

	if getCampaignsConfig != nil {
		if getCampaignsConfig.OrderBySendDate != nil {
			params.Add("orders[sdate]", string(*getCampaignsConfig.OrderBySendDate))
		}
		if getCampaignsConfig.Limit != nil {
			limit = *getCampaignsConfig.Limit
		}
		if getCampaignsConfig.Offset != nil {
			service.nextOffsets.Campaign = *getCampaignsConfig.Offset
		}
	}

	params.Add("limit", fmt.Sprintf("%v", limit))

	for true {
		params.Set("offset", fmt.Sprintf("%v", service.nextOffsets.Campaign))

		campaignsBatch := Campaigns{}

		requestConfig := go_http.RequestConfig{
			URL:           service.url(fmt.Sprintf("campaigns?%s", params.Encode())),
			ResponseModel: &campaignsBatch,
		}
		fmt.Println(service.url(fmt.Sprintf("campaigns?%s", params.Encode())))

		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		campaigns.Campaigns = append(campaigns.Campaigns, campaignsBatch.Campaigns...)

		if len(campaignsBatch.Campaigns) < int(limit) {
			service.nextOffsets.Campaign = 0
			break
		}

		service.nextOffsets.Campaign += limit
		rowCount += limit

		if rowCount >= service.maxRowCount {
			return &campaigns, nil
		}
	}

	return &campaigns, nil
}
