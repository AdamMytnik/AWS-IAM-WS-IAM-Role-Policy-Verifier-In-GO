package main

import (
	"fmt"

	"mymodule.com/jsonanalyze"
)

func main() {
	filePath := "../aws_iam_role.json"
	jsonData, err := jsonanalyze.LoadJsonFromFile(filePath)
	if err != nil {
		fmt.Printf("Error loading JSON from file: %v\n", err)
		return
	}
	fmt.Printf("Loaded JSON:\n")
	if err := jsonanalyze.PrintJSON(jsonData); err != nil {
		fmt.Printf("Error printing JSON: %v\n", err)
		return
	}
	result, err := jsonanalyze.AnalyzeAWSIAMROLEJSON(jsonData)
	if err != nil {
		fmt.Printf("Error while analyzing AWS::IAM::Role Policy Error: %v\n", err)
		return
	}
	fmt.Println("Analyzing json result:", result)
}
