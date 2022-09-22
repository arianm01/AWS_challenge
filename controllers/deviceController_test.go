package controllers

import (
	"AWS_Challenge/dynamoDB"
	"AWS_Challenge/models"
	"AWS_Challenge/services"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetDevice(t *testing.T) {

	_ = os.Setenv("ACCESS_KEY_ID", os.Getenv("AWS_ACCESS_KEY_ID"))
	_ = os.Setenv("SECRET_ACCESS_KEY", os.Getenv("AWS_SECRET_ACCESS_KEY"))
	_ = os.Setenv("AWS_TABLE_NAME", os.Getenv("AWS_TABLE"))

	input := models.Device{
		Id:          "A1",
		DeviceModel: "2016",
		Name:        "Car",
		Note:        "very beautiful",
		Serial:      "123456",
	}
	db := dynamoDB.GetDynamodb()
	service := &dynamoDB.Core{
		Db: db,
	}
	err := dynamoDB.PutItemInDB(service.Db, input)
	if err != nil {
		t.Fatal("error occurred while device creating", err)
	}

	tests := []struct {
		name   string
		id     string
		status int
		output interface{}
	}{
		{name: "success", status: http.StatusOK, output: input, id: input.Id},
		{name: "failure", status: http.StatusNotFound, output: services.Error{
			Message: "Could not find device with ID'not-found'",
		}, id: "not-found"},
		{name: "internal error", status: http.StatusInternalServerError, output: services.Error{
			Message: "internal Server Error",
		}, id: input.Id},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.name == "internal error" {
				_ = os.Unsetenv("AWS_TABLE_NAME")
			}
			router := mux.NewRouter()
			router.HandleFunc("/devices/{id}", GetDevice).Methods("GET")
			req, _ := http.NewRequest(http.MethodGet, "/devices/"+test.id, nil)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(t, test.status, res.Code)

			if test.status == http.StatusOK {
				var device models.Device
				_ = json.Unmarshal(res.Body.Bytes(), &device)
				assert.Equal(t, test.output.(models.Device), device)
			} else {
				var message services.Error
				_ = json.Unmarshal(res.Body.Bytes(), &message)
				assert.Equal(t, test.output.(services.Error), message)
			}
		})
	}

}

func TestCreateDevice(t *testing.T) {
	input := models.Device{
		Id:          "A1",
		DeviceModel: "2016",
		Name:        "Car",
		Note:        "very beautiful",
		Serial:      "123456",
	}
	tests := []struct {
		name   string
		input  models.Device
		status int
		output interface{}
	}{
		{name: "ok", input: input, status: http.StatusCreated, output: input},
		{name: "invalid", input: models.Device{Id: "wrong"}, status: http.StatusBadRequest, output: services.Error{
			Message: "invalid device properties",
		}},
		{name: "internal error", input: input, status: http.StatusInternalServerError, output: services.Error{
			Message: "internal server error",
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_ = os.Unsetenv("AWS_REGION")
			_ = os.Unsetenv("AWS_TABLE_NAME")
			_ = os.Setenv("ACCESS_KEY_ID", os.Getenv("AWS_ACCESS_KEY_ID"))
			_ = os.Setenv("SECRET_ACCESS_KEY", os.Getenv("AWS_SECRET_ACCESS_KEY"))
			if test.name != "internal error" {
				_ = os.Setenv("AWS_REGION", "us-east-1")
				_ = os.Setenv("AWS_TABLE_NAME", "aws_table_device")
			}

			router := mux.NewRouter()
			router.HandleFunc("/devices", CreateDevice).Methods("POST")

			marshal, _ := json.Marshal(test.input)
			req, _ := http.NewRequest(http.MethodPost, "/devices", bytes.NewBuffer(marshal))

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(t, test.status, res.Code)

			// check status
			if res.Code == http.StatusCreated {
				var device models.Device
				_ = json.Unmarshal(res.Body.Bytes(), &device)
				assert.Equal(t, test.output.(models.Device), device)
			} else {
				var err services.Error
				_ = json.Unmarshal(res.Body.Bytes(), &err)
				assert.Equal(t, test.output.(services.Error), err)
			}

		})
	}
}
