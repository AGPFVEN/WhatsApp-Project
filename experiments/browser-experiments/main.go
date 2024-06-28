package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
)

func main(){
	erg := "../../myUsers/erga0"

	allocatorCtx, allocatorCancel := chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:], 
			chromedp.Flag("headless", false),
			chromedp.UserDataDir(erg),
		)...
	)
	defer allocatorCancel()

	//Browser is closed at the end of this function
	browserCtx, browserCancel := chromedp.NewContext(allocatorCtx)
	defer browserCancel()

	//Go to Wss webpage, wait for QR and extract its information
	err := chromedp.Run(browserCtx,
		chromedp.Navigate("http://web.whatsapp.com/"),
		)
	if err != nil {
		log.Fatal(err)
	}

	//Where the attributes data will be stored
	var data map[string]string

	//Go to Wss webpage, wait for QR and extract its information
	err = chromedp.Run(browserCtx,
		chromedp.WaitEnabled("._ah_-", chromedp.ByQuery),
		chromedp.Attributes("/html/body/div[1]/div/div/div[2]/div[3]/div/div[1]/div/div[2]/div[1]", &data),
		)
	if err != nil {
		log.Fatal(err)
	}

	//Pass the QR data information to the channel
	println(fmt.Sprint(data))
}