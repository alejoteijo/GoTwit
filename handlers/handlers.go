package handlers

import (
	"github.com/alejoteijo/GoTwit/middleware"
	"github.com/alejoteijo/GoTwit/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

//Handlers config port & server listener
func Handlers() {
	router := mux.NewRouter()
	router.HandleFunc("/signup", middleware.CheckBD(routers.SignUp)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middleware.CheckBD(middleware.CheckJWT(routers.ShowProfile))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
