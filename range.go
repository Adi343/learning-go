package main

import "fmt"

func rangeTest(){
	coins := [4]int{1,9,4}
	var sum int = 0

	for _,num := range coins {
		sum += num
	}

	fmt.Println(coins)
	fmt.Println(sum)

}