package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := Pembagian(10, 2)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("error", err.Error())
	}
}
