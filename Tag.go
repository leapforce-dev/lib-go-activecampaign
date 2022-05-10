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

type Tags struct {
	Tags []Tag `json:"tags"`
	Meta Meta  `json:"meta"`
}

type Tag struct {
	TagType          string                         `json:"tagType"`
	Tag              string                         `json:"tag"`
	Description      *go_types.String               `json:"description"`
	SubscriberCount  go_types.Int64String           `json:"subscriber_count"`
	CreatedDate      a_types.DateTimeTimezoneString `json:"cdate"`
	CreatedTimestamp a_types.DateTimeString         `json:"created_timestamp"`
	UpdatedTimestamp a_types.DateTimeString         `json:"updated_timestamp"`
	CreatedBy        *go_types.Int64String          `json:"created_by"`
	UpdatedBy        *go_types.Int64String          `json:"updated_by"`
	Id               go_types.Int64String           `json:"id"`
	Links            *Links                         `json:"links"`
}

type GetTagsConfig struct {
	Limit  *uint64
	Offset *uint64
	Search *string
}

func (service *Service) GetTags(getTagsConfig *GetTagsConfig) (*Tags, *errortools.Error) {
	params := url.Values{}

	tags := Tags{}
	rowCount := uint64(0)
	limit := defaultLimit

	if getTagsConfig != nil {
		if getTagsConfig.Search != nil {
			params.Add("search", *getTagsConfig.Search)
		}
		if getTagsConfig.Limit != nil {
			limit = *getTagsConfig.Limit
		}
		if getTagsConfig.Offset != nil {
			service.nextOffsets.Tag = *getTagsConfig.Offset
		}
	}

	params.Add("limit", fmt.Sprintf("%v", limit))

	for {
		params.Set("offset", fmt.Sprintf("%v", service.nextOffsets.Tag))

		tagsBatch := Tags{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("tags?%s", params.Encode())),
			ResponseModel: &tagsBatch,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		tags.Tags = append(tags.Tags, tagsBatch.Tags...)

		if len(tagsBatch.Tags) < int(limit) {
			service.nextOffsets.Tag = 0
			break
		}

		service.nextOffsets.Tag += limit
		rowCount += limit

		if rowCount >= service.maxRowCount {
			return &tags, nil
		}
	}

	return &tags, nil
}
