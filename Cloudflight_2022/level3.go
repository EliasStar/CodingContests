package main

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

type Ghost3 struct {
	row, col int
	path     []string
}

func level3(filename string, wg *sync.WaitGroup) {
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

	ghostCount, _ := strconv.Atoi(input[boardSize+5])
	ghosts := make([]Ghost3, ghostCount)
	for i := 0; i < ghostCount; i++ {
		r, _ := strconv.Atoi(input[boardSize+6+4*i])
		c, _ := strconv.Atoi(input[boardSize+7+4*i])

		ghosts[i] = Ghost3{
			row:  r - 1,
			col:  c - 1,
			path: strings.Split(input[boardSize+9+4*i], ""),
		}
	}

	coins := 0
	survived := "YES"

done:
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

		for j := 0; j < ghostCount; j++ {
			switch ghosts[j].path[i] {
			case "L":
				ghosts[j].col--
			case "R":
				ghosts[j].col++
			case "U":
				ghosts[j].row--
			case "D":
				ghosts[j].row++
			}

			if row == ghosts[j].row && col == ghosts[j].col {
				survived = "NO"
				break done
			}
		}

		if board[row][col] == "C" {
			coins++
			board[row][col] = "E"
		}

		if board[row][col] == "W" {
			survived = "NO"
			break
		}
	}

	os.WriteFile(filename+".out", []byte(strconv.Itoa(coins)+" "+survived), os.ModePerm)
	wg.Done()
}
