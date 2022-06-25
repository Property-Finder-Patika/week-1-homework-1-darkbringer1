package main

import (
	"fmt"
)

func main() {
	var name string = "Zeynep"
	var greeting = createGreet(name)
	fmt.Printf("%s", greeting)
}

func greet(name string) {

}

func createGreet(name string) string {
	greeting := "Selam " + name + " :)"
	return greeting
}

// func createGreet(name string) string {
// 	return "Selam " + name + " :) "
// }
