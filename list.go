package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Eko")
	data.PushBack("Kurniawan")
	data.PushBack("Khannedy")
	data.PushFront("Budi")

	//dari depan ke belakang
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	//dari belakang ke depan
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
