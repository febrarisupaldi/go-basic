package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	slice := []string{"Eko", "Kurniawan", "Khannedy", "Budi", "Joko"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for _, value := range slice {
		//fmt.Println("index", i, "=", value)
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Eko"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
