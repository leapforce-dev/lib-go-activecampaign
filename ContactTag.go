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

type ContactTags struct {
	ContactTags []ContactTag `json:"contactTags"`
	Meta        Meta         `json:"meta"`
}

type ContactTag struct {
	ContactID        go_types.Int64String           `json:"contact"`
	TagID            go_types.Int64String           `json:"tag"`
	CreatedDate      a_types.DateTimeTimezoneString `json:"cdate"`
	CreatedTimestamp a_types.DateTimeString         `json:"created_timestamp"`
	UpdatedTimestamp a_types.DateTimeString         `json:"updated_timestamp"`
	CreatedBy        *go_types.String               `json:"created_by"`
	UpdatedBy        *go_types.String               `json:"updated_by"`
	Links            *Links                         `json:"links"`
	ID               go_types.Int64String           `json:"id"`
}

type GetContactTagsConfig struct {
	Limit     *uint64
	Offset    *uint64
	ContactID *int64
}

func (service *Service) GetContactTags(getContactTagsConfig *GetContactTagsConfig) (*ContactTags, *errortools.Error) {
	params := url.Values{}

	contactTags := ContactTags{}
	rowCount := uint64(0)
	limit := defaultLimit

	path := "contactTags"

	if getContactTagsConfig != nil {
		if getContactTagsConfig.ContactID != nil {
			path = fmt.Sprintf("contacts/%v/contactTags", *getContactTagsConfig.ContactID)
		}
		if getContactTagsConfig.Limit != nil {
			limit = *getContactTagsConfig.Limit
		}
		if getContactTagsConfig.Offset != nil {
			service.nextOffsets.ContactTag = *getContactTagsConfig.Offset
		}
	}

	params.Add("limit", fmt.Sprintf("%v", limit))

	for {
		params.Set("offset", fmt.Sprintf("%v", service.nextOffsets.ContactTag))

		contactTagsBatch := ContactTags{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			URL:           service.url(fmt.Sprintf("%s?%s", path, params.Encode())),
			ResponseModel: &contactTags,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		contactTags.ContactTags = append(contactTags.ContactTags, contactTagsBatch.ContactTags...)

		if len(contactTagsBatch.ContactTags) < int(limit) {
			service.nextOffsets.ContactTag = 0
			break
		}

		service.nextOffsets.ContactTag += limit
		rowCount += limit

		if rowCount >= service.maxRowCount {
			return &contactTags, nil
		}
	}

	return &contactTags, nil
}
