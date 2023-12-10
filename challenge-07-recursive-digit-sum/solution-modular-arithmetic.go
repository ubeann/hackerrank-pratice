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
 * Complete the 'superDigit' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING n
 *  2. INTEGER k
 */

func superDigit(n string, k int32) int32 {
	// NOTE: This solution is based on the modular arithmetic property: (a + b) mod n = ((a mod n) + (b mod n)) mod n
	//       For example:
	//       (123 + 456) mod 9 = ((1 + 2 + 3) + (4 + 5 + 6)) mod 9
	//             (579) mod 9 = ((6) + (15)) mod 9
	//                       3 = 3
	//       We can use recursive to calculate the digit sum, but it will cause stack overflow for large input
	//       So, we use iterative instead. The idea is to calculate the digit sum using modular arithmetic
	//       The time complexity is O(n), where n is the length of string n

	// Calculate the digit sum using modular arithmetic
	digitSum := int32(0)
	for _, v := range n {
		digitSum += int32(v - '0')
		digitSum %= 9
	}

	// Calculate the super digit using the modular arithmetic property
	superDigit := (digitSum * k) % 9

	// If the super digit is 0, return 9 (special case: when 9 mod 9 = 0, but the super digit is 9)
	if superDigit == 0 {
		return 9
	}

	// Return the super digit
	return superDigit
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	n := firstMultipleInput[0]

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := superDigit(n, k)

	fmt.Fprintf(writer, "%d\n", result)

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
