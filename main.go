package main

import "mylearning/myutil"

func main() {
	println("Hello World")
	println("----------------------------")
	myutil.PrintMessage("Package Imported")
	myutil.PrintMyName("Zakir Hossain")

	println("variable type learing")

	var name string = "Zakir Hossain"
	var age = 28
	var height float64 = 5.6
	var isAdult bool = true
	const pi = 3.14
	person := "Zahid Hasan"

	var private = "This is private variable"
	var Public = "This is public variable"

	println(private)
	println(Public)
	println(person)
	println(pi)

	println("My Name is", name, ". I am", age, "Years old. ", "I am ", height, "tall. ", "I am ", isAdult)

}
