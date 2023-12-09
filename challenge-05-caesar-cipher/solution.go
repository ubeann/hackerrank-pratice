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
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER n (length of string)
 *  2. STRING s (string to encrypt)
 *  3. INTEGER k (number of rotations)
 */

func caesarCipher(lengthOfString int32, text string, rotations int32) string {
	// NOTE: The time complexity of this solution is O(n) where n is the length of the string

	// Convert string to rune slice
	runes := []rune(text)

	// Iterate over each rune
	for i, currentRune := range runes {
		// Check if rune is a letter
		if currentRune >= 'a' && currentRune <= 'z' {
			// Rotate rune by k
			runes[i] = 'a' + (currentRune-'a'+rotations)%26
		} else if currentRune >= 'A' && currentRune <= 'Z' {
			// Rotate rune by k
			runes[i] = 'A' + (currentRune-'A'+rotations)%26
		}
	}

	// Return string
	return string(runes)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(n, s, k)

	fmt.Fprintf(writer, "%s\n", result)

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
