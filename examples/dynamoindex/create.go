package dynamoindex

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Create table Cache
func createCacheIfDoesNotExist(svc *dynamodb.DynamoDB, tableName string) {

	describeTableInput := &dynamodb.DescribeTableInput{TableName: &tableName}
	output, err := svc.DescribeTable(describeTableInput)

	if err == nil {
		fmt.Printf("Table exists and has %v items\n", *output.Table.ItemCount)
	} else {
		input := &dynamodb.CreateTableInput{
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("Token"),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String("DocumentID"),
					AttributeType: aws.String("S"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("Token"),
					KeyType:       aws.String("HASH"),
				},
				{
					AttributeName: aws.String("DocumentID"),
					KeyType:       aws.String("RANGE"),
				},
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(10),
				WriteCapacityUnits: aws.Int64(10),
			},
			TableName: aws.String(tableName),
		}

		_, err := svc.CreateTable(input)
		if err != nil {
			fmt.Println("Got error calling CreateTable:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Created the table", tableName)
	}
}
