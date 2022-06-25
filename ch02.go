package main

import (
	"fmt"
)

const boilingF = 212.0

func main() {
	boilingPoint()
	ftocPrint()
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
