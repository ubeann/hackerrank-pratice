package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'zigZagSequence' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY sequence
 *  2. INTEGER n (length of sequence)
 */
func zigZagSequence(sequence []int, n int) []int {
	// NOTE: This is an implementation of the zig-zag sequence algorithm, which has a time complexity of O(n)
	// Therefore, the overall time complexity of this solution is O(n^2) due to the insertion sort algorithm.

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

	// Return zig-zag sequence
	return result
}

func sort(sequence []int, n int) []int {
	// NOTE: This is an implementation of insertion sort, which has a time complexity of O(n^2)

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

func main() {
	// Enter your code here. Read input from STDIN. Print output to STDOUT

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int(nTemp)

		sequenceTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var sequence []int

		for i := 0; i < int(n); i++ {
			sequenceItemTemp, err := strconv.ParseInt(sequenceTemp[i], 10, 64)
			checkError(err)
			sequenceItem := int(sequenceItemTemp)
			sequence = append(sequence, sequenceItem)
		}

		result := zigZagSequence(sequence, n)

		for i, resultItem := range result {
			fmt.Fprintf(writer, "%d", resultItem)

			if i != len(result)-1 {
				fmt.Fprintf(writer, " ")
			} else {
				fmt.Fprintf(writer, "\n")
			}
		}
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
