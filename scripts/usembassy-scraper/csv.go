package main

import (
	"encoding/csv"
	"log"
	"os"
)

// writeToCSV saves the contact data for all the posts in a tabular format.
func writeToCSV(posts []Post) {
	file, err := os.Create("post-contact.csv")

	if err != nil {
		log.Fatal("Failed to create CSV file", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	headers := []string{
		"name",
		"country",
		"website",
		"contact",
	}

	writer.Write(headers)

	for _, post := range posts {
		record := []string{
			post.name,
			post.country,
			post.website,
			post.contact,
		}

		writer.Write(record)
	}

	defer writer.Flush()
}
