// function with dynamic parameters

// Note - The return statement should be the last statement of the function. When a return statement is encountered, the function terminates.
// We get a "missing return at the end of function" error as we have added a line after the return statement.

// function with two return statement

// variable scope - local & global

// recurisve function - we are calling inside its body itself
// Recursion with stopping condition ( using if else )

// calculate sum of positive numbers.

// ------------

// anonymous function with parameters.
// Anonymous Function as Arguments to Other Functions
// Return an Anonymous Function in Go
// nested function

// go closure : is a nested function that allows us to access variables of the outer function even after the outer function is closed.
// create a function which will increment numbers by go closure

// ================
/*

package main

import (
	"fmt"
)

// global variable
var global_num int

// Dynamic Functions.
func sum(num1 int, num2 int) int {
	output := num1 + num2 // local variable.
	return output

	// fmt.Println("adding line after return statement")      // this will give the error because we have added print statement after return statement. because function ends after return statement
}

// recursive function
// sum of positive numbers
func recsum(num1 int) int {
	if num1 == 0 {
		return 0
	} else {
		return num1 + recsum(num1-1)
	}
}

func factorial_fun(n1 int) int {
	if n1 == 0 {
		return 0
	} else {
		return n1 * (n1 - 1)
	}
}

func main() {

	// out := sum(9, 11)
	// fmt.Println(out)
	// fmt.Println(global_num)

	// out := recsum(6)
	// fmt.Println("output : ", out)

	// factorial of numbers
	out := factorial_fun(5)
	fmt.Println(out)

}

*/