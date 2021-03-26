package activecampaign

import (
	"fmt"
	"net/url"

	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type Segments struct {
	Segments []Segment `json:"segments"`
	Meta     Meta      `json:"meta"`
}

type Segment struct {
	Name             string                 `json:"name"`
	Logic            string                 `json:"logic"`
	Hidden           go_types.BoolString    `json:"hidden"`
	SeriesID         go_types.Int64String   `json:"seriesid"`
	CreatedTimestamp a_types.DateTimeString `json:"created_timestamp"`
	UpdatedTimestamp a_types.DateTimeString `json:"updated_timestamp"`
	CreatedBy        *go_types.String       `json:"created_by"`
	UpdatedBy        *go_types.String       `json:"updated_by"`
	Links            *Links                 `json:"links"`
	ID               go_types.Int64String   `json:"id"`
}

type GetSegmentsConfig struct {
	Limit *uint
}

func (service *Service) GetSegments(getSegmentsConfig *GetSegmentsConfig) (*Segments, *errortools.Error) {
	params := url.Values{}

	segments := Segments{}
	offset := uint(0)
	limit := defaultLimit

	if getSegmentsConfig != nil {
		if getSegmentsConfig.Limit != nil {
			limit = *getSegmentsConfig.Limit
		}
	}

	params.Add("limit", fmt.Sprintf("%v", limit))

	for true {
		params.Set("offset", fmt.Sprintf("%v", offset))

		segmentsBatch := Segments{}

		requestConfig := go_http.RequestConfig{
			URL:           service.url(fmt.Sprintf("segments?%s", params.Encode())),
			ResponseModel: &segmentsBatch,
		}

		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		segments.Segments = append(segments.Segments, segmentsBatch.Segments...)

		if len(segmentsBatch.Segments) < int(limit) {
			break
		}
		offset += limit
	}

	return &segments, nil
}
