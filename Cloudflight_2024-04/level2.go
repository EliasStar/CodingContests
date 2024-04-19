package main

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

type input2 struct {
	pathCount int
	paths     []string
}

func level2(filename string, wg *sync.WaitGroup) {
	in, err := os.ReadFile(filename + ".in")
	if err != nil {
		panic(err)
	}

	input := parseInput2(string(in))

	result := ""
	for _, path := range input.paths {
		width, height := minLawnSize(path)
		result += strconv.Itoa(width) + " " + strconv.Itoa(height) + "\n"
	}

	os.WriteFile(filename+".out", []byte(result), os.ModePerm)
	wg.Done()
}

func parseInput2(input string) input2 {
	fields := strings.Fields(input)

	pathCount, _ := strconv.Atoi(fields[0])

	return input2{
		pathCount: pathCount,
		paths:     fields[1 : pathCount+1],
	}
}

func minLawnSize(path string) (int, int) {
	// Track changes in X and Y coordinates
	posX, posY := 0, 0
	minX, maxX, minY, maxY := 0, 0, 0, 0

	for _, direction := range path {
		switch direction {
		case 'W':
			posY++
			maxY = max(maxY, posY)
		case 'D':
			posX++
			maxX = max(maxX, posX)
		case 'S':
			posY--
			minY = min(minY, posY)
		case 'A':
			posX--
			minX = min(minX, posX)
		}
	}

	width := abs(minX) + maxX + 1
	height := abs(minY) + maxY + 1

	return width, height
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
