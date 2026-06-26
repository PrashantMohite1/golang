// pointer variable : in go we can store memory of address to other variable using pointer variable
// Inside your running Go program, a pointer means: “work on this value in memory, not a duplicate.”
// So when you pass &account or use (a *Account), you’re saying: don’t give me a copy — give me the real one (or its address).
// get actual value of variable in pointer : use fmt.println(*pointer-var)

// ----------------------------------

// pointers with functions

// pointer with structs
// two ways declare pointer struct variable
// access struct using pointer, update the value of struct using pointer

// we use pointer as argument and return nothing : to update pointer variable value
//  we use argument and return pointer var : to update a pointer and return it to same or another pointer

// ===================

/*
package main

func increment(n *int) *int {
	*n = 10
	return n
}

func main() {
	num := 2
	var ptr1 *int
	ptr1 = &num

	println("address of num : ", &num)
	println("value of num before update : ", num)

	*ptr1 = 5

	println("address of pointer : ", ptr1)
	println("value of num after pointer update : ", num)

	// passing and pointer to function
	increment(ptr1)

	println("value of num after increment func : ", num)

}

*/