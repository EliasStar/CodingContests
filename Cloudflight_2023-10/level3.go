package main

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

type input3 struct {
	mapSize   int
	pirateMap [][]string
	pathNum   int
	paths     [][]point2D
}

func level3(filename string, wg *sync.WaitGroup) {
	in, err := os.ReadFile(filename + ".in")
	if err != nil {
		panic(err)
	}

	input := parseInput3(string(in))

	result := ""
	for _, path := range input.paths {
		if isPathIntersecting3(path) {
			result += "INVALID\n"
		} else {
			result += "VALID\n"
		}
	}

	os.WriteFile(filename+".out", []byte(result), os.ModePerm)
	wg.Done()
}

func parseInput3(input string) input3 {
	fields := strings.Split(input, "\r\n")

	mapSize, _ := strconv.Atoi(fields[0])
	pirateMap := make([][]string, mapSize)
	for i := 0; i < mapSize; i++ {
		pirateMap[i] = strings.Split(fields[i+1], "")
	}

	pathNum, _ := strconv.Atoi(fields[mapSize+1])
	paths := make([][]point2D, pathNum)
	for i := 0; i < pathNum; i++ {
		coords := strings.Split(fields[i+mapSize+2], " ")

		path := make([]point2D, len(coords))
		for j, pair := range coords {
			coord := strings.Split(pair, ",")
			x, _ := strconv.Atoi(coord[0])
			y, _ := strconv.Atoi(coord[1])

			path[j] = point2D{
				x: x,
				y: y,
			}
		}

		paths[i] = path
	}

	return input3{
		mapSize:   mapSize,
		pirateMap: pirateMap,
		pathNum:   pathNum,
		paths:     paths,
	}
}

func isPathIntersecting3(path []point2D) bool {
	// check for same tile
	coords := make(map[point2D]struct{})
	for _, coord := range path {
		_, exists := coords[coord]
		if exists {
			return true
		}

		coords[coord] = struct{}{}
	}

	//check for diagonal intersections
	for i := 0; i+1 < len(path); i++ {
		dx := path[i+1].x - path[i].x
		dy := path[i+1].y - path[i].y

		if dx == 0 || dy == 0 {
			continue
		}

		candiate1 := point2D{x: path[i].x + dx, y: path[i].y}
		candiate2 := point2D{x: path[i].x, y: path[i].y + dy}

		for j := i + 2; j+1 < len(path); j++ {
			if (path[j] == candiate1 && path[j+1] == candiate2) ||
				(path[j] == candiate2 && path[j+1] == candiate1) {
				return true
			}
		}
	}

	return false
}
