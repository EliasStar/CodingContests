package main

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

func level2(filename string, wg *sync.WaitGroup) {
	in, err := os.ReadFile(filename + ".in")
	if err != nil {
		panic(err)
	}

	input := strings.Fields(string(in))

	boardSize, _ := strconv.Atoi(input[0])
	board := make([][]string, boardSize)
	for i := 0; i < boardSize; i++ {
		board[i] = strings.Split(input[i+1], "")
	}

	row, _ := strconv.Atoi(input[boardSize+1])
	row--

	col, _ := strconv.Atoi(input[boardSize+2])
	col--

	path := strings.Split(input[boardSize+4], "")

	coins := 0
	for i := 0; i < len(path); i++ {
		switch path[i] {
		case "L":
			col--
		case "R":
			col++
		case "U":
			row--
		case "D":
			row++
		}

		if board[row][col] == "C" {
			coins++
			board[row][col] = "E"
		}
	}

	os.WriteFile(filename+".out", []byte(strconv.Itoa(coins)), os.ModePerm)
	wg.Done()
}
