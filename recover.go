package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Terjadi error", message)
	}
	fmt.Println("aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Aplikasi ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Eko")
}
