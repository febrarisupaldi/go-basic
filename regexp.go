package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("eKo"))

	search := regex.FindAllString("eko eka edo eto emo", -1)
	fmt.Println(search)
}
