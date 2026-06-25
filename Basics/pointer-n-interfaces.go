// pointer variable : in go we can store memory of address to other variable using pointer variable
// Inside your running Go program, a pointer means: “work on this value in memory, not a duplicate.”
// So when you pass &account or use (a *Account), you’re saying: don’t give me a copy — give me the real one (or its address).
// get actual value of variable in pointer : use fmt.println(*pointer-var)

package main

func main() {
	num := 10
	var ptr1 *int
	ptr1 = &num

	println("address of num : ", ptr1)
	println("value of pointer before update : ", *ptr1)
}
