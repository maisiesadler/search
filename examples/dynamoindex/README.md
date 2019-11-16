# DynamoIndex

DynamoIndex is an implementation of Index backed by dynamodb.

Credentials from the shared credentials file ~/.aws/credentials
and region from the shared configuration file ~/.aws/config.

Creates a table called 'Cache' with the partition key 'Token' and sort key 'DocumentID'.

This can then be queried by token.

