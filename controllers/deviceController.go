package controllers

import (
	"AWS_Challenge/models"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

func CreateDevice(writer http.ResponseWriter, request *http.Request) {
	var device models.Device
	_ = json.NewDecoder(request.Body).Decode(&device)

	if err := validator.New().Struct(device); err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		result, _ := json.Marshal(Error{
			Message: "invalid device attribute",
		})
		_, _ = writer.Write(result)
		return
	}

}

func GetDevice(writer http.ResponseWriter, request *http.Request) {

}
