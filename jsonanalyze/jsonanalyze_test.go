package jsonanalyze

import (
	"testing"
)

func TestSingleAsterisk(t *testing.T) {
	filePath := "test_correct_json_files/single_asterisk.json"
	jsonData, err := LoadJsonFromFile(filePath)
	if err != nil {
		t.Fatalf("Test failed: err=%v", err)
	}
	want := false
	result, err := AnalyzeAWSIAMROLEJSON(jsonData)
	if result != want || err != nil {
		t.Fatalf("Test failed: result=%v, want=%v, err=%v", result, want, err)
	}
}

func TestNotSingleAsterisk(t *testing.T) {
	filePath := "test_correct_json_files/not_single_asterisk.json"
	jsonData, err := LoadJsonFromFile(filePath)
	if err != nil {
		t.Fatalf("Test failed: err=%v", err)
	}
	want := true
	result, err := AnalyzeAWSIAMROLEJSON(jsonData)
	if result != want || err != nil {
		t.Fatalf("Test failed: result=%v, want=%v, err=%v", result, want, err)
	}
}
