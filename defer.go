package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Application running...")
	result := 10 / value
	fmt.Println("result", result)
}

func main() {
	runApplication(0)
}
