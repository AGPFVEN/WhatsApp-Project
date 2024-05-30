package main

import (
	"net/http"

	"github.com/agpfven/WhatsApp_project/config"
	"github.com/agpfven/WhatsApp_project/registration"
	//"github.com/agpfven/WhatsApp_project/controller"
)

 func main(){
	// Define routes
	http.HandleFunc(config.WebPagesHome, registration.InitialPageLoader)

	http.ListenAndServe(":3000", nil)
}