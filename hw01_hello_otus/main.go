package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	greeting := "Hello, OTUS!"
	reversed := stringutil.Reverse(greeting)
	fmt.Println(reversed)
}
