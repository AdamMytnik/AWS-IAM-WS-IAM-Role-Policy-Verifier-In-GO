package jsonanalyze

import (
	"bufio"
	"fmt"
	"os"
)

func LoadJsonFromFile(filePath string) (string, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := jsonFile.Close(); err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}()

	scanner := bufio.NewScanner(jsonFile)
	var jsonData string
	for scanner.Scan() {
		jsonData += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return jsonData, nil
}
