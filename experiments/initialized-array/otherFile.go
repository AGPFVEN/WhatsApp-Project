package main

import (
	"log"
	"strconv"
)

var matrix [3][2]string

func populateMatrix(){
	for i:= 0; i < 3; i++{
		matrix[i][0] = ""
	}
}

func populateMatrix1(){
	matrix[1][0] = "Hh"
}

func showMatrix(){
	for i:= 0; i < 3; i++{
		for i1:= 0; i1 < 2; i1++{
			log.Println("[" + strconv.Itoa(i) + "]" + "[" + strconv.Itoa(i1) + "]=" + strconv.FormatBool("" == matrix[i][i1]))
		}
	}
}