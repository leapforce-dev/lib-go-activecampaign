package activecampaign

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

const (
	apiName            string = "ActiveCampaign"
	defaultMaxRowCount uint64 = ^uint64(0)
	defaultLimit       uint64 = 100
	timestampLayout    string = "2006-01-02 15:04:05"
)

type Service struct {
	host        string
	apiKey      string
	maxRowCount uint64
	httpService *go_http.Service
	nextOffsets struct {
		Automation        uint64
		Campaign          uint64
		Contact           uint64
		ContactAutomation uint64
		ContactField      uint64
		ContactFieldValue uint64
		ContactTag        uint64
		Deal              uint64
		DealField         uint64
		DealGroup         uint64
		DealStage         uint64
		List              uint64
		Message           uint64
		Segment           uint64
		Tag               uint64
	}
}

type ServiceConfig struct {
	Host        string
	ApiKey      string
	MaxRowCount *uint64
}

func NewService(serviceConfig *ServiceConfig) (*Service, *errortools.Error) {
	if serviceConfig.Host == "" {
		return nil, errortools.ErrorMessage("Host not provided")
	}

	if serviceConfig.ApiKey == "" {
		return nil, errortools.ErrorMessage("ApiKey not provided")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	maxRowCount := defaultMaxRowCount
	if serviceConfig.MaxRowCount != nil {
		maxRowCount = *serviceConfig.MaxRowCount
	}
	return &Service{
		host:        serviceConfig.Host,
		apiKey:      serviceConfig.ApiKey,
		maxRowCount: maxRowCount,
		httpService: httpService,
	}, nil
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	// add authentication header
	header := http.Header{}
	header.Set("Api-Token", service.apiKey)
	(*requestConfig).NonDefaultHeaders = &header

	// add error model
	errorResponse := ErrorResponse{}
	(*requestConfig).ErrorModel = &errorResponse

	request, response, e := service.httpService.HttpRequest(requestConfig)
	if len(errorResponse.Errors) > 0 {
		e.SetMessage(errorResponse.Errors[0].Title)
	}

	// activecampaign sometimes returns an error while the action has succesfully been performed
	if response != nil {
		if response.StatusCode >= 200 && response.StatusCode <= 299 {
			errortools.CaptureError(e)
			return request, response, nil
		}
	}

	return request, response, e
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("https://%s/api/3/%s", service.host, path)
}

func (service Service) ApiName() string {
	return apiName
}

func (service Service) ApiKey() string {
	return service.apiKey
}

func (service Service) ApiCallCount() int64 {
	return service.httpService.RequestCount()
}

func (service Service) ApiReset() {
	service.httpService.ResetRequestCount()
}
