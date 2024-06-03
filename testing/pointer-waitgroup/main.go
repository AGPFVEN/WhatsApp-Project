package main

import (
	"log"
	"sync"
)

const m1dsize = 3

var matrix [m1dsize][2]string

func main(){ //main
	log.Println("Before function")
	f0()
}

//InitialPageQrMsg
func f0(){
	//c := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	
	var q string = "1"
	var p *string = &q

	go f1(p, wg)

	//log.Println("InitalPageQrMsg " + <-c)
	wg.Wait()
	log.Println("InitalPageQrMsg " + *p)

	//close(c)
}

//registrationDataHandler
//func f1 (ch chan string){
func f1 (p1 *string, wg1 *sync.WaitGroup){
	//f2(ch)
	f2(p1, wg1)
	log.Println("registrationDataHandler before handling channel")
	//matrix[0][0] = <-ch
	wg1.Wait()
	matrix[0][0] = *p1
	matrix[0][1] = "phone number"

	log.Println("registrationDataHandler " + matrix[0][0])
}

//getQrCode
//func f1 (ch chan string){
func f2 (p2 *string, wg2 *sync.WaitGroup){
	//ch1 <- "channel"
	*p2 = "qr data"
	wg2.Done()
}