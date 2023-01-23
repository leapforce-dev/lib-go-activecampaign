package activecampaign

import (
	a_types "github.com/leapforce-libraries/go_activecampaign/types"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
	"net/http"
)

type UsersResponse struct {
	Users []User `json:"users"`
	Meta  Meta   `json:"meta"`
}

type User struct {
	Username                    *string                         `json:"username,omitempty"`
	FirstName                   *string                         `json:"firstName,omitempty"`
	LastName                    *string                         `json:"lastName,omitempty"`
	Email                       *string                         `json:"email,omitempty"`
	Password                    *string                         `json:"password,omitempty"`
	Phone                       *string                         `json:"phone,omitempty"`
	Signature                   *string                         `json:"signature,omitempty"`
	Language                    *string                         `json:"lang,omitempty"`
	LocalZoneId                 *string                         `json:"localZoneid,omitempty"`
	PasswordUpdatedUtcTimestamp *a_types.DateTimeString         `json:"password_updated_utc_timestamp,omitempty"`
	CreatedDate                 *a_types.DateTimeTimezoneString `json:"cdate,omitempty"`
	UpdatedDate                 *a_types.DateTimeTimezoneString `json:"udate,omitempty"`
	MfaEnabled                  *go_types.BoolString            `json:"mfaEnabled,omitempty"`
	Links                       *Links                          `json:"links,omitempty"`
	Id                          *go_types.Int64String           `json:"id,omitempty"`
	Group                       *int64                          `json:"group,omitempty"`
}

func (service *Service) GetUsers() (*[]User, *errortools.Error) {

	usersResponse := UsersResponse{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url("users"),
		ResponseModel: &usersResponse,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &usersResponse.Users, nil
}

func (service *Service) CreateUser(user *User) (*User, *errortools.Error) {
	if user == nil {
		return nil, nil
	}

	userResponse := struct {
		User User `json:"user"`
	}{}

	requestConfig := go_http.RequestConfig{
		Method: http.MethodPost,
		Url:    service.url("users"),
		BodyModel: struct {
			User User `json:"user"`
		}{*user},
		ResponseModel: &userResponse,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &userResponse.User, nil
}
