package error

import (
	"TestProject/internal/config"
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go"
	smithyEndpoints "github.com/aws/smithy-go/endpoints"
	"net/url"
	"os"
)

type localEndpointResolver struct {
	Host string
	Port string
}

func (r *localEndpointResolver) ResolveEndpoint(_ context.Context, _ dynamodb.EndpointParameters) (smithyEndpoints.Endpoint, error) {
	var zero smithyEndpoints.Endpoint
	urlParsed, err := url.Parse(fmt.Sprintf("http://%s:%s", r.Host, r.Port))
	if err != nil {
		return zero, errors.New("error while parsing dynamo url")
	}
	return smithyEndpoints.Endpoint{
		URI:        *urlParsed,
		Headers:    nil,
		Properties: smithy.Properties{},
	}, nil
	return zero, nil
}

var DynamoErrorConnection *dynamodb.Client

func NewDynamoDBConnection(config *config.DynamoDbConfig) (*dynamodb.Client, error) {
	if DynamoErrorConnection != nil {
		return DynamoErrorConnection, nil
	}
	os.Setenv("AWS_ACCESS_KEY_ID", config.AccessKey)
	os.Setenv("AWS_SECRET_ACCESS_KEY", config.KeyId)
	os.Setenv("AWS_REGION", config.Region)
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	client := dynamodb.NewFromConfig(cfg, dynamodb.WithEndpointResolverV2(&localEndpointResolver{config.Host, config.Port}))
	DynamoErrorConnection = client
	err = setUpDynamoTables(DynamoErrorConnection)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func setUpDynamoTables(client *dynamodb.Client) error {
	tableExist, err := checkTableExists(client, "ErrorsTable")
	if err != nil {
		return err
	}
	if tableExist {
		return nil
	}
	input := &dynamodb.CreateTableInput{
		TableName: aws.String("ErrorsTable"),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("errorId"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("errorCode"),
				AttributeType: types.ScalarAttributeTypeN,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("errorId"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("errorCode"),
				KeyType:       types.KeyTypeRange,
			},
		},
		BillingMode: types.BillingModePayPerRequest,
	}
	_, err = client.CreateTable(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("error creant la taula: %w", err)
	}
	return nil
}

func checkTableExists(client *dynamodb.Client, tableName string) (bool, error) {
	output, err := client.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		return false, err
	}
	for _, name := range output.TableNames {
		if name == tableName {
			return true, nil
		}
	}
	return false, nil
}
