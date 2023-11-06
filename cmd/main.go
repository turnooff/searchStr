package main

import (
	"fmt"

	searchstring "github.com/turnooff/searchStr/pkg/searchString"
)

func main() {
	flag, err := searchstring.Contains("test.txt", "qwerty")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(flag)
}
