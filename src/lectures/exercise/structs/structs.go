//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	length int
	width  int
}

func calcRectArea(rectangle Rectangle) int {
	return rectangle.length * rectangle.width
}

func calcRectPerim(rectangle Rectangle) int {
	return (rectangle.length * 2) + (rectangle.width * 2)
}

func printRectInfo(rectangle Rectangle) {
	fmt.Println("The area of the given rectangle is:", calcRectArea(rectangle))
	fmt.Println("The perimeter of the given rectangle is:", calcRectPerim(rectangle))
}

func main() {
	var rect = Rectangle{257, 145}
	printRectInfo(rect)

}
