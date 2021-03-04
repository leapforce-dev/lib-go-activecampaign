package activecampaign

import (
	"fmt"
	"strconv"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Fields struct {
	FieldOptions   interface{}     `json:"fieldOptions"`
	FieldRelations []FieldRelation `json:"fieldRels"`
	Fields         []Field         `json:"fields"`
	Meta           FieldsMeta      `json:"meta"`
}

type FieldsMeta struct {
	Total string `json:"total"`
}

type FieldRelation struct {
	Field      string `json:"field"`
	RelationID string `json:"relid"`
	DOrder     string `json:"dorder"`
	CreateDate string `json:"cdate"`
	//Links      interface{} `json:"links"`
	ID string `json:"id"`
}

type Field struct {
	Title        string `json:"title"`
	Description  string `json:"descript"`
	Type         string `json:"type"`
	IsRequired   string `json:"isrequired"`
	Perstag      string `json:"perstag"`
	DefaultValue string `json:"defval"`
	Visible      string `json:"visible"`
	Service      string `json:"service"`
	Ordernum     string `json:"ordernum"`
	CreateDate   string `json:"cdate"`
	UpdateDate   string `json:"udate"`
	//Options      interface{}    `json:"options"`
	Relations []string  `json:"relations"`
	Links     FieldLink `json:"links"`
	ID        string    `json:"id"`
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

func (service *Service) GetCustomFields() (*Fields, *errortools.Error) {
	rowCount := 0

	fields := Fields{}

	for true {
		fields_ := Fields{}

		requestConfig := go_http.RequestConfig{
			URL:           service.url(fmt.Sprintf("fields?limit=%v&offset=%v", limit, rowCount)),
			ResponseModel: &fields_,
		}

		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		fields.Fields = append(fields.Fields, fields_.Fields...)
		rowCount += len(fields_.Fields)

		total, err := strconv.Atoi(fields_.Meta.Total)
		if err != nil {
			return nil, errortools.ErrorMessage(err)
		}

		if rowCount >= total {
			break
		}
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
