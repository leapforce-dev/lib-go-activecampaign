package activecampaign

import (
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Tags struct {
	Tags []Tag `json:"tags"`
	//Meta     FieldValuesMeta `json:"meta"`
}

type Tag struct {
	TagType          string   `json:"tagType"`
	Tag              string   `json:"tag"`
	Description      string   `json:"description"`
	SubscriberCount  string   `json:"subscriber_count"`
	CreatedDate      string   `json:"cdate"`
	CreatedTimestamp string   `json:"created_timestamp"`
	UpdatedTimestamp string   `json:"updated_timestamp"`
	CreatedBy        string   `json:"created_by"`
	UpdatedBy        string   `json:"updated_by"`
	ID               string   `json:"id"`
	Links            TagLinks `json:"links"`
}

type TagLinks struct {
	ContactGoalTags string `json:"contactGoalTags"`
}

type GetTagsFilter struct {
	Email *string
}

func (service *Service) GetTags() (*Tags, *errortools.Error) {
	tags := Tags{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url("tags"),
		ResponseModel: &tags,
	}

	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &tags, nil
}
