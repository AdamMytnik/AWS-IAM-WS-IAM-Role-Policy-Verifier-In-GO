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

	//we can have NotResource instead of Resource so this could be an option too
	//its up to you to decide which one to use
	// var isSingleAsterisk bool
	// if policy.PolicyDocument.Statement[0].Resource.isEmpty() {
	// 	isSingleAsterisk = policy.PolicyDocument.Statement[0].Resource.containsAsterisk()
	// } else {
	// 	isSingleAsterisk = policy.PolicyDocument.Statement[0].NotResource.containsAsterisk()
	// }
	//Also we can have a situation that multiple Statements contain multiple Resources so we could
	//check if any statement contains single asterisk resource
	// var isSingleAsterisk bool

	// for _, statement := range policy.PolicyDocument.Statement {
	// 	if statement.Resource.isEmpty() {
	// 		isSingleAsterisk = statement.Resource.containsAsterisk()
	// 	} else {
	// 		isSingleAsterisk = statement.NotResource.containsAsterisk()
	// 	}
	// 	if isSingleAsterisk{return isSingleAsterisk, nil}
	// }

	//for the case of this task i will asume the most basic example with single statement
	//single resource and not using NotResource but those combinations are also implemented.
	isSingleAsterisk := policy.PolicyDocument.Statement[0].Resource.containsAsterisk()

	return isSingleAsterisk, nil
}
