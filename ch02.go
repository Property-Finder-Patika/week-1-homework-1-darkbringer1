package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
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
	flagMethod()
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

func flagMethod() {
	n := flag.Bool("n", false, "omit trailing newline")
	sep := flag.String("s", " ", "seperator")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}