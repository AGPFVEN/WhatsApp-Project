package main

import (
	"log"
	"net/http"

	"github.com/agpfven/WhatsApp_project/config"
	"github.com/agpfven/WhatsApp_project/registration"
	//"github.com/agpfven/WhatsApp_project/controller"
)

 func main(){
	// Define routes
	http.HandleFunc(config.WebPagesLanding, registration.InitialPageLoader)
	http.HandleFunc(config.WebPagesLandingMsg, registration.InitialPageMsg)

	port := ":3000"
	log.Println("Serving port " + port)
	http.ListenAndServe(port, nil)
}