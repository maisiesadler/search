package dynamoindex

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/maisiesadler/search/index"
)

func (di *dynamoDictionary) Find(key string) (bool, *index.DictionaryResult) {
	output, err := di.queryCache(key)

	if err != nil {
		fmt.Printf("got error %v", err)
		return false, nil
	}

	if *output.Count == 0 {
		return false, nil
	}

	return true, di.parseResults(key, output)
}

func (di *dynamoDictionary) queryCache(word string) (*dynamodb.QueryOutput, error) {
	expressionAttributeNames := make(map[string]*string)
	expressionAttributeNames["#dictkey"] = di.keyName
	expressionAttributeValues := make(map[string]*dynamodb.AttributeValue)
	expressionAttributeValues[":v_DictKey"] = &dynamodb.AttributeValue{S: &word}
	queryInput := &dynamodb.QueryInput{
		TableName:                 di.tableName,
		ExpressionAttributeValues: expressionAttributeValues,
		ExpressionAttributeNames:  expressionAttributeNames,
		KeyConditionExpression:    aws.String("#dictkey = :v_DictKey"),
	}
	return di.svc.Query(queryInput)
}

func (di *dynamoDictionary) parseResults(key string, output *dynamodb.QueryOutput) *index.DictionaryResult {
	result := &index.DictionaryResult{
		Key:             key,
		ValueOccurences: make(map[string]int),
	}

	for _, item := range output.Items {
		token := item[*di.keyName]
		if *token.S != key {
			fmt.Printf("Unexpected token retreived. Expected %v, actual %v.", key, token)
		}
		docID := item[*di.valueName]
		result.ValueOccurences[*docID.S] = 1 // todo: does this ever need to be > 1
	}

	return result
}
