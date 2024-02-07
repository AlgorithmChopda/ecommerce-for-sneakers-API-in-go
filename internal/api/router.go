package api

import (
	"net/http"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app"
	"github.com/gorilla/mux"
)

func NewRouter(deps app.Dependencies) *mux.Router {
	router := mux.NewRouter()

	// UserRoutes
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/login", LoginUserHandler(deps.UserService))
	userRouter.HandleFunc("/register", RegisterUserHandler(deps.UserService)).Methods(http.MethodPost)

	return router
}
