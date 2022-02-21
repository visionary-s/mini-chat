package routers

import (
	"mini-chat/middlewares/home"
	"mini-chat/middlewares/user"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	userRouter := router.PathPrefix("/user").Subrouter()

	router.HandleFunc("/", home.Index).Methods("GET", "OPTIONS")
	userRouter.HandleFunc("/login", user.Login).Methods("POST", "OPTIONS")
	userRouter.HandleFunc("/logout", user.Logout).Methods("GET", "OPTIONS")

	return router
}
