package jsonanalyze

import "encoding/json"

type ResourceValue struct {
	Resource  string   `json:"Resource"` //can be given either as list or str
	Resources []string `json:"Resources"`
}

func (rv ResourceValue) containsAsterisk() bool {
	if rv.IsSlice() {
		for _, resource := range rv.Resources {
			if resource == "*" {
				return false
			}
		}
	} else {
		return rv.Resource != "*"
	}
	return true
}

func (rv ResourceValue) isEmpty() bool {
	return rv.IsString() || rv.IsSlice()
}

func (rv ResourceValue) IsString() bool {
	return rv.Resource != ""
}

func (rv ResourceValue) IsSlice() bool {
	return len(rv.Resources) > 0
}

func (rv *ResourceValue) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		// If the data starts with a double quote, it's a string value.
		var value string
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		rv.Resource = value
	} else {
		if err := json.Unmarshal(data, &rv.Resources); err != nil {
			return err
		}
	}
	return nil
}
