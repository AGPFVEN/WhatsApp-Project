package main

import (
	"log"
	"net/http"

	"github.com/agpfven/WhatsApp_project/config"
	controller_registration "github.com/agpfven/WhatsApp_project/controller_Registration"
	//"github.com/agpfven/WhatsApp_project/controller"
)

func main(){
	// Define routes (make them safe from saturation) --------
	http.HandleFunc(config.WebPagesLanding, controller_registration.InitialPageLoader)
	http.HandleFunc(config.WebPagesLandingMsg, controller_registration.InitialPageQrMsg)
	http.HandleFunc(config.WebPagesLandingMsg1, controller_registration.InitialPagePhoneMsg)

	port := ":3000"
	log.Println("Serving port " + port)
	http.ListenAndServe(port, nil)
}