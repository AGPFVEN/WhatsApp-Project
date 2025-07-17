package utils

import (
	"context"
	"log"
	"time"

	"github.com/agpfven/WhatsApp_project/config"
	"github.com/chromedp/chromedp"
)

func SelectContact(givenCtx context.Context) (){
	//This function checks the number of the user
	log.Println("Selecting contact ...")
	err := chromedp.Run(givenCtx,
		chromedp.WaitNotPresent(config.QrDivByQuery1, chromedp.ByQuery),
		chromedp.WaitNotPresent(config.QrDivByQuery2, chromedp.ByQuery),
		chromedp.Sleep(3 * time.Second),
		//chromedp.WaitReady("body"),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Going to click")
	err = chromedp.Run(givenCtx,
		chromedp.Click(config.QrDivFullXPATH2),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Select contact page ready")
}