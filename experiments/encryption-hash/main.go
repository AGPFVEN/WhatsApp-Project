package main

import (
	"crypto/sha256"
	"fmt"
	"log"
)

func main(){
	const inp = "Simple test"

	h := sha256.New()

	h.Write([]byte(inp))

	t := fmt.Sprintf("%x.zip", h.Sum((nil)))

	log.Printf(t)
}