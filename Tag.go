package activecampaign

import (
	"fmt"
	"net/url"

	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type Tags struct {
	Tags []Tag `json:"tags"`
	Meta Meta  `json:"meta"`
}

type Tag struct {
	TagType          string                         `json:"tagType"`
	Tag              string                         `json:"tag"`
	Description      *string                        `json:"description"`
	SubscriberCount  go_types.Int64String           `json:"subscriber_count"`
	CreatedDate      a_types.DateTimeTimezoneString `json:"cdate"`
	CreatedTimestamp a_types.DateTimeString         `json:"created_timestamp"`
	UpdatedTimestamp a_types.DateTimeString         `json:"updated_timestamp"`
	CreatedBy        *go_types.Int64String          `json:"created_by"`
	UpdatedBy        *go_types.Int64String          `json:"updated_by"`
	ID               go_types.Int64String           `json:"id"`
	Links            TagLinks                       `json:"links"`
}

type TagLinks struct {
	ContactGoalTags string `json:"contactGoalTags"`
}

type GetTagsConfig struct {
	Search *string
}

func (service *Service) GetTags(getTagsConfig *GetTagsConfig) (*Tags, *errortools.Error) {
	params := url.Values{}

	tags := Tags{}

	if getTagsConfig != nil {
		if getTagsConfig.Search != nil {
			params.Add("search", *getTagsConfig.Search)
		}
	}

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("tags?%s", params.Encode())),
		ResponseModel: &tags,
	}

	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &tags, nil
}
