package jsonanalyze

import (
	"testing"
)

func TestSingleAsterisk(t *testing.T) {
	filePath := "../aws_iam_role.json"
	want := true
	result, err := AnalyzeAWSIAMROLEJSON(filePath)
	if result != want || err != nil {
		t.Fatalf("Test failed: result=%v, err=%v, want=%v", result, err, want)
	}
}
