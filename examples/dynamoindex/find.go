package dynamoindex

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/maisiesadler/search/index"
)

func (di *dynamoIndex) Find(key string) (bool, *index.DictionaryResult) {
	output, err := di.queryCache(key)

	if err != nil {
		fmt.Printf("got error %v", err)
		return false, nil
	}

	if *output.Count == 0 {
		return false, nil
	}

	return true, parseResults(key, output)
}

func (di *dynamoIndex) queryCache(word string) (*dynamodb.QueryOutput, error) {
	expressionAttributeNames := make(map[string]*string)
	expressionAttributeNames["#token"] = aws.String("Token")
	expressionAttributeValues := make(map[string]*dynamodb.AttributeValue)
	expressionAttributeValues[":v_Token"] = &dynamodb.AttributeValue{S: &word}
	queryInput := &dynamodb.QueryInput{
		TableName:                 di.tableName,
		ExpressionAttributeValues: expressionAttributeValues,
		ExpressionAttributeNames:  expressionAttributeNames,
		KeyConditionExpression:    aws.String("#token = :v_Token"),
	}
	return di.svc.Query(queryInput)
}

func parseResults(key string, output *dynamodb.QueryOutput) *index.DictionaryResult {
	result := &index.DictionaryResult{
		Key:             key,
		ValueOccurences: make(map[string]int),
	}

	for _, item := range output.Items {
		token := item["Token"]
		if *token.S != key {
			fmt.Printf("Unexpected token retreived. Expected %v, actual %v.", key, token)
		}
		docID := item["DocumentID"]
		result.ValueOccurences[*docID.S] = 1 // todo: does this ever need to be > 1
	}

	return result
}
