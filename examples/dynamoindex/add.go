package dynamoindex

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (di *dynamoDictionary) Add(key string, value string) {

	item := make(map[string]*dynamodb.AttributeValue)
	item[*di.keyName] = &dynamodb.AttributeValue{S: &key}
	item[*di.valueName] = &dynamodb.AttributeValue{S: &value}

	putItem := &dynamodb.PutItemInput{TableName: di.tableName, Item: item}
	_, err := di.svc.PutItem(putItem)
	if err != nil {
		fmt.Printf("got error %v", err)
	}
}
