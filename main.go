package main

import (
	"log"
	"net/http"

	"github.com/agpfven/WhatsApp_project/config"
	"github.com/agpfven/WhatsApp_project/registration"
)

 func main(){
	// Define routes for First landing 
	http.HandleFunc(config.WebPagesFirstLand, registration.InitialPageLoader)
	http.HandleFunc(config.WebPagesMessageToFirstLand, registration.InitialPageLoader)

	const port = ":3000"
	
	log.Println("Serving port " + port)
	http.ListenAndServe(port, nil)
}