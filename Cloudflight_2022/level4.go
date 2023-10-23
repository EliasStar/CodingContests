package main

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

type Field4 struct {
	up *Field4
	down *Field4
	left *Field4
	right *Field4
	visited bool
}

func level4(filename string, wg *sync.WaitGroup) {
	in, err := os.ReadFile(filename + ".in")
	if err != nil {
		panic(err)
	}

	input := strings.Fields(string(in))

	boardSize, _ := strconv.Atoi(input[0])
	board := make([][]*Field4, boardSize)
	for i := 0; i < boardSize; i++ {
		board[i] = make([]*Field4, boardSize)

		raw := strings.Split(input[i+1], "")
		for j := 0; j < boardSize; j++ {
			if raw[j] == "P" || raw[j] == "C" {
				board[i][j] = &Field4{}
			}
		}
	}

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == nil { continue }

			if i > 0 {
				board[i][j].up = board[i-1][j]
			}

			if j > 0 {
				board[i][j].left = board[i][j-1]
			}

			if i < boardSize-1 {
				board[i][j].down = board[i+1][j]
			}

			if j < boardSize-1 {
				board[i][j].right = board[i][j+1]
			}
		}
	}

	row, _ := strconv.Atoi(input[boardSize+1])
	row--

	col, _ := strconv.Atoi(input[boardSize+2])
	col--

	path := visitChildFields4(board[row][col])

	os.WriteFile(filename+".out", []byte(path), os.ModePerm)
	wg.Done()
}

func visitChildFields4(field *Field4) (move string) {
	field.visited = true

	if field.up != nil && !field.up.visited {
		move += "U" + visitChildFields4(field.up) + "D"
	}

	if field.down != nil && !field.down.visited {
		move += "D" + visitChildFields4(field.down) + "U"
	}

	if field.left != nil && !field.left.visited {
		move += "L" + visitChildFields4(field.left) + "R"
	}

	if field.right != nil && !field.right.visited {
		move += "R" + visitChildFields4(field.right) + "L"
	}

	return
}
