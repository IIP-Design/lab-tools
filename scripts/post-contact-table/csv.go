package main

import (
	"encoding/csv"
	"os"
)

// processCSV opens the designated file and loops through each row therein
// uploading each entry as a PostItem to DynamoDB
func processCSV(csvFile string) {
	// Open the CSV file
	file, err := os.Open(csvFile)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)

	data, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	// Loop through rows uploading them to DynamoDB.
	for _, row := range data {
		var post PostItem

		post.name = row[0]
		post.country = row[1]
		post.website = row[2]
		post.contact = row[3]

		putDynamoItem(post)
	}
}
