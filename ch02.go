package main

import (
	"fmt"
	"math/rand"
)

const boilingF = 212.0

func main() {
	boilingPoint()
	ftocPrint()
	gif()
}

func boilingPoint() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %g C\n", f, c)
}

func ftocPrint() {
	const freezingF, boilingF = 32.0, 212.0

	fmt.Printf("%gF = %gC\n", freezingF, fToC(freezingF))
	fmt.Printf("%gF = %gC\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func gif() {
	freq := rand.Float64() * 3.0
	t := 0.0

	fmt.Println(freq, t)
}