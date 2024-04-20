package jsonanalyze

import (
	"fmt"
	"regexp"
)

type IAMRolePolicy struct {
	PolicyName     string         `json:"PolicyName"`
	PolicyDocument PolicyDocument `json:"PolicyDocument"`
}

func (p *IAMRolePolicy) Validate() error {
	var validationErrors []error

	if p.PolicyName == "" {
		validationErrors = append(validationErrors, fmt.Errorf("missing PolicyName"))
	} else {
		matched, _ := regexp.MatchString(`^[\w+=,.@-]+$`, p.PolicyName)
		if !matched {
			validationErrors = append(validationErrors, fmt.Errorf("PolicyName must match the pattern: ^[\\w+=,.@-]+$"))
		}
		if len(p.PolicyName) < 1 || len(p.PolicyName) > 128 {
			validationErrors = append(validationErrors, fmt.Errorf("PolicyName length must be between 1 and 128 characters"))
		}
	}

	if err := p.PolicyDocument.Validate(); err != nil {
		validationErrors = append(validationErrors, fmt.Errorf("%s", err.Error()))
	}

	if len(validationErrors) > 0 {
		errorMsg := formatErrors("vvalidation errors in IAMRolePolicy:", validationErrors)
		return fmt.Errorf(errorMsg)
	}

	return nil
}
