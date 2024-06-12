package config

import "log"

const(
	//Web pages
	WebPagesLanding = "/landing"
	WebPagesLandingMsg = "/landing-msg"
	WebPagesLandingMsg1 = "/landing-msg1"

	//Console Prompts
	PromptStartBrowser = "Initializing Browser..."

	//Wss selectors
	//Initial page
	QrDivByQuery1 = "._akaz"
	QrDivFullXPATH1 = "/html/body/div[2]/div/div/div[2]/div[3]/div[1]/div/div/div[2]/div"

	//Select contact	
	QrDivByQuery2 = ".xurb0ha"
	QrDivFullXPATH2 = "/html/body/div[1]/div/div/div[2]/div[3]/header/div[2]/div/span/div[4]/div"

	//Own number
	QrDivFullXPATH3 = "/html/body/div[1]/div/div/div[2]/div[2]/div[1]/span/div/span/div/div[2]/div[3]/div/div/div[11]/div[1]/div/div[2]/div[1]/div/div/span[1]"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}