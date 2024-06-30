package utils

import (
	"context"
	"log"
	"net/http"
	"text/template"

	"github.com/agpfven/WhatsApp_project/config"
	"github.com/chromedp/chromedp"
)

func SelectContact(givenCtx context.Context) (){
	//This function checks the number of the user
	log.Println("Selecting contact ...")
	err := chromedp.Run(givenCtx,
		chromedp.WaitNotPresent("_aly_", chromedp.ByQuery),
		chromedp.WaitReady(config.QrDivByQuery2, chromedp.ByQuery),
		chromedp.Click(config.QrDivFullXPATH2),
		//chromedp.WaitReady("body"),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Select contact page ready")
}

func SendHTMLToBrowser(filename string, w http.ResponseWriter, data any) (error){
		//Load html file with qr code
		t, err := template.ParseFiles(filename)
		if err != nil {
			print(err)
		}
		
		//Execute template into user browser
		if t.Execute(w, data) != nil{
			print(err)
		}
	return nil
}