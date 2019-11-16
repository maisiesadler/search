package dynamoindex

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/maisiesadler/search/index"
)

func (di *dynamoIndex) Find(word string) (bool, []*index.Result) {
	output, err := di.queryCache(word)

	if err != nil {
		fmt.Printf("got error %v", err)
		return false, []*index.Result{}
	}

	if *output.Count == 0 {
		return false, []*index.Result{}
	}

	results := parseResults(output)

	return true, toArray(results)
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

func parseResults(output *dynamodb.QueryOutput) (m *map[string]*map[string]int) {
	results := make(map[string]*map[string]int)
	getOrAdd := func(token *string) *map[string]int {
		if result, ok := results[*token]; ok {
			return result
		}

		m := make(map[string]int)
		return &m
	}

	for _, item := range output.Items {
		token := item["Token"]
		docID := item["DocumentID"]
		results[*token.S] = getOrAdd(token.S)
		(*results[*token.S])[*docID.S] = 1 // todo: does this ever need to be > 1
	}

	return &results
}

func toArray(m *map[string]*map[string]int) []*index.Result {
	s := []*index.Result{}
	for k, v := range *m {
		result := &index.Result{Matches: *v, Word: k}
		s = append(s, result)
	}

	return s
}
