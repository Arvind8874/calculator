package main

import "fmt"

func main() {

	fmt.Println("welcome to a class on poineters")
	// var pre *int
	// fmt.Println("value of pointer is", pre)
	myNumber := 12

	var ptr = &myNumber
	fmt.Println("value of a class on pointer", ptr)
	fmt.Println("value of a class on pointer", *ptr)

	*ptr = *ptr*2 + 3
	fmt.Println(myNumber)
}
