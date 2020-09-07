package router

import (
	"cdi/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/add", middleware.CreateIdea)
	router.HandleFunc("/api/see", middleware.ShowData)
	router.HandleFunc("/api/make/{t}/{d}", middleware.MakeIdea)
	router.HandleFunc("/api/remove/{id}", middleware.DeleteIdea)

	return router
}
