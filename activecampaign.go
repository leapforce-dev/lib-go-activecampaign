package activecampaign

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	utilities "github.com/leapforce-libraries/go_utilities"
)

const (
	apiName string = "ActiveCampaign"
	apiURL  string = "https://%s.api-us1.com/api/3"
	limit   int    = 20
)

type ActiveCampaign struct {
	accountName string
	apiKey      string
	Client      http.Client
}

type ActiveCampaignErrorSource struct {
	Pointer string `json:"pointer"`
}

type ActiveCampaignError struct {
	Title  string                    `json:"title"`
	Detail string                    `json:"detail"`
	Code   string                    `json:"code"`
	Error  string                    `json:"error"`
	Source ActiveCampaignErrorSource `json:"error"`
}

type CustomField struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

func NewActiveCampaign(accountName string, apiKey string) *ActiveCampaign {
	ac := ActiveCampaign{}
	ac.accountName = accountName
	ac.apiKey = apiKey
	ac.Client = http.Client{}

	return &ac
}

func (ac *ActiveCampaign) baseURL() string {
	return fmt.Sprintf(apiURL, ac.accountName)
}

func (ac *ActiveCampaign) limit() int {
	return limit
}

func (ac *ActiveCampaign) httpRequest(httpMethod string, url string, body io.Reader) (*http.Request, *http.Response, *errortools.Error) {
	e := new(errortools.Error)

	request, err := http.NewRequest(httpMethod, url, body)
	e.SetRequest(request)
	if err != nil {
		e.SetMessage(err)
		return request, nil, e
	}

	// Add authorization token to header
	request.Header.Set("Api-Token", ac.apiKey)
	request.Header.Set("Accept", "application/json")
	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	response, ee := utilities.DoWithRetry(&ac.Client, request, 10, 5)
	e.SetResponse(response)

	if ee != nil {
		e.SetMessage(ee)
		return request, response, e
	}

	return request, response, nil
}

func (ac *ActiveCampaign) get(url string, model interface{}) *errortools.Error {
	request, response, e := ac.httpRequest(http.MethodGet, url, nil)
	if e != nil {
		return e
	}

	e = new(errortools.Error)
	e.SetRequest(request)
	e.SetResponse(response)

	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		e.SetMessage(err)
		return e
	}

	err = json.Unmarshal(b, &model)
	if err != nil {
		e.SetMessage(err)
		return e
	}

	return nil
}

func (ac *ActiveCampaign) post(url string, buf *bytes.Buffer, model interface{}) *errortools.Error {
	request, response, e := ac.httpRequest(http.MethodPost, url, buf)

	if e != nil {
		if response != nil {

			e = new(errortools.Error)
			e.SetRequest(request)
			e.SetResponse(response)

			defer response.Body.Close()

			b, err := ioutil.ReadAll(response.Body)
			if err != nil {
				e.SetMessage(err)
				return e
			}

			var errors struct {
				Errors []ActiveCampaignError `json:"errors"`
			}

			err = json.Unmarshal(b, &errors)
			if err != nil {
				e.SetMessage(err)
				return e
			}

			e.SetMessage(fmt.Sprintf("Error: %v, title: %s", err, errors.Errors[0].Title))
			return e
		} else {
			return e
		}
	}

	e = new(errortools.Error)
	e.SetRequest(request)
	e.SetResponse(response)

	if model != nil {
		defer response.Body.Close()

		b, err := ioutil.ReadAll(response.Body)
		if err != nil {
			e.SetMessage(err)
			return e
		}

		err = json.Unmarshal(b, &model)
		if err != nil {
			e.SetMessage(err)
			return e
		}
	}

	return nil
}

func (ac *ActiveCampaign) put(url string, buf *bytes.Buffer, model interface{}) *errortools.Error {
	request, response, e := ac.httpRequest(http.MethodPut, url, buf)

	if e != nil {
		if response != nil {

			e = new(errortools.Error)
			e.SetRequest(request)
			e.SetResponse(response)

			defer response.Body.Close()

			b, err := ioutil.ReadAll(response.Body)
			if err != nil {
				e.SetMessage(err)
				return e
			}

			var errors struct {
				Errors []ActiveCampaignError `json:"errors"`
			}

			err = json.Unmarshal(b, &errors)
			if err != nil {
				e.SetMessage(err)
				return e
			}

			e.SetMessage(fmt.Sprintf("Error: %v, title: %s", err, errors.Errors[0].Title))
			return e
		} else {
			return e
		}
	}

	e = new(errortools.Error)
	e.SetRequest(request)
	e.SetResponse(response)

	if model != nil {
		defer response.Body.Close()

		b, err := ioutil.ReadAll(response.Body)
		if err != nil {
			e.SetMessage(err)
			return e
		}

		err = json.Unmarshal(b, &model)
		if err != nil {
			e.SetMessage(err)
			return e
		}
	}

	return nil
}

func (ac *ActiveCampaign) delete(url string) *errortools.Error {
	_, _, e := ac.httpRequest(http.MethodDelete, url, nil)
	if e != nil {
		return e
	}

	return nil
}
