package main

import "fmt"

type PostItem struct {
	contact string `dynamodbav:"contact"`
	country string `dynamodbav:"country"`
	name    string `dynamodbav:"name"`
	website string `dynamodbav:"website"`
}

const (
	AWS_REGION = "us-east-1"
	CSV_FILE   = "post-contact.csv"
	TABLE      = "travel-alerts-marek-posts"
)

func main() {
	tableExists := checkForTable()

	if tableExists {
		fmt.Printf("Processing the file %s...\n", CSV_FILE)

		processCSV(CSV_FILE)

		fmt.Print("\nUpload complete!")
	} else {
		fmt.Printf("The DynamoDB table %s was not found.", TABLE)
	}
}
