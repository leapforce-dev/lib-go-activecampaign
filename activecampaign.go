package activecampaign

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	types "github.com/Leapforce-nl/go_types"
)

const (
	apiName string = "ActiveCampaign"
	apiURL  string = "https://%s.api-us1.com/api/3"
)

type ActiveCampaign struct {
	accountName string
	apiKey      string
	Client      http.Client
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

func (ac *ActiveCampaign) httpRequest(httpMethod string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(httpMethod, url, nil)
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
		return nil, &types.ErrorString{message}
	}
	if err != nil {
		return nil, err
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
