package googlefunctions

import (
	"context"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

// Suma dos enteros y devuelve el resultado.
func getAndRetrieveLogInQr() {
    //Initializing Browser Context (if headless mode is not disabled this doesn't work)
	log.Println("Creando browser")
	allocatorCtx, allocatorCancel := chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:], 
			chromedp.Flag("headless", false),
			chromedp.UserDataDir(os.Getenv("USER_DATA_DIR")),
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
}