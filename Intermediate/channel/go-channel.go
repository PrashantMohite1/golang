// channel - is use to make communication between go routines
//  channel is a memory address, which acts as a medium through which goroutines send and receive data to communicate.

// go channel - we use make () function

// 1. Send data to the channel

// channelName <- data

// Receive data from the channel
// <- channelName

/*

package main

import (
	"fmt"
	"time"
)

func sender(ch chan string) {
	fmt.Println("Sender: Sending A")
	ch <- "A" // function/operation is pause until this channel used by other go routine

	fmt.Println("Sender: Sending B")
	ch <- "B" // function/operation is pause until this channel used by other go routine

	fmt.Println("Sender: Sending C")
	ch <- "C"

	fmt.Println("Sender: Finished")
}

func receiver(ch chan string) {
	time.Sleep(2 * time.Second)

	fmt.Println("Receiver got:", <-ch)

	time.Sleep(2 * time.Second)

	fmt.Println("Receiver got:", <-ch)

	time.Sleep(2 * time.Second)

	fmt.Println("Receiver got:", <-ch)
}

func main() {
	ch := make(chan string)

	go sender(ch)
	go receiver(ch)

	time.Sleep(10 * time.Second)
}


*/