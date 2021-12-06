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

type Messages struct {
	Messages []Message `json:"messages"`
	Meta     Meta      `json:"meta"`
}

type Message struct {
	UserID        go_types.Int64String           `json:"userid"`
	EdInstanceID  go_types.Int64String           `json:"ed_instanceid"`
	EdVersion     go_types.Int64String           `json:"ed_version"`
	CreatedDate   a_types.DateTimeTimezoneString `json:"cdate"`
	ModifiedDate  a_types.DateTimeTimezoneString `json:"mdate"`
	Name          string                         `json:"name"`
	FromName      string                         `json:"fromname"`
	FromEmail     string                         `json:"fromemail"`
	Reply2        string                         `json:"reply2"`
	Priority      go_types.Int64String           `json:"priority"`
	Charset       string                         `json:"charset"`
	Encoding      string                         `json:"encoding"`
	Format        string                         `json:"format"`
	Subject       string                         `json:"subject"`
	PreHeaderText string                         `json:"preheader_text"`
	Text          string                         `json:"text"`
	HTML          string                         `json:"html"`
	HTMLFetch     string                         `json:"htmlfetch"`
	TextFetch     string                         `json:"textfetch"`
	Hidden        go_types.BoolString            `json:"hidden"`
	PreviewMime   *go_types.String               `json:"preview_mime"`
	PreviewData   *go_types.String               `json:"preview_data"`
	Links         *Links                         `json:"links"`
	ID            go_types.Int64String           `json:"id"`
}

type GetMessagesConfig struct {
	Limit  *uint64
	Offset *uint64
}

func (service *Service) GetMessages(getMessagesConfig *GetMessagesConfig) (*Messages, *errortools.Error) {
	params := url.Values{}

	messages := Messages{}
	rowCount := uint64(0)
	limit := defaultLimit

	if getMessagesConfig != nil {
		if getMessagesConfig.Limit != nil {
			limit = *getMessagesConfig.Limit
		}
		if getMessagesConfig.Offset != nil {
			service.nextOffsets.Message = *getMessagesConfig.Offset
		}
	}

	params.Add("limit", fmt.Sprintf("%v", limit))

	for {
		params.Set("offset", fmt.Sprintf("%v", service.nextOffsets.Message))

		messagesBatch := Messages{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			URL:           service.url(fmt.Sprintf("messages?%s", params.Encode())),
			ResponseModel: &messagesBatch,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		messages.Messages = append(messages.Messages, messagesBatch.Messages...)

		if len(messagesBatch.Messages) < int(limit) {
			service.nextOffsets.Message = 0
			break
		}

		service.nextOffsets.Message += limit
		rowCount += limit

		if rowCount >= service.maxRowCount {
			return &messages, nil
		}
	}

	return &messages, nil
}
