package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", hostname)
	} else {
		fmt.Println("error", err.Error())
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}
