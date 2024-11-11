package main

import (
	"fmt"
)

func main() {
	age := 22
	name := "John Doe"
	height := 6.75

	println("name: ", name, "height: ", height, "age: ", age)

	fmt.Printf("%s\n", name)
	fmt.Printf("Age is %d\n", age)
	fmt.Printf("Height is %.2f\n", height)
	fmt.Printf("Type of height is %T\n", height)
	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)

}
