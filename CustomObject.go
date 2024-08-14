package activecampaign

import (
	"fmt"
	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	"net/http"
	"net/url"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

const CustomObjectRecordTimestampLayout string = "2006-01-02T15:04:05+00:00"

type CustomObjectRecordsResponse struct {
	Records []CustomObjectRecordRecord `json:"records"`
	Meta    Meta                       `json:"meta"`
}

type CustomObjectRecord struct {
	Record CustomObjectRecordRecord `json:"record"`
}

type CustomObjectRecordRecord struct {
	SchemaId         string                          `json:"schemaId"`
	Id               string                          `json:"id"`
	ExternalId       string                          `json:"externalId"`
	Fields           []CustomObjectRecordField       `json:"fields"`
	Relationships    CustomObjectRecordRelationships `json:"relationships"`
	CreatedTimestamp a_types.DateTimeString          `json:"createdTimestamp"`
	UpdatedTimestamp a_types.DateTimeString          `json:"updatedTimestamp"`
}

type CustomObjectRecordRelationships struct {
	PrimaryContact []go_types.Int64String `json:"primary-contact"`
}

type CustomObjectRecordField struct {
	Id    string      `json:"id"`
	Value interface{} `json:"value"`
}

type GetCustomObjectRecordsConfig struct {
	SchemaId     string
	Limit        *uint64
	CreatedAfter *time.Time
	UpdatedAfter *time.Time
}

func (service *Service) GetCustomObjectRecords(getCustomObjectRecordsConfig *GetCustomObjectRecordsConfig) (*CustomObjectRecordsResponse, bool, *errortools.Error) {
	params := url.Values{}

	customObjectRecords := CustomObjectRecordsResponse{}
	limit := defaultLimit

	if getCustomObjectRecordsConfig != nil {
		if getCustomObjectRecordsConfig.Limit != nil {
			limit = getLimit(*getCustomObjectRecordsConfig.Limit)
		}
		if getCustomObjectRecordsConfig.CreatedAfter != nil {
			params.Set("filters[createdTimestamp][gt]", (*getCustomObjectRecordsConfig.CreatedAfter).Format(timestampLayout2))
		}
		if getCustomObjectRecordsConfig.UpdatedAfter != nil {
			params.Set("filters[updatedTimestamp][gt]", (*getCustomObjectRecordsConfig.UpdatedAfter).Format(timestampLayout2))
		}
	}

	params.Set("limit", fmt.Sprintf("%v", limit))
	params.Set("orders[createdTimestamp]", "ASC")

	var maxCreatedTimestamp *time.Time = nil

	for {
		if maxCreatedTimestamp != nil {
			params.Set("filters[createdTimestamp][gt]", maxCreatedTimestamp.Format(timestampLayout2))
		}

		customObjectRecordsBatch := CustomObjectRecordsResponse{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("customObjects/records/%s?%s", getCustomObjectRecordsConfig.SchemaId, params.Encode())),
			ResponseModel: &customObjectRecordsBatch,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, false, e
		}

		rowCount := len(customObjectRecordsBatch.Records)

		if rowCount == 0 {
			break
		}

		customObjectRecords.Records = append(customObjectRecords.Records, customObjectRecordsBatch.Records...)

		maxCreatedTimestamp_ := customObjectRecordsBatch.Records[len(customObjectRecordsBatch.Records)-1].CreatedTimestamp.Value()
		maxCreatedTimestamp = &maxCreatedTimestamp_
	}

	return &customObjectRecords, false, nil
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
