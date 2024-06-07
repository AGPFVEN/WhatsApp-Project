package controller_registration

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	//"os"
	//"strconv"
	"sync"
	"text/template"
	"time"

	"github.com/agpfven/WhatsApp_project/config"
	"github.com/agpfven/WhatsApp_project/utils"
	"github.com/chromedp/chromedp"
)

const registration_qr_phone_size = 3

var registration_qr_phone [registration_qr_phone_size][2]string


// struct for the template of the log in page
type logInData struct{
	QrMsgURL string
	PhoneMsgURL string
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
	p := logInData{QrMsgURL: config.WebPagesLandingMsg,
		PhoneMsgURL: config.WebPagesLandingMsg1,
	}
	
	//Execute template into user browser
	if t.Execute(w, p) != nil{
		print(err)
	}
}

func InitialPageQrMsg(w http.ResponseWriter, r *http.Request) {
	//Create string channel (in order to use concurrency)
	//qrData := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	var qrData string

	//Retrive qr from whatsapp web page and handle all data retrieval
	go registrationDataHandler(&qrData, wg)

	//Prepare message
	w.Header().Set("Content-Type", "application/json")
	wg.Wait()
	response := Response{qrData}
	//close(qrData)
	json.NewEncoder(w).Encode(response)	
}

func InitialPagePhoneMsg(w http.ResponseWriter, r *http.Request) {
	log.Println("1er paso")
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println("Received qr data: " + string(body))
		for i, j := 0, 0; j == 0; i++{
		  if (string(body) == registration_qr_phone[i][0]){
			io.WriteString(w, registration_qr_phone[i][1])
			registration_qr_phone[i][0] = ""
			registration_qr_phone[i][1] = ""
			j = 1
		  }

		  if (i == 2){
			i = 0
		  }
		}
		//io.WriteString(w, "Received text: %s" + string(body))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

//func registrationDataHandler(ch chan string) (){
func registrationDataHandler(qrDataPtrLocal *string, wgLocal *sync.WaitGroup) (){
	var erg string

	//for i:=0; i < registration_qr_phone_size; i++{
		//erg = "myUsers/erga" + strconv.Itoa(i)
		//_, err := os.Stat(erg)
		//if os.IsNotExist(err){
			//log.Println("Going good")
			//i = registration_qr_phone_size
		//}
		//log.Println("Not going good")
	//}

	//Initializing Browser Context (if headless mode is not disabled this doesn't work)
	erg = "myUsers/erga0"
	log.Println(config.PromptStartBrowser)
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

	//Extract QR data from wss page
	//getQrCode(qrDataPtrLocal, browserCtx, wgLocal)

	//Retrive User's phone number
	userPhoneNumber := retriveNumber(browserCtx)
	log.Printf("Users phone number: %s", userPhoneNumber)

	//
	go storeQrPhone(*qrDataPtrLocal, userPhoneNumber, wgLocal)

	//This is done in order to let the whatsapp web page to synchronize with the mobile app
	//time.Sleep(1 * time.Minute)
	time.Sleep(3 * time.Second)

	chromedp.Cancel(browserCtx)


	//Next step of the process
	go HandlerRegistrationUpload(userPhoneNumber, allocatorCtx, browserCtx)
}

func storeQrPhone(qrData string, phoneNumber string, wg *sync.WaitGroup){
	for i:= 0; i < registration_qr_phone_size; i++{
		if(registration_qr_phone[i][0] == ""){
			wg.Wait()
			registration_qr_phone[i][0] = qrData
			registration_qr_phone[i][1] = phoneNumber
			
		}
	}

	//NEED TO CARE WHEN THIS HAPPEN --------------------------------------------
	log.Println("Registration is full ...")
}

//This function retrives the user phone number
func retriveNumber(givenCtx context.Context) (string){
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
//func getQrCode(auxiliarCh chan string, browserCtx context.Context) () {
func getQrCode(extractedQr *string, browserCtx context.Context, wgLocal *sync.WaitGroup) () {
	log.Println("Extracting QR data...")

	//Where the attributes data will be stored
	var data map[string]string

	//Go to Wss webpage, wait for QR and extract its information
	err := chromedp.Run(browserCtx,
		chromedp.WaitEnabled(config.QrDivByQuery1, chromedp.ByQuery),
		chromedp.Attributes(config.QrDivFullXPATH1, &data),
		)
	if err != nil {
		log.Fatal(err)
	}

	//Pass the QR data information to the channel
	println(data["data-ref"])
	*extractedQr = data["data-ref"]
	wgLocal.Done()
}