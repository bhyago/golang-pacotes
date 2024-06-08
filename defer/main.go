package main

import (
	"fmt"
)

func main() {
	fmt.Println("primeira linha")
	defer fmt.Println("segunda linha")
	defer fmt.Println("terceira linha")
	fmt.Println("quarta linha")

}
