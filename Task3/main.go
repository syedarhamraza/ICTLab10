// SYED ARHAM RAZA 2430-0159
package main

import "fmt"

func main() {
	var sliceOfNumbers = []int{}
	var numbers int
	for i := 0; i < 5; i++ {
		fmt.Print("Enter some Numbers: ")
		fmt.Scan(&numbers)
		sliceOfNumbers = append(sliceOfNumbers, numbers)
	}
	var sum int = 0
	for i := 0; i < len(sliceOfNumbers); i++ {
		sum += sliceOfNumbers[i]
	}
	avg := float64(sum) / float64(len(sliceOfNumbers))
	fmt.Println("The sum of slice is:", sum)
	fmt.Println("The avg of slice is:", avg)

}
