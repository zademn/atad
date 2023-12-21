package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

/*
 * Complete the 'encryption' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func encryption(s string) string {
	w := strings.ReplaceAll(s, " ", "")
	l := len(s)
	t := int(math.Floor(math.Sqrt(float64(l))))
	row_opt := [4]int{t, t, t + 1, t + 1}
	col_opt := [4]int{t, t + 1, t, t + 1}
	n_rows := 0
	n_cols := 0

	for i := 0; i < 4; i++ {
		if row_opt[i]*col_opt[i] >= l {
			n_rows = row_opt[i]
			n_cols = col_opt[i]
			break
		}
	}

	v := make([][]int, n_rows)
	for i := 0; i < n_rows; i++ {
		v[i] = make([]int, n_cols)
		for j := i * n_cols; j < (i+1)*n_cols; j++ {
			if j < l {
				v[i][j-i*n_cols] = int(w[j])
			}
		}
	}

	var result string
	for i := 0; i < n_cols; i++ {
		for j := 0; j < n_rows; j++ {
			if v[j][i] != 0 {
				result += string(v[j][i])
			}
		}
		result += " "
	}
	fmt.Println(result)
	return result

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	s := readLine(reader)

	result := encryption(s)
	fmt.Println(result)
	// fmt.Fprintf(writer, "%s\n", result)

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
