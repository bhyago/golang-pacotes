package main

import (
	"github.com/bhyago/golang-pacotes/packaging/go-mod-replace/math"
	"github.com/google/uuid"
)

func main() {
	m := math.NewMath(1, 3)
	println(m.Add())
	println(uuid.New().String())
}
