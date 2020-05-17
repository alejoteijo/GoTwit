package main

import (
	"github.com/alejoteijo/GoTwit/bd"
	"github.com/alejoteijo/GoTwit/handlers"
	"log"
)

func main() {
	if !bd.CheckConnection() {
		log.Fatal("Can't connect to database")
	}
	handlers.Handlers()
}
