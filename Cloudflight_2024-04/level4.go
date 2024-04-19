package main

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

type input4 struct {
	lawnCount int
	lawns     []lawn4
}

type lawn4 struct {
	width   int
	height  int
	lawn    [][]byte
	visited [][]bool
}

func level4(filename string, wg *sync.WaitGroup) {
	in, err := os.ReadFile(filename + ".in")
	if err != nil {
		panic(err)
	}

	input := parseInput4(string(in))

	result := ""
	for _, lawn := range input.lawns {
		path, _ := bfs(lawn.lawn, 0, 0)
		result += path + "\n"
	}

	os.WriteFile(filename+".out", []byte(result), os.ModePerm)
	wg.Done()
}

func parseInput4(input string) input4 {
	fields := strings.Fields(input)

	lawnCount, _ := strconv.Atoi(fields[0])

	lawns := []lawn4{}
	for i := 1; i < len(fields); {
		width, _ := strconv.Atoi(fields[i])
		height, _ := strconv.Atoi(fields[i+1])
		i += 2

		lawn := [][]byte{}
		for _, line := range fields[i : i+height] {
			lawn = append(lawn, []byte(line))
		}
		i += height

		visited := make([][]bool, height)
		for i := range visited {
			visited[i] = make([]bool, width)
		}

		lawns = append(lawns, lawn4{width, height, lawn, visited})
	}

	return input4{
		lawnCount: lawnCount,
		lawns:     lawns,
	}
}

type point struct {
	x int
	y int
}

func isValid4(grid [][]byte, x int, y int) bool {
	return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid) && grid[y][x] != 'X'
}

func bfs(grid [][]byte, startX int, startY int) (string, bool) {
	visited := make(map[point]bool)
	queue := []point{{startX, startY}}
	path := ""

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if !visited[current] {
			visited[current] = true

			for _, dir := range []direction{north, east, south, west} {
				nextX, nextY := current.x, current.y
				switch dir {
				case north:
					nextY--
				case east:
					nextX++
				case south:
					nextY++
				case west:
					nextX--
				}
				if isValid4(grid, nextX, nextY) {
					queue = append(queue, point{nextX, nextY})
				}
			}
		}

	}

	// Check if all free cells are visited
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '.' && !visited[point{x, y}] {
				return "", false
			}
		}
	}

	return path, true
}
