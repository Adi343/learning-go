package main

import (
	"fmt"
	"os"
)

func fileTest(){

	file,err := os.Open("./maps.go")
	defer file.Close()

	if(err!=nil){
		fmt.Println(err)
	}
	fmt.Println(file)
	

}