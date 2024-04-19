package main

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

type input3 struct {
	lawnCount int
	lawns     []lawn3
}

type lawn3 struct {
	width  int
	height int
	lawn   [][]byte
	path   string
}

func level3(filename string, wg *sync.WaitGroup) {
	in, err := os.ReadFile(filename + ".in")
	if err != nil {
		panic(err)
	}

	input := parseInput3(string(in))

	result := ""
	for _, lawn := range input.lawns {
		if isValidLawn(lawn) {
			result += "VALID\n"
		} else {
			result += "INVALID\n"
		}
	}

	os.WriteFile(filename+".out", []byte(result), os.ModePerm)
	wg.Done()
}

func parseInput3(input string) input3 {
	fields := strings.Fields(input)

	lawnCount, _ := strconv.Atoi(fields[0])

	lawns := []lawn3{}
	for i := 1; i < len(fields); i++ {
		width, _ := strconv.Atoi(fields[i])
		height, _ := strconv.Atoi(fields[i+1])
		i += 2

		lawn := [][]byte{}
		for _, line := range fields[i : i+height] {
			lawn = append(lawn, []byte(line))
		}
		i += height

		lawns = append(lawns, lawn3{width, height, lawn, fields[i]})
	}

	return input3{
		lawnCount: lawnCount,
		lawns:     lawns,
	}
}

type direction byte

const (
	north direction = 'W'
	east  direction = 'D'
	south direction = 'S'
	west  direction = 'A'
)

type position struct {
	x int
	y int
}

func isValidLawn(lawn lawn3) bool {

	for i := 0; i < lawn.height; i++ {
		for j := 0; j < lawn.width; j++ {
			if lawn.lawn[i][j] == '.' {
				if checkPath(lawn, position{j, i}, make(map[position]bool)) {
					return true
				}
			}
		}
	}

	return false
}

func checkPath(lawn lawn3, current position, visited map[position]bool) bool {
	visited[current] = true

	for _, move := range lawn.path {
		newX, newY := movePosition(current.x, current.y, direction(move))

		// Check if new position is valid
		if !isValid(newX, newY, lawn.height, lawn.width, lawn.lawn) || visited[position{newX, newY}] {
			return false
		}

		current.x = newX
		current.y = newY
		visited[position{newX, newY}] = true
	}

	// Check if all free cells are visited
	return len(visited) == (lawn.height*lawn.width - 1)
}

func movePosition(x, y int, dir direction) (int, int) {
	switch dir {
	case north:
		return x, y - 1
	case east:
		return x + 1, y
	case south:
		return x, y + 1
	case west:
		return x - 1, y
	default:
		panic("invalid direction")
	}
}

func isValid(x, y int, rows, cols int, lawn [][]byte) bool {
	return x >= 0 && x < cols && y >= 0 && y < rows && lawn[y][x] != 'X'
}
