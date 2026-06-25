// packages

// create a custom packages :
// create a alias of custom package

package main

import (
	"example/first-module/Basics/package-pract/calculator"
	"fmt"
)

func main() {

	addition := calculator.Add(5, 6)
	substraction := calculator.Sub(7, 2)

	fmt.Printf("addition = %d \n substraction = %d", addition, substraction)

}
