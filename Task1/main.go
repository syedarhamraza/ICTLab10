// SYED ARHAM RAZA 2430-0159
package main

import "fmt"

func main() {
	const PI float64 = 3.14159
	var radius float64
	var area float64
	fmt.Print("Enter the value of radius: ")
	fmt.Scan(&radius)
	area = PI * (radius * radius)
	fmt.Print("The value of area is: ", area)
}
