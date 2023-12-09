package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	// Get number of test cases
	var t int
	fmt.Scanf("%d", &t)

	// Iterate over test cases
	for i := 0; i < t; i++ {
		// Get number of elements in sequence
		var n int
		fmt.Scanf("%d", &n)

		// Get sequence
		var sequence = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &sequence[j])
		}

		// Run zig-zag sequence
		zigZagSequence(sequence, n)
	}
}

func zigZagSequence(sequence []int, n int) {
	// NOTE: This solution is based on the solution to the Longest Increasing Subsequence problem
	// The time complexity of this solution is O(n^2) because we are looping through the array twice

	// If sequence is empty or has only one element, return sequence
	if n <= 1 {
		fmt.Println(sequence)
	}

	// Sort sequence in ascending order
	result := sort(sequence, n)

	// Find middle index
	middle := n / 2
	result[middle], result[n-1] = result[n-1], result[middle]

	// Find left and right indices
	left := middle + 1
	right := n - 2

	// Iterate over sequence and swap elements
	for left <= right {
		// Swap left and right elements
		result[left], result[right] = result[right], result[left]

		// Increment left and right indices
		left++
		right--
	}

	// Iterate over sequence and print elements
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Printf("%d\n", result[i])
		} else {
			fmt.Printf("%d ", result[i])
		}
	}
}

func sort(sequence []int, n int) []int {
	// Iterate over sequence
	for i := 1; i < n; i++ {
		// Get current element
		current := sequence[i]

		// Get previous element
		j := i - 1

		// Shift elements of sequence that are greater than current to the right
		for j >= 0 && sequence[j] > current {
			sequence[j+1] = sequence[j]
			j--
		}

		// Insert current element into sequence
		sequence[j+1] = current
	}

	// Return sorted sequence
	return sequence
}
