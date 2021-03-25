package activecampaign

import (
	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	go_types "github.com/leapforce-libraries/go_types"
)

type FieldValuesDeal struct {
	FieldValues []FieldValueDeal `json:"fieldValues"`
	Meta        Meta             `json:"meta"`
}

type FieldValueDeal struct {
	DealID               go_types.Int64String    `json:"deal_id"`
	FieldID              go_types.Int64String    `json:"custom_field_id"`
	TextValue            *go_types.String        `json:"custom_field_text_value"`
	TextBlob             *go_types.String        `json:"custom_field_text_blob"`
	DateValue            *a_types.DateTimeString `json:"custom_field_date_value"`
	NumberValue          *go_types.Float64String `json:"custom_field_number_value"`
	CurrencyValue        *go_types.Float64String `json:"custom_field_currency_value"`
	CurrencyType         *go_types.String        `json:"custom_field_currency_type"`
	CreatedTimestamp     a_types.DateTimeString  `json:"created_timestamp"`
	UpdatedTimestamp     a_types.DateTimeString  `json:"updated_timestamp"`
	CreatedBy            *go_types.String        `json:"created_by"`
	UpdatedBy            *go_types.String        `json:"updated_by"`
	Links                Links                   `json:"links"`
	ID                   go_types.Int64String    `json:"id"`
	DealCustomFieldMetum go_types.Int64String    `json:"dealCustomFieldMetum"`
}
