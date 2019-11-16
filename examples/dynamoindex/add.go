package dynamoindex

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (di *dynamoIndex) Add(docID string, tokens <-chan string) {
	for token := range tokens {
		di.AddOne(docID, token)
	}
}

func (di *dynamoIndex) AddOne(docID string, token string) {
	tableName := "Cache"

	item := make(map[string]*dynamodb.AttributeValue)
	item["Token"] = &dynamodb.AttributeValue{S: &token}
	item["DocumentID"] = &dynamodb.AttributeValue{S: &docID}

	putItem := &dynamodb.PutItemInput{TableName: &tableName, Item: item}
	_, err := di.svc.PutItem(putItem)
	if err != nil {
		fmt.Printf("got error %v", err)
	}
}
