package jsonanalyze

import (
	"fmt"
)

type PolicyDocument struct {
	Version   string      `json:"Version"`
	Id        string      `json:"Id"` //Optional, recommended to be UUID
	Statement []Statement `json:"Statement"`
}

func (p *PolicyDocument) Validate() error {
	var validationErrors []error
	sidCount := make(map[string]int) //We can't have 2 identical Sid in same policy!

	if p.Version == "" {
		err := fmt.Errorf("missing Version in PolicyDocument: %s", p.Version)
		validationErrors = append(validationErrors, err)
	}

	if len(p.Statement) == 0 {
		err := fmt.Errorf("missing Statement in PolicyDocument")
		validationErrors = append(validationErrors, err)
	}
	for _, statement := range p.Statement {
		if err := statement.Validate(); err != nil {
			validationErrors = append(validationErrors, err)
		}
		if statement.Sid != "" {
			sidCount[statement.Sid]++
			if sidCount[statement.Sid] > 1 {
				validationErrors = append(validationErrors, fmt.Errorf("statement error: Sid '%s' appeared in %d statements; Sid must be unique within the policy", statement.Sid, sidCount[statement.Sid]))
			}
		}
	}

	if len(validationErrors) > 0 {
		errorMsg := formatErrors("validation errors in PolicyDocument:", validationErrors)
		return fmt.Errorf(errorMsg)
	}

	return nil
}
