package dynamoDB

import (
	"AWS_Challenge/mocks"
	"AWS_Challenge/models"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGetDevice(t *testing.T) {
	item := map[string]*dynamodb.AttributeValue{
		"id": {S: aws.String("A1")},
	}

	tests := []struct {
		name           string
		item           map[string]*dynamodb.AttributeValue
		getItemError   error
		errorExpected  error
		outputExpected models.Device
	}{
		{name: "id not found", errorExpected: errors.New("could not find device with ID'A1'")},
		{name: "received", item: item, outputExpected: models.Device{Id: "A1"}},
		{name: "internal error", getItemError: errors.New("internal error"),
			errorExpected: errors.New("internal server error")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dbMock := new(mocks.DeviceDynamoDB)
			dbMock.On("GetItem", mock.Anything).Return(&dynamodb.GetItemOutput{
				Item: test.item,
			}, test.getItemError)
			service := Core{Db: dbMock}
			output, err := GetDevice(service.Db, "")
			if err == nil {
				assert.Nil(t, test.errorExpected)
			} else {
				assert.Error(t, test.errorExpected, err.Error())
			}
			assert.Equal(t, test.outputExpected, output)
		})
	}
}

func TestPutItemInDB(t *testing.T) {
	tests := []struct {
		name          string
		putItemError  error
		errorExpected error
	}{
		{name: "Created"},
		{name: "internal server error", putItemError: errors.New("internal error"), errorExpected: errors.New("internal error")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dbMock := new(mocks.DeviceDynamoDB)
			dbMock.On("PutItem", mock.Anything).Return(&dynamodb.PutItemOutput{}, test.putItemError)
			service := Core{
				Db: dbMock,
			}
			err := PutItemInDB(service.Db, models.Device{})

			if err == nil {
				assert.Nil(t, test.errorExpected)
			} else {
				assert.Error(t, test.errorExpected, err.Error())
			}
		})
	}
}
