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
	router.Post("/send-message", handlerService.StoreMessage)
	router.Post("/get-all-chat", handlerService.IndexUserChat)
	// ROUTE ENVOYER UN PREMIER MESSAGE
	// ROUTE RECUPERER TOUT LES USERS AVEC QUI J'AI EU UNE CONV

	router.Group(func(r chi.Router) {
		// Route User
		r.Get("/profils", handlerService.IndexUser)
		r.Patch("/profil/user/{id}", handlerService.Update)
		r.Delete("/profil/user/{id}", handlerService.Delete)
	})

	router.NotFound(notfound)
}
