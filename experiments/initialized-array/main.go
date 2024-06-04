package main

import "log"

/*
	Conclusion of the experiment:
initialize strings are assigned ""
I can trust the init function (execute before main)
*/

func init(){
	showMatrix()
	populateMatrix()
	showMatrix()
}

func main(){
	log.Println("jj")
	populateMatrix1()
	showMatrix()
}

