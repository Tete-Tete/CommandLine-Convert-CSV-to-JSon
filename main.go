package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Check Command-line Arguments //
	args := os.Args
	if len(args) < 3 {
		log.Fatal("Please provide the input CSV file path and output JSON file path. Example: go run main.go input.csv output.json")
	}

	inputFile := args[1]  // Get the input CSV File path
	outputFile := args[2] // Get the output JSON file path

	//  Open The Input CSV File From Above //
	csvFile, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Sry, couldn't open the input file: %v, plz check again", err)
	}
	defer csvFile.Close() // Make sure close the function if we can't open the input CSV file

	// Read The CSV File From Above //
	csvReader := csv.NewReader(csvFile)
	records, err := csvReader.ReadAll() // Make sure get all data from the CSV File
	if err != nil {
		log.Fatalf("Sry, couldn't read the CSV file: %v, plz check again", err)
	}

	//Create The Output JSON File //
	jsonFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Sry, couldn't create the JSON file: %v, plz check again", err)
	}
	defer jsonFile.Close() // Make sure close the function if we can't create the output JSON file

	// Make A Loop Which Converts CSV To JSON//
	headers := records[0]
	for i, record := range records {
		if i == 0 {
			continue
		}

		// Create A map To Store Key-Value Pairs For JSON Conversion //
		jsonData := make(map[string]interface{})
		for j, value := range record {
			// Convert value to the correct type: try Float first, then Int, otherwise keep as String //
			if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
				jsonData[headers[j]] = floatValue
			} else if intValue, err := strconv.Atoi(value); err == nil {
				jsonData[headers[j]] = intValue
			} else {
				jsonData[headers[j]] = value
			}
		}

		// Convert to JSON //
		jsonBytes, err := json.Marshal(jsonData)
		if err != nil {
			log.Fatalf("Sry, couldn't convert the data to JSON: %v, plz check again", err)
		}

		// Put the converted JSON data to output file //
		if _, err = jsonFile.WriteString(fmt.Sprintf("%s\n", jsonBytes)); err != nil {

			log.Fatalf("Sry, couldn't put the converted the data to output JSON file: %v, plz check again", err)
		}
	}

	// Print success message //
	fmt.Println("CONVERT CSV TO JSON IS SUCCESSED!!!")
}
