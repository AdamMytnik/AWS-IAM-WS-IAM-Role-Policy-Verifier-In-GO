package jsonanalyze

import (
	"encoding/json"
	"fmt"
	"os"
)

func AnalyzeAWSIAMROLEJSON(filePath string) (bool, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return false, err
	}

	defer func() {
		if err := jsonFile.Close(); err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}()

	var data map[string]interface{}
	if err := json.NewDecoder(jsonFile).Decode(&data); err != nil {
		return false, err
	}

	fmt.Println(data)
	return true, nil
}
