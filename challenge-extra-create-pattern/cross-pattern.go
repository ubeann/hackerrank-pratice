package main

import "fmt"

// when number is 1
// *
// when number is 2
// *-*-*
// -*-*-
// *-*-*
// when number is 3
// *---*---*---*
// -*-*-*-*-*-*-
// --*---*---*--
// -*-*-*-*-*-*-
// *---*---*---*

func main() {
	// Create a variable to store the number
	var number int

	// Ask the user to enter a number
	fmt.Println("Enter a number: ")

	// Read the number from the console
	fmt.Scanf("%d", &number)

	// Calculate lines and characters
	lines := (number * 2) - 1
	characters := (lines * number) - (number - 1)

	// Create a loop to print the pattern
	for line := 1; line <= lines; line++ {
		// Create a loop to print the character
		for character := 1; character <= characters; character++ {
			if (character-line)%(lines-1) == 0 {
				// fmt.Print("     ")
				fmt.Print("*")
			} else {
				fmt.Print(" ")
				// fmt.Print("[",line, character, "]")
			}
		}

		// Print the new line
		fmt.Println()
	}
}
