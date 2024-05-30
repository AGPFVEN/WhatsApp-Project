package config

import (
	"os"

	"golang.org/x/oauth2"
)

const(
	//Web pages
	WebPagesHome = "/landing"

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
	QrDivFullXPATH3 = "/html/body/div[1]/div/div/div[2]/div[2]/div[1]/span/div/span/div/div[2]/div[5]/div/div/div[11]/div[1]/div/div[2]/div[1]/div/div/span[1]"
)

type product struct{
	phone string
	zip_bytes []byte
}

//NEED TO CHANGE THIS IN ORDER TO DO CLIENT CREDENTIAL FLOW
func LoadOauthConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID: os.Getenv("WP_CLIENT_ID"),
		ClientSecret: os.Getenv("WP_TENANT_ID"),
		Scopes: []string{"files.readwrite", "offline_access"},
		Endpoint: oauth2.Endpoint{
			AuthURL: os.Getenv("ONEDRIVE_AUTH_ENDPOINT"),
			TokenURL: os.Getenv("ONEDRIVE_TOKEN_ENDPOINT"),
		},
	}

	return conf
}