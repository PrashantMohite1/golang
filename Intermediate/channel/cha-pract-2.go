package main

import (
	"fmt"
	"time"
)

func generateLogs(logs chan string) {
	logs <- "keycloak is waiting for the postgres"
	logs <- "importing realm in kc"
	logs <- "used given realm in keycloak"
	logs <- "container is running on port 8000"
	// close(logs)
}

func printLogs(logs chan string) {
	for i := range logs {
		fmt.Println(i)
	}
}

func main() {

	logs := make(chan string)
	go generateLogs(logs)
	go printLogs(logs)

	time.Sleep(2 * time.Second)

}
