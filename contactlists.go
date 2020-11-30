package activecampaign

import (
	"bytes"
	"encoding/json"
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
)

type ContactLists struct {
	ContactLists []ContactList `json:"contactLists"`
	//Meta     FieldValuesMeta `json:"meta"`
}

type ContactList struct {
	ID                    string           `json:"id"`
	Automation            string           `json:"automation"`
	Contact               string           `json:"contact"`
	List                  string           `json:"list"`
	Form                  json.RawMessage  `json:"form"`
	SeriesID              string           `json:"seriesid"`
	SubscribeDate         string           `json:"sdate"`
	UnsubscribeDate       string           `json:"udate"`
	Status                string           `json:"status"`
	Responder             string           `json:"responder"`
	Sync                  string           `json:"sync"`
	UnsubscribeReason     string           `json:"unsubreason"`
	Campaign              string           `json:"campaign"`
	Message               string           `json:"message"`
	FirstName             string           `json:"first_name"`
	LastName              string           `json:"last_name"`
	IPSubcribe            string           `json:"ip4sub"`
	SourceID              string           `json:"sourceid"`
	AutosyncLog           json.RawMessage  `json:"autosyncLog"`
	IPLast                string           `json:"ip4_last"`
	IPUnsubscribe         string           `json:"ip4Unsub"`
	CreatedTimestamp      string           `json:"created_timestamp"`
	UpdatedTimestamp      string           `json:"updated_timestamp"`
	CreatedBy             string           `json:"created_by"`
	UpdatedBy             string           `json:"updated_by"`
	UnsubscribeAutomation string           `json:"unsubscribeAutomation"`
	Links                 ContactListLinks `json:"links"`
}

type ContactListLinks struct {
	Automation            string `json:"automation"`
	List                  string `json:"list"`
	Contact               string `json:"contact"`
	Form                  string `json:"form"`
	AutosyncLog           string `json:"autosyncLog"`
	Campaign              string `json:"campaign"`
	UnsubscribeAutomation string `json:"unsubscribeAutomation"`
	Message               string `json:"message"`
}

func (ac *ActiveCampaign) GetContactLists(contactID string) (*ContactLists, *errortools.Error) {
	urlStr := fmt.Sprintf("%s/contacts/%s/contactLists", ac.baseURL(), contactID)

	contactLists := ContactLists{}

	e := ac.get(urlStr, &contactLists)
	if e != nil {
		return nil, e
	}

	return &contactLists, nil
}

func (ac *ActiveCampaign) Subscribe(contactID int, listID int) *errortools.Error {
	return ac.setContactLists(contactID, listID, 1)
}

func (ac *ActiveCampaign) Unsubscribe(contactID int, listID int) *errortools.Error {
	return ac.setContactLists(contactID, listID, 2)
}

func (ac *ActiveCampaign) setContactLists(listID int, contactID int, status int) *errortools.Error {
	urlStr := fmt.Sprintf("%s/contactLists", ac.baseURL())

	type contactList struct {
		List    int `json:"list"`
		Contact int `json:"contact"`
		Status  int `json:"status"`
	}

	type data struct {
		ContactList contactList `json:"contactList"`
	}

	d := data{
		contactList{
			List:    listID,
			Contact: contactID,
			Status:  status,
		},
	}

	b, err := json.Marshal(d)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	buf := bytes.NewBuffer(b)

	e := ac.post(urlStr, buf, nil)
	if e != nil {
		return e
	}

	return nil
}
