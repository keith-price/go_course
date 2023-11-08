//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func displayName(name string) {
	fmt.Println("Hello", name)
}

func anyMessage() string {
	return "Pooped mi pants again missus!"
}

func addThreeNums(num1 float32, num2 float32, num3 float32) float32 {
	return num1 + num2 + num3
}

func anyNum() float32 {
	return 10.8
}

func anyTwoNums() (float32, float32) {
	return 28.8, 87.4
}

func main() {
	displayName("Keith")

	fmt.Println(anyMessage())

	fmt.Println(addThreeNums(3.2, 9.1, 8.5))

	num1, num2 := anyTwoNums()
	fmt.Println(addThreeNums(anyNum(), num1, num2))

}
