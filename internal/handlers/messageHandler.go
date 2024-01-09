package handlers

import (
	"App/internal/helpers"
	"App/internal/models"
	"App/internal/requests"
	"App/internal/resources"
	"encoding/json"
	"fmt"
	"net/http"
)

func (handler *HandlerService) StoreMessage(w http.ResponseWriter, r *http.Request) {
	var form requests.SendMessageRequest

	errPayload := ProcessRequest(&form, r, w)

	if errPayload != nil {
		return
	}

	var message models.Message

	helpers.FillStruct(&message, form.Data.Attributes)

	fmt.Printf("user: %+v\n", message)

	if err := handler.use.CreateMessage(&message, w); err != nil {
		success := models.Success{Success: false}

		successStatus, _ := json.Marshal(success)

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(successStatus)
		return
	}

	var messageRessource resources.MessageRessource

	resources.GenerateResource(&messageRessource, message, w)
}

func (handler *HandlerService) IndexUserChat(w http.ResponseWriter, r *http.Request) {
	var form requests.ChatRequest
	errPayload := ProcessRequest(&form, r, w)
	if errPayload != nil {
		return
	}
	var receiver models.Sender

	helpers.FillStruct(&receiver, form.Data.Attributes)

	found, err := handler.use.GetAllLinkedChat(receiver.Sender)

	if err != nil {
		success := models.Success{Success: false}

		successStatus, _ := json.Marshal(success)

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(successStatus)
		return
	}

	frontRessource, _ := json.Marshal(found)
	w.Write(frontRessource)

	return
}
