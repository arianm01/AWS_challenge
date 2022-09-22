package controllers

import (
	"AWS_Challenge/dynamoDB"
	"AWS_Challenge/models"
	"AWS_Challenge/services"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateDevice(writer http.ResponseWriter, request *http.Request) {
	var device models.Device
	_ = json.NewDecoder(request.Body).Decode(&device)

	if err := validator.New().Struct(device); err != nil {
		log.Println(err)
		services.SendError(writer, "invalid device properties", http.StatusBadRequest)
		return
	}
	db := dynamoDB.GetDynamodb()
	service := &dynamoDB.Core{
		Db: db,
	}
	err := dynamoDB.PutItemInDB(service.Db, device)
	if err != nil {
		log.Println(err)
		services.SendError(writer, "internal server error", http.StatusInternalServerError)
		return
	}
	services.SendSuccess(writer, device, http.StatusCreated)
}

func GetDevice(writer http.ResponseWriter, request *http.Request) {
	db := dynamoDB.GetDynamodb()
	params := mux.Vars(request)
	service := &dynamoDB.Core{
		Db: db,
	}
	result, err := dynamoDB.GetDevice(service.Db, params["id"])
	if err != nil {
		if err.Error() == "internal Server Error" {
			services.SendError(writer, err.Error(), http.StatusInternalServerError)
		} else {
			services.SendError(writer, err.Error(), http.StatusNotFound)
		}
		return
	}
	services.SendSuccess(writer, result, http.StatusOK)
}
