package main

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

type input2 struct {
	mapSize         int
	pirateMap       [][]string
	coordinateNum   int
	coordinatePairs [][]point2D
}

func level2(filename string, wg *sync.WaitGroup) {
	in, err := os.ReadFile(filename + ".in")
	if err != nil {
		panic(err)
	}

	input := parseInput2(string(in))

	result := ""
	for _, coordPair := range input.coordinatePairs {
		if isIslandConnected2(copyPirateMap2(input.pirateMap), input.mapSize, coordPair[0], coordPair[1]) {
			result += "SAME\n"
		} else {
			result += "DIFFERENT\n"
		}
	}

	os.WriteFile(filename+".out", []byte(result), os.ModePerm)
	wg.Done()
}

func parseInput2(input string) input2 {
	fields := strings.Split(input, "\r\n")

	mapSize, _ := strconv.Atoi(fields[0])
	pirateMap := make([][]string, mapSize)
	for i := 0; i < mapSize; i++ {
		pirateMap[i] = strings.Split(fields[i+1], "")
	}

	coordinateNum, _ := strconv.Atoi(fields[mapSize+1])
	coordinatePairs := make([][]point2D, coordinateNum)
	for i := 0; i < coordinateNum; i++ {
		pairs := strings.Split(fields[i+mapSize+2], " ")

		coordinatePair := make([]point2D, len(pairs))
		for j, pair := range pairs {
			coords := strings.Split(pair, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])

			coordinatePair[j] = point2D{
				x: x,
				y: y,
			}
		}

		coordinatePairs[i] = coordinatePair
	}

	return input2{
		mapSize:         mapSize,
		pirateMap:       pirateMap,
		coordinateNum:   coordinateNum,
		coordinatePairs: coordinatePairs,
	}
}

func copyPirateMap2(pirateMap [][]string) [][]string {
	newPirateMap := make([][]string, len(pirateMap))

	for i := range pirateMap {
		newPirateMap[i] = make([]string, len(pirateMap[i]))
		copy(newPirateMap[i], pirateMap[i])
	}

	return newPirateMap
}

func isIslandConnected2(pirateMap [][]string, mapSize int, startPoint, endPoint point2D) bool {
	if startPoint.x < 0 ||
		startPoint.x >= mapSize ||
		startPoint.y < 0 ||
		startPoint.y >= mapSize ||
		pirateMap[startPoint.y][startPoint.x] != "L" {
		return false
	}

	if startPoint == endPoint {
		return true
	}

	pirateMap[startPoint.y][startPoint.x] = "#"

	if isIslandConnected2(pirateMap, mapSize, point2D{x: startPoint.x, y: startPoint.y - 1}, endPoint) {
		return true
	}

	if isIslandConnected2(pirateMap, mapSize, point2D{x: startPoint.x + 1, y: startPoint.y}, endPoint) {
		return true
	}

	if isIslandConnected2(pirateMap, mapSize, point2D{x: startPoint.x, y: startPoint.y + 1}, endPoint) {
		return true
	}

	if isIslandConnected2(pirateMap, mapSize, point2D{x: startPoint.x - 1, y: startPoint.y}, endPoint) {
		return true
	}

	return false
}
