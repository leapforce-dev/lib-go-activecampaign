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
	Links            *Links                         `json:"links"`
}

type GetTagsConfig struct {
	Limit  *uint
	Search *string
}

func (service *Service) GetTags(getTagsConfig *GetTagsConfig) (*Tags, *errortools.Error) {
	params := url.Values{}

	tags := Tags{}
	offset := uint(0)
	limit := defaultLimit

	if getTagsConfig != nil {
		if getTagsConfig.Search != nil {
			params.Add("search", *getTagsConfig.Search)
		}
		if getTagsConfig.Limit != nil {
			limit = *getTagsConfig.Limit
		}
	}

	params.Add("limit", fmt.Sprintf("%v", limit))

	for true {
		params.Set("offset", fmt.Sprintf("%v", offset))

		tagsBatch := Tags{}

		requestConfig := go_http.RequestConfig{
			URL:           service.url(fmt.Sprintf("tags?%s", params.Encode())),
			ResponseModel: &tagsBatch,
		}

		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		tags.Tags = append(tags.Tags, tagsBatch.Tags...)

		if len(tagsBatch.Tags) < int(limit) {
			break
		}
		offset += limit
	}

	return &tags, nil
}
