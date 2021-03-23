package activecampaign

import (
	"fmt"
	"net/http"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

const (
	limit           int    = 20
	TimestampFormat string = "2006-01-02 15:04:05"
	XDateFormat     string = "2006-01-02T15:04:05"
)

type CustomField struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type Service struct {
	host        string
	apiKey      string
	httpService *go_http.Service
}

type ServiceConfig struct {
	Host   string
	APIKey string
}

func NewService(serviceConfig *ServiceConfig) (*Service, *errortools.Error) {
	if serviceConfig.Host == "" {
		return nil, errortools.ErrorMessage("Host not provided")
	}

	if serviceConfig.APIKey == "" {
		return nil, errortools.ErrorMessage("APIKey not provided")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	return &Service{
		host:        serviceConfig.Host,
		apiKey:      serviceConfig.APIKey,
		httpService: httpService,
	}, nil
}

func (service *Service) httpRequest(httpMethod string, requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	// add authentication header
	header := http.Header{}
	header.Set("Api-Token", service.apiKey)
	(*requestConfig).NonDefaultHeaders = &header

	// add error model
	errorResponse := ErrorResponse{}
	(*requestConfig).ErrorModel = &errorResponse

	request, response, e := service.httpService.HTTPRequest(httpMethod, requestConfig)
	if len(errorResponse.Errors) > 0 {
		e.SetMessage(errorResponse.Errors[0].Title)
	}

	// activecampaign sometimes returns an error while the action has succesfully been performed
	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		errortools.CaptureError(e)
		return request, response, nil
	}

	return request, response, e
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("https://%s/api/3/%s", service.host, path)
}

func (service *Service) get(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodGet, requestConfig)
}

func (service *Service) post(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPost, requestConfig)
}

func (service *Service) put(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPut, requestConfig)
}

func (service *Service) delete(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodDelete, requestConfig)
}

func ParseXDate(xdate string) (*time.Time, *errortools.Error) {

	t, err := time.Parse(XDateFormat, xdate[:len(XDateFormat)])
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	return &t, nil
}
