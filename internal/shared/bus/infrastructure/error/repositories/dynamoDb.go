package repositories

import (
	error2 "TestProject/internal/error"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

const (
	PKFormat = "Pk#%s"
	SKFormat = "Sk#%s"
)

type DynamoDbErrorRepository struct {
	client *dynamodb.Client
}

func NewDynamoDbErrorRepository(client *dynamodb.Client) *DynamoDbErrorRepository {
	return &DynamoDbErrorRepository{client: client}
}

func (d *DynamoDbErrorRepository) RecordError(id string, err error) {
	var httpError *error2.HttpError
	ok := errors.As(err, &httpError)
	if !ok {
		httpError = error2.NewInternalServerError(err.Error()).HttpError
	}
	item := errorEntry{
		ErrorID:   id,
		ErrorCode: httpError.HttpCode(),
		Message:   httpError.Error(),
	}
	data, err := attributevalue.MarshalMap(item)
	if err != nil {
		log.Fatalf("Failed to marshal error entry: %v", err)
	}

	_, err = d.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("ErrorTable"),
		Item:      data,
	})
	if err != nil {
		println(err.Error())
	}
	return
}

func (d *DynamoDbErrorRepository) GetError(id string) ([]byte, error) {
	selectedKeys := map[string]string{
		"PK": fmt.Sprintf(PKFormat, id),
	}
	key, err := attributevalue.MarshalMap(selectedKeys)

	if err != nil {
		return nil, error2.NewInternalServerError("error marshalling selectedKeys")
	}

	data, err := d.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("ErrorTable"),
		Key:       key,
	},
	)
	if err != nil {
		return nil, error2.NewInternalServerError(fmt.Sprintf("GetItem: %v\n", err))
	}
	if data.Item == nil {
		return nil, error2.NewInternalServerError(fmt.Sprintf("GetItem: Data not found.\n"))
	}
	var item *errorEntry
	err = attributevalue.UnmarshalMap(data.Item, &item)
	if err != nil {
		return nil, error2.NewInternalServerError(fmt.Sprintf("UnmarshalMap: %v\n", err))
	}
	bytes, err := json.Marshal(item)
	if err != nil {
		return nil, error2.NewInternalServerError(fmt.Sprintf("Error marshalling item retrieved: %v\n", err))
	}
	return bytes, nil
}

type errorEntry struct {
	ErrorID   string `dynamodbav:"errorId"`
	ErrorCode int    `dynamodbav:"errorCode"`
	Message   string `dynamodbav:"message"`
}
