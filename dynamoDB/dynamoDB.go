package dynamoDB

import (
	"AWS_Challenge/models"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
)

type DeviceDynamoDB interface {
	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
}

type Core struct {
	Db DeviceDynamoDB
}

func GetDynamodb() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("ACCESS_KEY_ID"),
			os.Getenv("SECRET_ACCESS_KEY"), ""),
	}))
	return dynamodb.New(sess)
}

func PutItemInDB(db DeviceDynamoDB, device models.Device) error {
	deviceMap, _ := dynamodbattribute.MarshalMap(device)
	tableName := os.Getenv("AWS_TABLE_NAME")
	data := &dynamodb.PutItemInput{
		Item:      deviceMap,
		TableName: aws.String(tableName),
	}
	_, err := db.PutItem(data)
	if err != nil {
		return err
	}
	return nil
}

func GetDevice(db DeviceDynamoDB, id string) (models.Device, error) {
	tableName := os.Getenv("AWS_TABLE_NAME")
	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return models.Device{}, errors.New("internal Server Error")
	}
	if result.Item == nil {
		msg := "Could not find device with ID'" + id + "'"
		return models.Device{}, errors.New(msg)
	}
	var device models.Device
	_ = dynamodbattribute.UnmarshalMap(result.Item, &device)
	return device, nil
}
