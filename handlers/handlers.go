package handlers

import (
	"../middleware"
	"../routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

//Handlers config port & server listener
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middleware.CheckBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
