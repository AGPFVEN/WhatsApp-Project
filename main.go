package main

import (
	"log"
	"net/http"

	"github.com/agpfven/WhatsApp_project/config"
	controller_registration "github.com/agpfven/WhatsApp_project/controller_Registration"
	//"github.com/agpfven/WhatsApp_project/controller"
)

 func main(){
	// Define routes
	http.HandleFunc(config.WebPagesLanding, controller_registration.InitialPageLoader)
	http.HandleFunc(config.WebPagesLandingMsg, controller_registration.InitialPageQrMsg)

	port := ":3000"
	log.Println("Serving port " + port)
	http.ListenAndServe(port, nil)
}