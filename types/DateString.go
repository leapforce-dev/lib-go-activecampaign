package facebook

import (
	"encoding/json"
	"time"

	"cloud.google.com/go/civil"
)

const (
	DateFormat string = "2006-01-02"
)

type DateString civil.Date

func (d *DateString) UnmarshalJSON(b []byte) error {
	var s string

	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	if s == "" || s == "0000-00-00" {
		d = nil
		return nil
	}

	_t, err := time.Parse(DateFormat, s)
	if err != nil {
		return err
	}

	*d = DateString(civil.DateOf(_t))
	return nil
}

func (d *DateString) ValuePtr() *civil.Date {
	if d == nil {
		return nil
	}

	_d := civil.Date(*d)
	return &_d
}

func (d DateString) Value() civil.Date {
	return civil.Date(d)
}
