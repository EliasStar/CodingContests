package main

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

type point2D struct {
	x int
	y int
}

type input1 struct {
	mapSize       int
	pirateMap     [][]string
	coordinateNum int
	coordinates   []point2D
}

func level1(filename string, wg *sync.WaitGroup) {
	in, err := os.ReadFile(filename + ".in")
	if err != nil {
		panic(err)
	}

	input := parseInput1(string(in))

	result := ""
	for _, coord := range input.coordinates {
		result += input.pirateMap[coord.y][coord.x] + "\n"
	}

	os.WriteFile(filename+".out", []byte(result), os.ModePerm)
	wg.Done()
}

func parseInput1(input string) input1 {
	fields := strings.Fields(input)

	mapSize, _ := strconv.Atoi(fields[0])
	pirateMap := make([][]string, mapSize)
	for i := 0; i < mapSize; i++ {
		pirateMap[i] = strings.Split(fields[i+1], "")
	}

	coordinateNum, _ := strconv.Atoi(fields[mapSize+1])
	coordinates := make([]point2D, coordinateNum)
	for i := 0; i < coordinateNum; i++ {
		coords := strings.Split(fields[i+mapSize+2], ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])

		coordinates[i] = point2D{
			x: x,
			y: y,
		}
	}

	return input1{
		mapSize:       mapSize,
		pirateMap:     pirateMap,
		coordinateNum: coordinateNum,
		coordinates:   coordinates,
	}
}
