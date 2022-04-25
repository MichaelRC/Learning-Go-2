package main

import "fmt"

func newMain() {
	fmt.Println("Hello World")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world."
	fmt.Println(whatToSay)

	i = 42

	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()

	fmt.Println(whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
