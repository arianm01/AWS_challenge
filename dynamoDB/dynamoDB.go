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

func GetDynamodb() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("DYNAMODB_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	}))
	return dynamodb.New(sess)
}

func PutItemInDB(db dynamodb.DynamoDB, device models.Device) error {
	deviceMap, _ := dynamodbattribute.MarshalMap(device)
	tableName := "devices"
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

func GetDevice(db dynamodb.DynamoDB, id string) (models.Device, error) {
	tableName := "devices"
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
	err = dynamodbattribute.UnmarshalMap(result.Item, &device)
	if err != nil {
		return models.Device{}, errors.New("device is not valid")
	}
	return device, nil
}
