package main

import "strconv"

/*
 * Complete the 'superDigit' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING n
 *  2. INTEGER k
 */

func superDigitRecursive(n string, k int32) int32 {
	// Initialize sum
	var sum int32

	// Calculate sum (using ASCII code)
	for _, v := range n {
		sum += int32(v - '0')
	}

	// Multiply sum by k (which same to concatenate n k times)
	sum *= k

	// If sum is less than 10, return sum
	if sum < 10 {
		return sum
	}

	// Recursive call to superDigit (set k to 1, because we already concatenate n k times)
	return superDigit(strconv.Itoa(int(sum)), 1)
}
