package activecampaign

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_types "github.com/leapforce-libraries/go_types"
)

// Int64String is an extended version of go_types.Int64String.
// UnmarshalJSON also tries to unmarshal to int64 directly because ActiveCampaign sometimes returns an integer as int64, e.g. Account.Owner
type Int64String go_types.Int64String
type Int64Strings go_types.Int64Strings

func (i *Int64String) UnmarshalJSON(b []byte) error {
	var returnError = func() error {
		errortools.CaptureError(fmt.Sprintf("Cannot parse '%s' to Int64String", string(b)))
		return nil
	}

	var s string

	err := json.Unmarshal(b, &s)
	if err != nil {
		// try to marshal to int64
		var ii int64
		err := json.Unmarshal(b, &ii)
		if err != nil {
			return returnError()
		} else {
			*i = Int64String(ii)
			return nil
		}
	}

	s = strings.Trim(s, " ")

	if s == "" {
		i = nil
		return nil
	}

	_i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return returnError()
	}

	*i = Int64String(_i)
	return nil
}

func (i Int64String) MarshalJSON() ([]byte, error) {
	return i.MarshalJSON()
}

func (i *Int64String) ValuePtr() *int64 {
	return i.ValuePtr()
}

func (i Int64String) Value() int64 {
	return i.Value()
}

func (is *Int64Strings) ToInt64() []int64 {
	return is.ToInt64()
}
