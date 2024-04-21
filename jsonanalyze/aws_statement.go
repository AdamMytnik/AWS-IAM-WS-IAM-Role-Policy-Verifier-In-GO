package jsonanalyze

import (
	"fmt"
)

/*
Some JSON policy elements are mutually exclusive.
This means that you cannot create a policy that uses both.
For example, you cannot use both Action and NotAction in the same policy statement.
Other pairs that are mutually exclusive include Principal/NotPrincipal and Resource/NotResource.
*/

type Statement struct {
	Sid          string   `json:"Sid"`    //Optional but should be unique in Policy
	Effect       string   `json:"Effect"` //Required
	Action       []string `json:"Action"` //Required
	NotAction    []string `json:"NotAction"`
	Resource     string   `json:"Resource"`    //(Required in only some circumstances)
	NotResource  string   `json:"NotResource"` // you must provide at least one of them
	Principal    string   `json:"Principal"`   // (Required in only some circumstances) TODO
	NotPrincipal string   `json:"NotPrincipal"`
	Condition    string   `json:"Condition"` //Optional
}

func (s *Statement) Validate() error {
	var validationErrors []error

	if s.Effect == "" {
		validationErrors = append(validationErrors, fmt.Errorf("missing Effect; Set it to Allow or Deny"))
	} else if s.Effect != "Allow" && s.Effect != "Deny" {
		validationErrors = append(validationErrors, fmt.Errorf("invalid Effect value; Should be Allow/Deny"))
	}

	if (len(s.Action) == 0 && len(s.NotAction) == 0) || (len(s.Action) != 0 && len(s.NotAction) != 0) {
		validationErrors = append(validationErrors, fmt.Errorf("either Action or NotAction must be specified in the statement"))
	}

	if s.Principal != "" && s.NotPrincipal != "" {
		validationErrors = append(validationErrors, fmt.Errorf("either Principal or NotPrincipal can be specified in the statement"))
	}

	if (s.Resource == "" && s.NotResource == "") || (s.Resource != "" && s.NotResource != "") {
		validationErrors = append(validationErrors, fmt.Errorf("either Resource or NotResource can be specified in the statement"))
	}

	if len(validationErrors) > 0 {
		errorMsg := formatErrors("validation errors in Statement:", validationErrors)
		return fmt.Errorf(errorMsg)
	}

	return nil
}
