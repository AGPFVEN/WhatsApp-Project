package main

import (
	"crypto/sha256"
	"log"
)

func main(){
	const inp = "Simple test"

	h := sha256.New()

	h.Write([]byte(inp))

	log.Printf("%x\n", h.Sum((nil)))
}