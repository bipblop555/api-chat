package handlers

import (
	"App/internal/helpers"
	"App/internal/models"
	"App/internal/requests"
	"App/internal/resources"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (handler *HandlerService) StoreMessage(w http.ResponseWriter, r *http.Request) {
	var form requests.SendMessageRequest
	var message models.Message

	errPayload := ProcessRequest(&form, r, w)

	if errPayload != nil {
		return
	}

	helpers.FillStruct(&message, form.Data.Attributes)

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

	found, err := handler.use.GetAllLinkedChat(receiver.Sender_id)

	if err != nil {
		success := models.Success{Success: false}

		successStatus, _ := json.Marshal(success)

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(successStatus)
		return
	}

	var chatRessource []resources.ChatRessource

	resources.GenerateResource(&chatRessource, found, w)
}

func (handler *HandlerService) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	message := models.Message{}

	err := handler.use.ByID(id, &message)
	if err != nil {
		// Gérer l'erreur
		fmt.Println(err)
	}

	// Utilisez simplement Delete avec l'instance de la structure de modèle
	request := models.InitGorm.Db.Where("id = ?", id).Delete(&message)
	fmt.Println(request)
	/*if request != nil {
		success := models.Success{Success: false}

		successStatus, _ := json.Marshal(success)

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(successStatus)
		return
	} else {

		success := models.Success{Success: true}

		successStatus, _ := json.Marshal(success)

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(successStatus)
	}*/
}

func (handler *HandlerService) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	message := models.Message{}
	err := handler.use.ByID(id, &message)

	if err != nil {
		// Gérer l'erreur
		fmt.Println(err)
	}

	fmt.Println(message)

	err = handler.use.Update(&message, id, w)
	if err != nil {
		// Gérer l'erreur
		fmt.Println(err)
	}

	var messageRessource resources.MessageRessource

	resources.GenerateResource(&messageRessource, message, w)
}

func (handler *HandlerService) GetAllMessageFrom(w http.ResponseWriter, r *http.Request) {
	senderId := chi.URLParam(r, "senderId")
	receiverId := chi.URLParam(r, "receiverId")

	err, _ := handler.use.GetAllMessagesFromUser(senderId, receiverId)
	if err != nil {
		// Gérer l'erreur
		fmt.Println(err)
	}

	var messageRessource []resources.MessageRessource

	resources.GenerateResource(&messageRessource, err, w)
}
