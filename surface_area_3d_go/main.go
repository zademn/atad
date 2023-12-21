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
 * Complete the 'surfaceArea' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY A as parameter.
 */

func surfaceArea(A [][]int32) int32 {
	n_rows := len(A)
	n_cols := len(A[0])
	s := 0
	for i := 0; i < n_rows; i++ {
		for j := 0; j < n_cols; j++ {
			if A[i][j] != 0 {
				s += 2
			}
			if i == 0 {
				s += int(A[i][j])
			}
			if i == n_rows-1 {
				s += int(A[i][j])
			}
			if j == 0 {
				s += int(A[i][j])
			}
			if j == n_cols-1 {
				s += int(A[i][j])
			}

			if i > 0 {
				v := A[i][j] - A[i-1][j]
				if v > 0 {
					s += int(v)
				}
			}
			if i < n_rows-1 {
				v := A[i][j] - A[i+1][j]
				if v > 0 {
					s += int(v)
				}
			}
			if j > 0 {
				v := A[i][j] - A[i][j-1]
				if v > 0 {
					s += int(v)
				}
			}
			if j < n_cols-1 {
				v := A[i][j] - A[i][j+1]
				if v > 0 {
					s += int(v)
				}
			}
		}
	}
	// Write your code here
	return int32(s)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	HTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	H := int32(HTemp)

	WTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	W := int32(WTemp)

	var A [][]int32
	for i := 0; i < int(H); i++ {
		ARowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var ARow []int32
		for _, ARowItem := range ARowTemp {
			AItemTemp, err := strconv.ParseInt(ARowItem, 10, 64)
			checkError(err)
			AItem := int32(AItemTemp)
			ARow = append(ARow, AItem)
		}

		if len(ARow) != int(W) {
			panic("Bad input")
		}

		A = append(A, ARow)
	}

	result := surfaceArea(A)
	fmt.Println(result)
	// fmt.Fprintf(writer, "%d\n", result)

	// writer.Flush()
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
