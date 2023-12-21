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
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	nPlayers := len(ranked)
	nGames := len(player)

	playerRanks := make([]int32, nPlayers)

	playerRanks[0] = 1
	for i := 1; i < nPlayers; i++ {
		if ranked[i] == ranked[i-1] {
			playerRanks[i] = playerRanks[i-1]
			continue
		}
		if ranked[i] < ranked[i-1] {
			playerRanks[i] = playerRanks[i-1] + 1
			continue
		}
	}

	currentRanks := make([]int32, nGames)
	for i := 0; i < nGames; i++ {
		currentRanks[i] = 1
	}
	for j := 0; j < nGames; j++ {
		for i := 0; i < nPlayers; i++ {
			if player[j] < ranked[i] {
				currentRanks[j] = playerRanks[i] + 1
			}
		}
	}
	return currentRanks

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	rankedCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	rankedTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ranked []int32

	for i := 0; i < int(rankedCount); i++ {
		rankedItemTemp, err := strconv.ParseInt(rankedTemp[i], 10, 64)
		checkError(err)
		rankedItem := int32(rankedItemTemp)
		ranked = append(ranked, rankedItem)
	}

	playerCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	playerTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var player []int32

	for i := 0; i < int(playerCount); i++ {
		playerItemTemp, err := strconv.ParseInt(playerTemp[i], 10, 64)
		checkError(err)
		playerItem := int32(playerItemTemp)
		player = append(player, playerItem)
	}

	result := climbingLeaderboard(ranked, player)

	for i, resultItem := range result {
		// fmt.Fprintf(writer, "%d", resultItem)
		fmt.Println(resultItem)
		if i != len(result)-1 {
			// fmt.Fprintf(writer, "\n")
		}
	}

	// fmt.Fprintf(writer, "\n")

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
