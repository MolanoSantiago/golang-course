package main

import (
	"fmt"
)

func main() {

	fmt.Println("CONSTS")
	const name string = "Santiago"
	const lastName = "Molano"
	fmt.Println(name)
	fmt.Println(lastName)
	fmt.Println(" ")

	fmt.Println("VARIABLES")
	base := 12          // First use
	var altura int = 14 // Another use
	var area int        // Go not run if a variable is unused
	fmt.Println(base, altura, area)
	fmt.Println(" ")

	fmt.Println("DEFAULT VALUES")
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)
	fmt.Println(" ")

	fmt.Println("EXERCISE")
	fmt.Println("Calculate area of a square with base 15cm")
	const baseSquare = 15
	var areaSquare = baseSquare * baseSquare
	fmt.Println("Square area: ", areaSquare)
	fmt.Println(" ")

	x := 18
	y := 2

	fmt.Println("ADDITION")
	result := x + y
	fmt.Println("Result: ", result)
	fmt.Println(" ")

	fmt.Println("SUBTRACT")
	result = x - y
	fmt.Println("Result: ", result)
	fmt.Println(" ")

	fmt.Println("MULTIPLICATION")
	result = x * y
	fmt.Println("Result: ", result)
	fmt.Println(" ")

	fmt.Println("DIVISION")
	result = x / y
	fmt.Println("Result: ", result)
	fmt.Println(" ")

	fmt.Println("RESIDUE")
	result = x % y
	fmt.Println("Result: ", result)
	fmt.Println(" ")

	fmt.Println("INCREMENT")
	x++
	fmt.Println("Result X: ", x)
	fmt.Println(" ")

	fmt.Println("DECREMENT")
	x--
	fmt.Println("Result X: ", x)
	fmt.Println(" ")

	fmt.Println("EXERCISE")
	fmt.Println("Calculate area of a rectangle with width 15cm and height 6cm")
	const widthRectangle = 15
	const heightRectangle = 6
	var areaRectangle = widthRectangle * heightRectangle
	fmt.Println("Rectangle area: ", areaRectangle)
	fmt.Println(" ")
}
