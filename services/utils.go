package services

import (
	"AWS_Challenge/models"
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

func SendError(writer http.ResponseWriter, err string, status int) {
	writer.WriteHeader(status)
	result, _ := json.Marshal(Error{
		Message: err,
	})
	_, _ = writer.Write(result)
}

func SendSuccess(writer http.ResponseWriter, item models.Device, status int) {
	writer.WriteHeader(status)
	itemMarshal, _ := json.Marshal(item)
	_, _ = writer.Write(itemMarshal)
}
