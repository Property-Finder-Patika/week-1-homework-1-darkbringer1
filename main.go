package main

import (
	"fmt"
	"os"
)

func main() {
	var name string = "Zeynep"
	var greeting = createGreet(name)
	fmt.Printf("%s", greeting)
	greetCreatorForMain()
}

func greet(name string) {

}

func createGreet(name string) string {
	greeting := "Selam " + name + " :)"
	return greeting
}

func greetCreatorForMain() {
	name := os.Args[1]
	greeting := createGreet(name)
	fmt.Printf("%s\n", greeting)
}

// func createGreet(name string) string {
// 	return "Selam " + name + " :) "
// }
