package main

import (
	"fmt"
)

var std_a = 90

// std_b := 70
// std_c := 55

var num = std_a

func main() {

	if num >= 90 {
		fmt.Println("student grade : A")
	}

	// count 1 to 10
	fmt.Println("Printing 1 to 10 count")

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	addition_of_100 := 0
	// print even numbers
	fmt.Println("Print Sum")

	for i := 0; i <= 100; i++ {
		addition_of_100 = addition_of_100 + i
		// final_total := num +
	}
	// total := final_total
	fmt.Println("addition of 100 is ", addition_of_100)

	// total := 2 % 2
	// fmt.Println(total)

}
