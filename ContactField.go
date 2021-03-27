package activecampaign

import (
	"fmt"
	"net/url"

	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type ContactFields struct {
	FieldOptions   []ContactFieldOption   `json:"fieldOptions"`
	FieldRelations []ContactFieldRelation `json:"fieldRels"`
	ContactFields  []ContactField         `json:"fields"`
	Meta           Meta                   `json:"meta"`
}

type ContactFieldOption struct {
	FieldID     go_types.Int64String           `json:"field"`
	OrderID     go_types.Int64String           `json:"orderid"`
	Value       string                         `json:"value"`
	Label       string                         `json:"label"`
	IsDefault   go_types.BoolString            `json:"isdefault"`
	CreatedDate a_types.DateTimeTimezoneString `json:"cdate"`
	UpdatedDate a_types.DateTimeTimezoneString `json:"udate"`
	Links       *Links                         `json:"links"`
	ID          go_types.Int64String           `json:"id"`
}

type ContactFieldRelation struct {
	FieldID      go_types.Int64String           `json:"field"`
	RelationID   go_types.Int64String           `json:"relid"`
	DisplayOrder go_types.Int64String           `json:"dorder"`
	CreatedDate  a_types.DateTimeTimezoneString `json:"cdate"`
	Links        *Links                         `json:"links"`
	ID           go_types.Int64String           `json:"id"`
}

type ContactField struct {
	Title              string                         `json:"title"`
	Description        string                         `json:"descript"`
	Type               string                         `json:"type"`
	IsRequired         go_types.BoolString            `json:"isrequired"`
	PersonalizationTag string                         `json:"perstag"`
	DefaultValue       string                         `json:"defval"`
	ShowInList         go_types.BoolString            `json:"show_in_list"`
	Rows               go_types.Int64String           `json:"rows"`
	Columns            go_types.Int64String           `json:"cols"`
	Visible            go_types.BoolString            `json:"visible"`
	Service            string                         `json:"service"`
	OrderNumber        go_types.Int64String           `json:"ordernum"`
	CreatedDate        a_types.DateTimeTimezoneString `json:"cdate"`
	UpdatedDate        a_types.DateTimeTimezoneString `json:"udate"`
	Options            []go_types.Int64String         `json:"options"`
	Relations          []go_types.Int64String         `json:"relations"`
	Links              *Links                         `json:"links"`
	ID                 go_types.Int64String           `json:"id"`
}

type GetContactFieldsConfig struct {
	Limit *uint
}

func (service *Service) GetContactFields(getContactFieldsConfig *GetContactFieldsConfig) (*ContactFields, *errortools.Error) {
	params := url.Values{}

	contactFields := ContactFields{}
	offset := uint(0)
	limit := defaultLimit

	if getContactFieldsConfig != nil {
		if getContactFieldsConfig.Limit != nil {
			limit = *getContactFieldsConfig.Limit
		}
	}

	params.Add("limit", fmt.Sprintf("%v", limit))

	for true {
		params.Set("offset", fmt.Sprintf("%v", offset))

		contactFieldsBatch := ContactFields{}

		requestConfig := go_http.RequestConfig{
			URL:           service.url(fmt.Sprintf("fields?%s", params.Encode())),
			ResponseModel: &contactFieldsBatch,
		}

		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		contactFields.ContactFields = append(contactFields.ContactFields, contactFieldsBatch.ContactFields...)

		if len(contactFieldsBatch.ContactFields) < int(limit) {
			break
		}
		offset += limit
	}

	return &contactFields, nil
}
