package handlers

import (
	"App/internal/helpers"
	"App/internal/models"
	"App/internal/requests"
	"App/internal/resources"
	"encoding/json"
	"fmt"
	"net/http"
	// "net/http"
)

// Methode create pour ajotuer new user "POST / signup"
func (handler *HandlerService) StoreUser(w http.ResponseWriter, r *http.Request) {
	var form requests.StoreUserRequest

	errPayload := ProcessRequest(&form, r, w)

	if errPayload != nil {
		return
	}

	var user models.User

	form.Data.Attributes.Password = helpers.HashPassword(form.Data.Attributes.Password)

	helpers.FillStruct(&user, form.Data.Attributes)

	fmt.Printf("user: %+v\n", user)

	if err := handler.use.Create(&user, w); err != nil {
		success := models.Success{Success: false}

		successStatus, _ := json.Marshal(success)

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(successStatus)
		return
	}

	var userResource resources.UserResource

	resources.GenerateResource(&userResource, user, w)
}

func (handler *HandlerService) Login(w http.ResponseWriter, r *http.Request) {
	var form requests.UserLoginRequest

	errPayload := ProcessRequest(&form, r, w)
	if errPayload != nil {
		fmt.Println("errPayload:", errPayload)
		return
	}

	user, err := handler.use.Authenticate(form.Data.Attributes.Email, form.Data.Attributes.Password)

	if err != nil {
		success := models.Success{Success: false}
		successStatus, _ := json.Marshal(success)

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(successStatus)
		return
	}

	fmt.Printf("user: %+v\n", user)

	handler.setCookieFromJWT(w, user.Email)

	var userResource resources.UserResource

	resources.GenerateResource(&userResource, user, w)
}

func (handler *HandlerService) IndexUser(w http.ResponseWriter, r *http.Request) {
	var form requests.FindUserRequest

	errPayload := ProcessRequest(&form, r, w)

	if errPayload != nil {
		return
	}

	var user models.User

	helpers.FillStruct(&user, form.Data.Attributes)
	found, err := handler.use.ByUserName(user.Username)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		errMsg, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Write(errMsg)
		return
	}

	var userResource []resources.UserResource
	resources.GenerateResource(&userResource, found, w)
}
