/*  ############# CPU , RAM , THREAD ############################################

The CPU (The Chefs): The CPU actually does the work. If you have an 8-core CPU, you have 8 chefs standing at the counter.
A chef can only chop vegetables, cook, or read a recipe when they are actively working. CPUs do not store data; they process instructions.

The RAM (The Kitchen Counter): RAM is the workspace. Before a chef (CPU) can execute a program, that program’s code, variables, and tracking state must be loaded onto the kitchen counter (RAM).
If the counter fills up, the kitchen grinds to a halt.



// Threads : threads are logical subparts of a CPU core.
// threads are logical subparts of a CPU core.  While a core is a physical processing unit, a thread is a virtual execution path that allows that single physical core to manage multiple instruction streams concurrently
// we can check thread in windows from processor we are using - in spec details will get cpu cores and threads details


How Docker is using routing : - like we know to get routing execution our main should be continuously running
in docker main program is running in daemon baground mode - which does his job and when some ask any other functionality/task it uses go routine to do that in fasted ways without interrupting main task right



*/

// go routines

// normally waits until first statement complete then it moves to second but in go routine -
// if lets say we have gave go routine at first statemnt then second will executing immidiately without waiting for go routing execution

// Run Two Function concurrently using routing
// what happens before go routine execution our main will end so we need to add sleep in go routine
// so that main will not end and we can see then go routine output

// terms :
// concurrently : means happening/running at same time
// synchronous - means running at same time

// Concurrent programming is a technique where multiple tasks or processes execute in overlapping time periods

// Multiple goroutine : order of executing multiple go routine is random in go

// The CPU stops executing one task and starts executing another, while remembering where the first one stopped.

// ------------------------------

/*
package main

import (
	"fmt"
	"time"
)

func greet() {
	fmt.Println("Hey! what's up.")
}

func main() {

	go fmt.Println("Process : 1")
	go fmt.Println("Process : 2")
	go fmt.Println("Process : 3")

	time.Sleep(2)
}

*/