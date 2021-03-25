package activecampaign

import (
	"fmt"
	"net/url"

	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

type Fields struct {
	FieldOptions   []FieldOption   `json:"fieldOptions"`
	FieldRelations []FieldRelation `json:"fieldRels"`
	Fields         []Field         `json:"fields"`
	Meta           Meta            `json:"meta"`
}

type FieldOption struct {
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

type FieldRelation struct {
	FieldID      go_types.Int64String           `json:"field"`
	RelationID   go_types.Int64String           `json:"relid"`
	DisplayOrder go_types.Int64String           `json:"dorder"`
	CreatedDate  a_types.DateTimeTimezoneString `json:"cdate"`
	Links        *Links                         `json:"links"`
	ID           go_types.Int64String           `json:"id"`
}

type Field struct {
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

type FieldUpdate struct {
	Title        string `json:"title,omitempty"`
	Description  string `json:"descript,omitempty"`
	Type         string `json:"type,omitempty"`
	IsRequired   string `json:"isrequired,omitempty"`
	Perstag      string `json:"perstag,omitempty"`
	DefaultValue string `json:"defval,omitempty"`
	Visible      string `json:"visible,omitempty"`
	Service      string `json:"service,omitempty"`
	Ordernum     string `json:"ordernum,omitempty"`
}

type FieldLink struct {
	Options   string `json:"options"`
	Relations string `json:"Relations"`
}

type GetFieldsConfig struct {
	Limit *uint
}

func (service *Service) GetFields(getFieldsConfig *GetFieldsConfig) (*Fields, *errortools.Error) {
	params := url.Values{}

	fields := Fields{}
	offset := uint(0)
	limit := defaultLimit

	if getFieldsConfig != nil {
		if getFieldsConfig.Limit != nil {
			limit = *getFieldsConfig.Limit
		}
	}

	params.Add("limit", fmt.Sprintf("%v", limit))

	for true {
		params.Set("offset", fmt.Sprintf("%v", offset))

		fieldsBatch := Fields{}

		requestConfig := go_http.RequestConfig{
			URL:           service.url(fmt.Sprintf("fields?%s", params.Encode())),
			ResponseModel: &fieldsBatch,
		}

		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		fields.Fields = append(fields.Fields, fieldsBatch.Fields...)

		if len(fieldsBatch.Fields) < int(limit) {
			break
		}
		offset += limit
	}

	return &fields, nil
}

func (service *Service) CreateField(fieldUpdate *FieldUpdate) (*Field, *errortools.Error) {
	if fieldUpdate == nil {
		return nil, nil
	}

	d := struct {
		Field FieldUpdate `json:"field"`
	}{
		Field: *fieldUpdate,
	}

	var fieldUpdated struct {
		Field Field `json:"field"`
	}

	requestConfig := go_http.RequestConfig{
		URL:           service.url("fields"),
		BodyModel:     d,
		ResponseModel: &fieldUpdated,
	}

	_, _, e := service.post(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &fieldUpdated.Field, nil
}

func (service *Service) UpdateField(fieldID string, fieldUpdate *FieldUpdate) (*Field, *errortools.Error) {
	if fieldUpdate == nil {
		return nil, nil
	}

	d := struct {
		Field FieldUpdate `json:"field"`
	}{
		Field: *fieldUpdate,
	}

	var fieldUpdated struct {
		Field Field `json:"field"`
	}

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("fields/%s", fieldID)),
		BodyModel:     d,
		ResponseModel: &fieldUpdated,
	}

	_, _, e := service.put(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &fieldUpdated.Field, nil
}

func (service *Service) DeleteField(fieldID string) *errortools.Error {
	requestConfig := go_http.RequestConfig{
		URL: service.url(fmt.Sprintf("fields/%s", fieldID)),
	}

	_, _, e := service.delete(&requestConfig)
	if e != nil {
		return e
	}

	return nil
}

type FieldRelationUpdate struct {
	Field int32 `json:"field"`
	RelID int32 `json:"relid"`
}

func (service *Service) CreateFieldRelation(fieldID int32, listID int32) (*FieldRelation, *errortools.Error) {
	d := struct {
		FieldRelationUpdate `json:"fieldRel"`
	}{
		FieldRelationUpdate{fieldID, listID},
	}

	fieldRelation := FieldRelation{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url("fieldRels"),
		BodyModel:     d,
		ResponseModel: &fieldRelation,
	}

	_, _, e := service.post(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &fieldRelation, nil
}
