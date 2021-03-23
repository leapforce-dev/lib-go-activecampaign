package facebook

import (
	"encoding/json"
	"time"
)

const (
	DateTimeFormat string = "2006-01-02 15:04:05"
)

type DateTimeString time.Time

func (d *DateTimeString) UnmarshalJSON(b []byte) error {
	var s string

	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	if s == "" || s == "0000-00-00 00:00:00" {
		d = nil
		return nil
	}

	_t, err := time.Parse(DateTimeFormat, s)
	if err != nil {
		return err
	}

	*d = DateTimeString(_t)
	return nil
}

func (d *DateTimeString) ValuePtr() *time.Time {
	if d == nil {
		return nil
	}

	_d := time.Time(*d)
	return &_d
}

func (d DateTimeString) Value() time.Time {
	return time.Time(d)
}
