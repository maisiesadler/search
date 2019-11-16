package dynamoindex

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (di *dynamoIndex) Add(key string, value string) {
	tableName := "Cache"

	item := make(map[string]*dynamodb.AttributeValue)
	item["Key"] = &dynamodb.AttributeValue{S: &key}
	item["Value"] = &dynamodb.AttributeValue{S: &value}

	putItem := &dynamodb.PutItemInput{TableName: &tableName, Item: item}
	_, err := di.svc.PutItem(putItem)
	if err != nil {
		fmt.Printf("got error %v", err)
	}
}
