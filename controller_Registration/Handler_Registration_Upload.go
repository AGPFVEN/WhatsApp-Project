package controller_registration

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"

	"github.com/agpfven/WhatsApp_project/utils"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func HandlerRegistrationUpload(phoneNumber string, isAllocatorClosed context.Context, isBrowserClosed context.Context) (){
	//Check if the browser is closed
	if (struct {}{} == <-isBrowserClosed.Done() && struct{}{} == <-isAllocatorClosed.Done()){
		log.Println("Browser Closed")
	}

	//Make phone number phoneHashed
	phoneHashed := sha256.New()
	phoneHashed.Write([]byte(phoneNumber))

	//Compress browser sesion
    	zipPath := fmt.Sprintf("%x.zip", phoneHashed.Sum((nil)))
    	log.Println("Creating user zip ...")
	utils.MyZip(zipPath, "./myUsers")
    	log.Println("User zip created")

    	mf, _ := os.Open(zipPath)
	state, _:= mf.Stat()
	data := make([]byte, state.Size())
	mf.Read(data)

    	//Send browser session to azure blob storage
    	log.Println("Sending user zip to azure...")
    	//storeInAzure(phoneNumber, data)
    	//log.Println(data)
    	log.Println("User zip sent to azure") 
}

func storeInAzure(blobName string, blobToSend []byte) (err error){
    	storage_account_url := os.Getenv("STORAGE_ACCOUNT_URL")
    	ctx := context.Background()

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	client, err := azblob.NewClient(storage_account_url, credential, nil)
	handleError(err)

    	log.Printf("Uploading a blob named %s\n", blobName)
	_, err = client.UploadBuffer(ctx, "raw", blobName, blobToSend, &azblob.UploadBufferOptions{})
	handleError(err)

    return nil
}