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
	err := dynamoDB.PutItemInDB(*db, device)
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
	result, err := dynamoDB.GetDevice(*db, params["id"])
	if err != nil {
		if err.Error() == "internal Server Error" {
			services.SendError(writer, err.Error(), http.StatusInternalServerError)
		} else if err.Error() == "device is not valid" {
			services.SendError(writer, err.Error(), http.StatusNotAcceptable)
		} else {
			services.SendError(writer, err.Error(), http.StatusNotFound)
		}
		return
	}
	services.SendSuccess(writer, result, http.StatusOK)
}
