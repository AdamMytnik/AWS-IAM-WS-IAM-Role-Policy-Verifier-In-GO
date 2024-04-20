package jsonanalyze

import (
	"encoding/json"
)

func AnalyzeAWSIAMROLEJSON(jsonData string) (bool, error) {
	var policy IAMRolePolicy
	err := json.Unmarshal([]byte(jsonData), &policy)
	if err != nil {
		return false, err
	}

	if err := policy.Validate(); err != nil {
		return false, err
	}

	isSingleAsterisk := policy.PolicyDocument.Statement[0].Resource != "*"
	return isSingleAsterisk, nil
}
