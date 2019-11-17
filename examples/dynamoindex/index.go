package dynamoindex

import (
	"github.com/maisiesadler/search/index"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Create the credentials from AssumeRoleProvider to assume the role
// referenced by the "myRoleARN" ARN.
// creds := stscreds.NewCredentials(sess, "Dynamo")

type dynamoDictionary struct {
	svc       *dynamodb.DynamoDB
	tableName *string
	keyName   *string
	valueName *string
}

// CreateDynamoDictionary connects to dynamo and returns an instance of dynamoIndex
func CreateDynamoDictionary() (index.Dictionary, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	tableName := "Cache"
	keyName := "DictKey"
	valueName := "DictValue"
	createCacheIfDoesNotExist(svc, tableName, keyName, valueName)

	return &dynamoDictionary{svc, &tableName, &keyName, &valueName}, nil
}
