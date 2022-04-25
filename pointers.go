package main

import "log"

func main() {
	myString := "Green"

	log.Println("myString is set to", myString)
	changeUSingPointer(&myString)
	log.Panicln("after func call myString is set to", myString)
}

func changeUSingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
