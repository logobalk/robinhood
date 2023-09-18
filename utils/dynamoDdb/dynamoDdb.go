package dynamoddb

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var defaultClient *DynamoDdb
var localClient *DynamoDdb

type DynamoDdb struct {
	*dynamodb.Client
}

type LocalEndpointResolver struct{}
type LocalStackEndpointResolver struct{}

func New() *DynamoDdb {
	if defaultClient != nil {
		return defaultClient
	}

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	defaultClient = &DynamoDdb{
		Client: svc,
	}

	return defaultClient
}

func NewLocal() *DynamoDdb {
	if localClient != nil {
		return localClient
	}

	cfg, err := config.LoadDefaultConfig(context.Background(),
		// config.WithHTTPClient(httpClient),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(&LocalEndpointResolver{}),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "local", SecretAccessKey: "local", SessionToken: "local",
			},
		}),
	)
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	localClient = &DynamoDdb{
		Client: svc,
	}

	return localClient
}

func NewLocalStack() *DynamoDdb {
	if localClient != nil {
		return localClient
	}

	cfg, err := config.LoadDefaultConfig(context.Background(),
		// config.WithHTTPClient(httpClient),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(&LocalStackEndpointResolver{}),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "local", SecretAccessKey: "local", SessionToken: "local",
			},
		}),
	)
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	localClient = &DynamoDdb{
		Client: svc,
	}

	return localClient
}

func (r *LocalStackEndpointResolver) ResolveEndpoint(service, region string, options ...interface{}) (aws.Endpoint, error) {
	endpoint := os.Getenv("DYNAMODB_ENDPOINT_LOCAL_STACK")
	if endpoint == "" {
		endpoint = "http://localhost:4566"
	}

	return aws.Endpoint{
		URL:           endpoint,
		SigningRegion: region,
	}, nil
}

func (r *LocalEndpointResolver) ResolveEndpoint(service, region string, options ...interface{}) (aws.Endpoint, error) {
	endpoint := os.Getenv("DYNAMODB_ENDPOINT_LOCAL")
	if endpoint == "" {
		endpoint = "http://localhost:8000"
	}

	return aws.Endpoint{
		URL:           endpoint,
		SigningRegion: region,
	}, nil
}
