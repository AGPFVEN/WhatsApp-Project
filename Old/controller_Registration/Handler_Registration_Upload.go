package controller_registration

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"

	"github.com/agpfven/WhatsApp_project/config"
	"github.com/agpfven/WhatsApp_project/utils"
)

func HandlerRegistrationUpload(phoneNumber string, isAllocatorClosed context.Context, isBrowserClosed context.Context, specific_local_user int) (){
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

	//Presending data to Azure
    	log.Println("User zip preparation to send ...")
    	mf, _ := os.Open(zipPath)
	state, _:= mf.Stat()
	data := make([]byte, state.Size())
	mf.Read(data)
    	log.Println("User zip preparation to send completed")

	//Encrypting data NOT TESTED .....................................................
	log.Println("Starting encryption process ...")
    	secret_encryption_key1 := os.Getenv("SECRET_ENCRYPTION_KEY1")
    	secret_encryption_key2 := os.Getenv("SECRET_ENCRYPTION_KEY2")
	encrypted_data, err := utils.MyEncryption(data, []byte(secret_encryption_key1 + phoneNumber + secret_encryption_key2 + string(phoneHashed.Sum(nil))))
	config.HandleError(err)
	log.Println("Encryption process finished")

    	//Send browser session to azure blob storage
    	log.Println("Sending user zip to azure...")
    	storeInAzure(string(phoneHashed.Sum(nil)), encrypted_data)
    	log.Println("User zip sent to azure") 
}

func storeInAzure(blobName string, blobToSend []byte) (err error){
    	storage_account_url := os.Getenv("STORAGE_ACCOUNT_URL")
    	ctx := context.Background()

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	config.HandleError(err)

	client, err := azblob.NewClient(storage_account_url, credential, nil)
	config.HandleError(err)

    	log.Printf("Uploading a blob named %s\n", blobName)
	_, err = client.UploadBuffer(ctx, "raw", blobName, blobToSend, &azblob.UploadBufferOptions{})
	config.HandleError(err)

    return nil
}