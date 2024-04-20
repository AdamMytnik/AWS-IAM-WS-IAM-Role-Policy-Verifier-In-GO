package jsonanalyze

import (
	"strings"
)

func formatErrors(header string, errors []error) string {
	var errorMsg strings.Builder
	errorMsg.WriteString(header + "\n")
	for _, err := range errors {
		errorMsg.WriteString("\t" + strings.ReplaceAll(err.Error(), "\t", "\t\t") + "\n")
	}
	return errorMsg.String()
}
