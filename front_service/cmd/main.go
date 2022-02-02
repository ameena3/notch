package main

import (
	"log"
	"net/http"

	handlers "github.com/ameena3/notch_project/front_service"
	client "github.com/ameena3/notch_project/front_service/services/back_client"
)

func main() {
	log.Println("Strating the front service...")
	hsnlr := handlers.NewFrontService(client.NewClient())
	http.HandleFunc("/login", hsnlr.LoginHandler)
	http.HandleFunc("/home", hsnlr.HomeHandler)
	http.HandleFunc("/checkout", hsnlr.CheckoutHandler)
	http.HandleFunc("/cart", hsnlr.CartHandler)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
