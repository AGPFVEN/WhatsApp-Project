package main

import (
	"bytes"
	"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func main(){
	//mf, _ := os.Open("x.zip")
	//state, _:= mf.Stat()
	//data := make([]byte, state.Size())
	//mf.Read(data)

	//storeInAzure("experimenting", data)

	retriveFromAzure("experimenting")
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
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

func retriveFromAzure(blobName string) (err error){
	/* refs to do this (done in first try lol)
		https://learn.microsoft.com/en-us/azure/storage/blobs/storage-blob-download-go
		https://stackoverflow.com/questions/29237411/how-to-convert-type-bytes-buffer-to-use-as-byte-in-argument-to-w-write
		https://stackoverflow.com/questions/32687985/convert-back-byte-array-into-file-using-golang
	*/
    	storage_account_url := os.Getenv("STORAGE_ACCOUNT_URL")
    	ctx := context.Background()

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	client, err := azblob.NewClient(storage_account_url, credential, nil)
	handleError(err)

    	log.Printf("Uploading a blob named %s\n", blobName)
	get, err := client.DownloadStream(ctx, "raw", blobName, nil)
	handleError(err)

	downloadedData := bytes.Buffer{}
	retryReader := get.NewRetryReader(ctx, &azblob.RetryReaderOptions{})
	_, err = downloadedData.ReadFrom(retryReader)
	handleError(err)

	err = retryReader.Close()
	handleError(err)

	err = os.WriteFile("THE_exp.zip", downloadedData.Bytes(), 0644)
	handleError(err)

    return nil
}