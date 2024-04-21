package jsonanalyze

import (
	"encoding/json"
	"fmt"
)

type Resource struct {
	Value interface{} `json:"Resource"` //can be given either as list or str
}

func (r *Resource) UnmarshalJSON(data []byte) error {
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	switch v := value.(type) {
	case string:
		r.Value = v
	case []interface{}:
		var values []string
		for _, item := range v {
			if str, ok := item.(string); ok {
				values = append(values, str)
			} else {
				return fmt.Errorf("unexpected type in slice: %T", item)
			}
		}
		r.Value = values
	default:
		return fmt.Errorf("unexpected type: %T", value)
	}

	return nil
}

func (r *Resource) isEmpty() bool {
	switch v := r.Value.(type) {
	case string:
		return v == ""
	case []string:
		return len(v) == 0
	default:
		return true
	}
}

func (r *Resource) containsAsterisk() bool {
	switch v := r.Value.(type) {
	case string:
		return v != "*"
	case []string:
		for _, item := range v {
			if item == "*" {
				return false
			}
		}
		return true
	default:
		return true
	}
}
