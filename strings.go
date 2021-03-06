package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Eko Kurniawan", "Eko"))
	fmt.Println(strings.Split("Eko Kurniawan Khannedy", " "))
	fmt.Println(strings.ToLower("Eko Kurniawan Khannedy"))
	fmt.Println(strings.ToUpper("Eko Kurniawan Khannedy"))
	fmt.Println(strings.ToTitle("eko kurniawan khannedy"))
	fmt.Println(strings.Trim("                  eko kurniawan khannedy        ", " "))
	fmt.Println(strings.ReplaceAll("Eko Eko Eko Eko", "Eko", "Budi"))
}
