package activecampaign

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	types "github.com/leapforce-libraries/go_types"
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

func NewActiveCampaign(accountName string, apiKey string) (*ActiveCampaign, error) {
	ac := ActiveCampaign{}
	ac.accountName = accountName
	ac.apiKey = apiKey
	ac.Client = http.Client{}

	return &ac, nil
}

func (ac *ActiveCampaign) baseURL() string {
	return fmt.Sprintf(apiURL, ac.accountName)
}

func (ac *ActiveCampaign) limit() int {
	return limit
}

func (ac *ActiveCampaign) httpRequest(httpMethod string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		return nil, err
	}

	// Add authorization token to header
	req.Header.Set("Api-Token", ac.apiKey)

	// Send out the HTTP request
	response, err := ac.Client.Do(req)

	// Check HTTP StatusCode
	if response.StatusCode < 200 || response.StatusCode > 299 {
		message := fmt.Sprintf("Server returned statuscode %v", response.StatusCode)
		return response, &types.ErrorString{message}
	}
	if err != nil {
		return response, err
	}

	return response, nil
}

func (ac *ActiveCampaign) get(url string, model interface{}) error {
	res, err := ac.httpRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &model)
	if err != nil {
		return err
	}

	return nil
}

func (ac *ActiveCampaign) post(url string, buf *bytes.Buffer, model interface{}) error {
	res, err := ac.httpRequest(http.MethodPost, url, buf)
	if err != nil {
		if res != nil {
			defer res.Body.Close()

			b, errRead := ioutil.ReadAll(res.Body)
			if errRead != nil {
				return err
			}

			var errors struct {
				Errors []ActiveCampaignError `json:"errors"`
			}

			errUnmarshal := json.Unmarshal(b, &errors)
			if errUnmarshal != nil {
				return err
			}

			return &types.ErrorString{fmt.Sprintf("Error: %v, title: %s", err, errors.Errors[0].Title)}
		} else {
			return err
		}
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &model)
	if err != nil {
		return err
	}

	return nil
}
