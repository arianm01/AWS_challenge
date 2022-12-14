// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
	mock "github.com/stretchr/testify/mock"
)

// DeviceDynamoDB is an autogenerated mock type for the DeviceDynamoDB type
type DeviceDynamoDB struct {
	mock.Mock
}

// GetItem provides a mock function with given fields: _a0
func (_m *DeviceDynamoDB) GetItem(_a0 *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	ret := _m.Called(_a0)

	var r0 *dynamodb.GetItemOutput
	if rf, ok := ret.Get(0).(func(*dynamodb.GetItemInput) *dynamodb.GetItemOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.GetItemOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dynamodb.GetItemInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutItem provides a mock function with given fields: _a0
func (_m *DeviceDynamoDB) PutItem(_a0 *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	ret := _m.Called(_a0)

	var r0 *dynamodb.PutItemOutput
	if rf, ok := ret.Get(0).(func(*dynamodb.PutItemInput) *dynamodb.PutItemOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.PutItemOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dynamodb.PutItemInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDeviceDynamoDB interface {
	mock.TestingT
	Cleanup(func())
}

// NewDeviceDynamoDB creates a new instance of DeviceDynamoDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDeviceDynamoDB(t mockConstructorTestingTNewDeviceDynamoDB) *DeviceDynamoDB {
	mock := &DeviceDynamoDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
