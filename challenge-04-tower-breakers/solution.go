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
 * Complete the 'towerBreakers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER m
 */

func towerBreakers(n int32, m int32) int32 {
	// NOTES: Time complexity is O(1) because the function is a simple calculation.

	// EXPLANATION:

	// - If the height of each tower is 1, Player 2 wins. Why?
	//   Because Player 1 can only reduce the height to a divisor of 1, which is 1 itself.
	//   Player 2 then mirrors the move, reducing the height to 1 again. This cycle continues,
	//   and Player 1 will eventually run out of moves, leaving Player 2 with the last move
	//   to remove the last tower. Player 2 wins.

	// - If the number of towers is even, Player 2 wins. Why?
	//   If the number of towers is even, Player 1 can only remove towers in pairs to maintain
	//   an even number of towers. Player 2 can then mirror each move, ensuring that the number
	//   of towers remains even. This process continues until Player 1 has no valid moves left,
	//   and Player 2 wins by removing the last tower.

	// - If the number of towers is odd, Player 1 wins. Why?
	//   If the number of towers is odd, Player 1 can consistently remove one tower at a time,
	//   maintaining an odd number of towers. Player 2, forced to mirror the move, will eventually
	//   run out of moves as the number of towers becomes 1. Player 1 wins by removing the last tower.

	// If the height of each tower is 1 or the number of towers is even, Player 2 wins.
	if m == 1 || n%2 == 0 {
		return 2
	}

	// If the number of towers is odd, Player 1 wins.
	return 1
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
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		result := towerBreakers(n, m)

		fmt.Fprintf(writer, "%d\n", result)
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
