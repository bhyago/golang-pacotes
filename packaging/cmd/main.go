package main

import (
	"fmt"

	"github.com/bhyago/golang-pacotes/packaging/math"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
	fmt.Println("Hello, World!")
}
