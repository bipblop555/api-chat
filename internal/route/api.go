package route

import (
	"App/internal/handlers"
	"App/internal/middlewares"
	"github.com/go-chi/chi/v5"
)

func route(router *chi.Mux, handlerService *handlers.HandlerService) {

	router.Use(middlewares.SetJSONHeaders)

	router.Post("/get-user", handlerService.IndexUser)
	router.Post("/signup", handlerService.StoreUser)
	router.Post("/login", handlerService.Login)

	router.Post("/get-all-chat", handlerService.IndexUserChat)

	router.Post("/send-message", handlerService.StoreMessage)
	router.Delete("/delete-message/{id}", handlerService.Delete)
	router.Patch("/update-message/{id}", handlerService.Update)
	router.Get("/messages/get-all/{senderId}/{receiverId}", handlerService.GetAllMessageFrom)

	router.NotFound(notfound)
}
