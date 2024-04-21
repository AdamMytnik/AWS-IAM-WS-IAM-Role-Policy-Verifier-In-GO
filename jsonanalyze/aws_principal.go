package jsonanalyze

import (
	"encoding/json"
	"fmt"
)

type Principal struct {
	AWS           interface{} `json:"AWS"`
	CanonicalUser string      `json:"CanonicalUser"`
}

func (p *Principal) UnmarshalJSON(data []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if canonicalUser, ok := value["CanonicalUser"].(string); ok {
		p.CanonicalUser = canonicalUser
	}

	if awsValue, ok := value["AWS"]; ok {
		switch aws := awsValue.(type) {
		case string:
			p.AWS = aws
		case []interface{}:
			var values []string
			for _, item := range aws {
				if str, ok := item.(string); ok {
					values = append(values, str)
				} else {
					return fmt.Errorf("unexpected type in 'AWS' slice: %T", item)
				}
			}
			p.AWS = values
		default:
			return fmt.Errorf("unexpected type in 'AWS' field: %T", awsValue)
		}
	}
	return nil
}

func (p *Principal) isEmpty() bool {
	return p.isFieldEmpty(p.AWS) && p.isFieldEmpty(p.CanonicalUser)
}

func (p *Principal) isFieldEmpty(field interface{}) bool {
	switch v := field.(type) {
	case string:
		return v == ""
	case []string:
		return len(v) == 0
	default:
		return true
	}
}
