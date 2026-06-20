// Change value of a map in Golang
// Add Element of Go Map Element
// Delete Element of Go Map Element

package main

import "fmt"

func main() {
	// mp1 = map[keytype]valuetype {}

	mp1 := map[int]string{1: "Prashant", 2: "Sanat", 3: "Rakesh"}
	fmt.Println(mp1)

	fmt.Println(mp1[1])

	submark := map[string]float32{"English": 90, "Science": 76}
	fmt.Println(submark["Science"]) // map access are case sensitive - if we change it science it will return zero , because we Capital S

	//  add element in map
	submark["maths"] = 55
	fmt.Println("updated map ", submark)

	// delete elements from map
	delete(submark, "Science")
	fmt.Println("map after delete opern : ", submark)

}
