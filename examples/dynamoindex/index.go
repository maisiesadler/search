package dynamoindex

import (
	"github.com/maisiesadler/search/index"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Create the credentials from AssumeRoleProvider to assume the role
// referenced by the "myRoleARN" ARN.
// creds := stscreds.NewCredentials(sess, "Dynamo")

type dynamoIndex struct {
	svc       *dynamodb.DynamoDB
	tableName *string
}

// CreateDynamoIndex connects to dynamo and returns an instance of dynamoIndex
func CreateDynamoIndex() (index.Index, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	tableName := "Cache"
	createCacheIfDoesNotExist(svc, tableName)

	return &dynamoIndex{svc, &tableName}, nil
}
