package activecampaign

import "encoding/json"

type Links map[string]string

func (l *Links) UnmarshalJSON(b []byte) error {
	m := make(map[string]string)

	err := json.Unmarshal(b, &m)
	if err != nil {
		l = nil
		return nil
	}

	*l = m
	return nil
}

func (l *Links) ValuePtr() *map[string]string {
	if l == nil {
		return nil
	}

	_l := map[string]string(*l)
	return &_l
}

func (l Links) Value() map[string]string {
	return map[string]string(l)
}
