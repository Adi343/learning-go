package main
import 	"fmt"

func mapsTest(){
	m:= make(map[string]string)
	m["question"] = "This is a question"
	m["answer"] = "This is a answer"
	fmt.Println("Hello maps")
	fmt.Println("m is ",m)
	fmt.Println(m["answer"])
}
