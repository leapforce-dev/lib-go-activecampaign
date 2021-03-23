package facebook

import (
	"encoding/json"
	"time"
)

const (
	DateTimeTimezoneFormat string = time.RFC3339
)

type DateTimeTimezoneString time.Time

func (d *DateTimeTimezoneString) UnmarshalJSON(b []byte) error {
	var s string

	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	if s == "" || s == "0000-00-00T00:00:00Z00:00" {
		d = nil
		return nil
	}

	_t, err := time.Parse(DateTimeTimezoneFormat, s)
	if err != nil {
		return err
	}

	*d = DateTimeTimezoneString(_t)
	return nil
}

func (d *DateTimeTimezoneString) ValuePtr() *time.Time {
	if d == nil {
		return nil
	}

	_d := time.Time(*d)
	return &_d
}

func (d DateTimeTimezoneString) Value() time.Time {
	return time.Time(d)
}
