package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// connectToDB uses local configurations to open a connection to AWS DynamoDB.
func connectToDB() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		func(o *config.LoadOptions) error {
			o.Region = AWS_REGION
			return nil
		},
	)

	if err != nil {
		panic(err)
	}

	db := dynamodb.NewFromConfig(cfg)

	return db
}

// checkForTable verifies whether or not a DynamoDB table with a given name exists.
// The name of the table to look for is defined using the TABLE constant in main.go
func checkForTable() bool {
	db := connectToDB()

	tables, err := db.ListTables(context.TODO(), nil)

	if err != nil {
		panic(err)
	}

	found := false

	for _, val := range tables.TableNames {
		if val == TABLE {
			found = true
		}
	}

	return found
}

// putDynamoItem saves the data for a post item to DynamoDB. The name of the table
// where this item is saved is defined using the TABLE constant in main.go
func putDynamoItem(post PostItem) {
	db := connectToDB()

	item := map[string]types.AttributeValue{
		"contact": &types.AttributeValueMemberS{Value: post.contact},
		"country": &types.AttributeValueMemberS{Value: post.country},
		"name":    &types.AttributeValueMemberS{Value: post.name},
		"website": &types.AttributeValueMemberS{Value: post.website},
	}

	_, err := db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(TABLE),
		Item:      item,
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("\nSuccessfully added the post - %s", post.name)
}
