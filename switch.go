package main

import "fmt"

func main() {
	name := "Eko"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Kenalan dong")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama cukup panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
