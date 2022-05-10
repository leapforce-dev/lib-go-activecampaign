package activecampaign

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

const CustomObjectRecordTimestampLayout string = "2006-01-02T15:04:05+00:00"

type CustomObjectRecord struct {
	Record CustomObjectRecordRecord `json:"record"`
}

type CustomObjectRecordRecord struct {
	SchemaId      string                          `json:"schemaId"`
	Id            *string                         `json:"id"`
	ExternalId    *string                         `json:"externalId"`
	Fields        []CustomObjectRecordField       `json:"fields"`
	Relationships CustomObjectRecordRelationships `json:"relationships"`
}

type CustomObjectRecordRelationships struct {
	PrimaryContact []go_types.Int64String `json:"primary-contact"`
}

type CustomObjectRecordField struct {
	Id    string      `json:"id"`
	Value interface{} `json:"value"`
}

func (service *Service) CreateCustomObjectRecord(r *CustomObjectRecord) (*CustomObjectRecord, *errortools.Error) {
	if r == nil {
		return nil, errortools.ErrorMessage("CustomObjectRecord is nil")
	}

	customObjectRecord := CustomObjectRecord{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url(fmt.Sprintf("customObjects/records/%s", r.Record.SchemaId)),
		BodyModel:     r,
		ResponseModel: &customObjectRecord,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &customObjectRecord, nil
}

func (service *Service) DeleteCustomObjectRecordByExternalId(schemaId string, externalId string) *errortools.Error {
	if externalId == "" {
		return errortools.ErrorMessage("externalId is empty")
	}

	requestConfig := go_http.RequestConfig{
		Method: http.MethodDelete,
		Url:    service.url(fmt.Sprintf("customObjects/records/%s/external/%s", schemaId, externalId)),
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return e
	}

	return nil
}
