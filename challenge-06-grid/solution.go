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
 * Complete the 'gridChallenge' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING_ARRAY grid as parameter.
 */

func gridChallenge(grid []string) string {
	// NOTE: The time complexity of this solution is O(n^2) where n is the length of the grid
	//       It is possible to reduce the time complexity to O(n) by using counting sort instead of bubble sort
	//       However, used bubble sort for simplicity. It is also possible to use quick sort or merge sort

	// Loop through each row
	for i := 0; i < len(grid); i++ {
		// Convert string to rune slice
		runes := []rune(grid[i])

		// Sort rune slice (using bubble sort)
		for j := 0; j < len(runes)-1; j++ {
			for k := 0; k < len(runes)-j-1; k++ {
				if runes[k] > runes[k+1] {
					runes[k], runes[k+1] = runes[k+1], runes[k]
				}
			}
		}

		// Convert rune slice back to string
		grid[i] = string(runes)
	}

	// Loop through each column
	for i := 0; i < len(grid[0]); i++ {
		// Loop through each row
		for j := 0; j < len(grid)-1; j++ {
			// Check if current character is greater than the next character
			if grid[j][i] > grid[j+1][i] {
				// Return "NO"
				return "NO"
			}
		}
	}

	// Return "YES"
	return "YES"
}

func main() {
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
		n := int32(nTemp)

		var grid []string

		for i := 0; i < int(n); i++ {
			gridItem := readLine(reader)
			grid = append(grid, gridItem)
		}

		result := gridChallenge(grid)

		fmt.Fprintf(writer, "%s\n", result)
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
