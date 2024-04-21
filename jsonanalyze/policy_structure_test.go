package jsonanalyze

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func readErrorFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text()), nil
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", nil
}

func TestCorrectPolicyStructure(t *testing.T) {
	testDir := "test_correct_json_files"
	err := filepath.Walk(testDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			t.Run(path, func(t *testing.T) {
				jsonData, err := LoadJsonFromFile(path)
				if err != nil {
					t.Fatalf("Test failed: err=%v", err)
				}
				_, err = AnalyzeAWSIAMROLEJSON(jsonData)
				if err != nil {
					t.Fatalf("Test failed: err=%v", err)
				}
			})
		}
		return nil
	})
	if err != nil {
		t.Fatalf("Error walking directory: %v", err)
	}
}

func TestInCorrectPolicyStructure(t *testing.T) {
	testDir := "test_incorrect_json_files"
	err := filepath.Walk(testDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && !strings.HasSuffix(info.Name(), ".error") {
			t.Run(path, func(t *testing.T) {
				errorFilePath := strings.TrimSuffix(path, ".json") + ".error"
				expectedError, err := readErrorFile(errorFilePath)
				if err != nil {
					t.Fatalf("Test failed: error reading expected errors file: %v", err)
				}
				jsonData, err := LoadJsonFromFile(path)
				if err != nil {
					t.Fatalf("Test failed: err=%v", err)
				}
				_, err = AnalyzeAWSIAMROLEJSON(jsonData)
				if err == nil {
					t.Fatalf("Test failed: error should be thrown! expected error: %v got %v", expectedError, err)
				}
			})
		}
		return nil
	})
	if err != nil {
		t.Fatalf("Error walking directory: %v", err)
	}
}
