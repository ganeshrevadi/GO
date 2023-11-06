package routes

import (
	"ganeshrevadi/GO/Mongo/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllCollection).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatch).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOne).Methods("DELETE")
	router.HandleFunc("/api/deleteAll", controller.DeleteAll).Methods("DELETE")

	return router
}
