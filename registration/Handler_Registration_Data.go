package registration

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/agpfven/WhatsApp_project/config"
	"github.com/agpfven/WhatsApp_project/utils"
	"github.com/chromedp/chromedp"
)

// struct for the template of the log in page
type logInData struct{
	QrImage string
}

// struct for the template of the log in page messenger
type Response struct {
	Message string `json:"message"`
}

func InitialPageLoader(w http.ResponseWriter, r *http.Request) {
	//Create template
	t, err := template.ParseFiles("log_in.html")
	config.HandleError(err)

	//Fill template
	p := logInData{QrImage: config.WebPagesLandingMsg}
	
	//Execute template into user browser
	if t.Execute(w, p) != nil{
		print(err)
	}
}

func InitialPageMsg(w http.ResponseWriter, r *http.Request) {
	//Create string channel (in order to use concurrency)
	qrData := make(chan string)

	//Retrive qr from whatsapp web page and handle all data retrieval
	go RegistrationDataHandler(qrData)

	//Prepare message
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: <-qrData}
	close(qrData)
	json.NewEncoder(w).Encode(response)	
}

func RegistrationDataHandler(ch chan string) (){
	//Initializing Browser Context (if headless mode is not disabled this doesn't work)
	log.Println(config.PromptStartBrowser)
	allocatorCtx, allocatorCancel := chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:], 
			chromedp.Flag("headless", false),
			chromedp.UserDataDir("myUsers"),
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

	//Extract QR data from wss page
	GetQrCode(ch, browserCtx)

	//Retrive User's phone number
	userPhoneNumber := RetriveNumber(browserCtx)
	log.Printf("Users phone number: %s", userPhoneNumber)

	//This is done in order to let the whatsapp web page to synchronize with the mobile app
	//time.Sleep(1 * time.Minute)
	time.Sleep(3 * time.Second)

	chromedp.Cancel(browserCtx)

	//Next step of the process
	go HandlerRegistrationUpload(userPhoneNumber, allocatorCtx, browserCtx)
}

//This function retrives the user phone number
func RetriveNumber(givenCtx context.Context) (string){
	//This function checks the number of the user using a channel
	utils.SelectContact(givenCtx)

	var data map[string] string
	err := chromedp.Run(givenCtx,
		chromedp.Attributes(config.QrDivFullXPATH3, &data),
	)
	if err != nil {
		log.Fatal(err)
	}

	return data["title"]
}

//This functions retrives the image of the qr code of the wss page
func GetQrCode(auxiliarCh chan string, browserCtx context.Context) () {
	log.Println("Extracting QR data...")

	//Where the attributes data will be stored
	var data map[string]string

	//Go to Wss webpage, wait for QR and extract its information
	err := chromedp.Run(browserCtx,
		chromedp.Navigate("http://web.whatsapp.com/"),
		chromedp.WaitEnabled(config.QrDivByQuery1, chromedp.ByQuery),
		chromedp.Attributes(config.QrDivFullXPATH1, &data),
		)
	if err != nil {
		log.Fatal(err)
	}

	//Pass the QR data information to the channel
	println(data["data-ref"])
	auxiliarCh <- data["data-ref"]
}