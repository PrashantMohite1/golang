package main

import "fmt"
import "example/first-module/greeting"


func main(){

	names := []string{"sam", "john", "jane", "doe"}
	// names := "sam"

	message, err := greeting.MultiplePrints(names)


	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(message)
}


