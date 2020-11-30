package activecampaign

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
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

func (ac *ActiveCampaign) GetTags() (*Tags, *errortools.Error) {
	urlStr := fmt.Sprintf("%s/tags", ac.baseURL())

	tags := Tags{}

	e := ac.get(urlStr, &tags)
	if e != nil {
		return nil, e
	}

	return &tags, nil
}
