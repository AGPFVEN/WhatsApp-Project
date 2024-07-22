package utils

import (
	"crypto/aes"
	"log"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func MyEncryption(blobToSend []byte, secret []byte) (result []byte, err error){
	print("Encrypting data ...")

	//Encrypt
	c, err := aes.NewCipher(blobToSend)
	handleError(err)
	
	encrypted_blob := make([]byte, len(blobToSend))
	
	c.Encrypt(encrypted_blob, secret)
	
	print("Encryption finished")
	return encrypted_blob, nil
}