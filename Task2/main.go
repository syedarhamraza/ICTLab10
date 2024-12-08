package main

import "fmt"

func main() {
	var number int
	var square int
	fmt.Println("Enter a number: ")
	fmt.Scan(&number)
	if number%2 == 0 {
		square = number * number
		fmt.Println("Number is even so the Square is: ", square)
	} else {
		cube := number * number * number
		fmt.Println("Number is odd so the cube is: ", cube)
	}
}
