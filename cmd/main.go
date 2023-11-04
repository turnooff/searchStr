package main

import (
	"fmt"

	searchstring "github.com/turnooff/searchStr/pkg/searchString"
)

func main() {
	fmt.Println(searchstring.Contains("test.txt", "qwerty"))
}
