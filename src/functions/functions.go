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

func greetPersonByName(name string) {
	fmt.Println("Hello", name)
}

func greetingMessage() string {
	return "greeting!!"
}

func addThree(x, y, z int) int {
	return x + y + z
}

func seventySeven() int {
	return 77
}

func fortyThreeSixtySeven() (int, int) {
	return 43, 67
}

func main() {

	greetPersonByName("Mohamed")

	fmt.Println(greetingMessage())

	fmt.Println(addThree(23, 54, 6))

	fmt.Println(seventySeven())

	fmt.Println(fortyThreeSixtySeven())

	y, z := fortyThreeSixtySeven()
	fmt.Println(addThree(seventySeven(), y, z))
}
