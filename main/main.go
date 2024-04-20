package main

import (
	"bufio"
	"fmt"
	"os"

	"mymodule.com/jsonanalyze"
)

func loadJsonFromFile(filePath string) (string, error) {
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

func main() {
	filePath := "../aws_iam_role.json"
	jsonData, err := loadJsonFromFile(filePath)
	if err != nil {
		fmt.Printf("Error loading JSON from file: %v\n", err)
		return
	}
	result, err := jsonanalyze.AnalyzeAWSIAMROLEJSON(jsonData)
	if err != nil {
		fmt.Printf("Error while analyzing AWS::IAM::Role Policy Error: %v\n", err)
		return
	}
	fmt.Println("Analyzing json result:", result)
}
