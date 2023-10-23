package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type input4 struct {
	mapSize         int
	pirateMap       [][]string
	coordinateNum   int
	coordinatePairs [][]point2D
}

func level4(filename string, wg *sync.WaitGroup) {
	in, err := os.ReadFile(filename + ".in")
	if err != nil {
		panic(err)
	}

	input := parseInput4(string(in))

	result := ""
	for _, coordPair := range input.coordinatePairs {
		path := []string{}

		out := findWaterConnection4(copyPirateMap4(input.pirateMap), input.mapSize, coordPair[0], coordPair[1])
		for _, point := range out {
			path = append(path, fmt.Sprintf("%v,%v", point.x, point.y))
		}

		result += strings.Join(path, " ") + "\n"
	}

	os.WriteFile(filename+".out", []byte(result), os.ModePerm)
	wg.Done()
}

func parseInput4(input string) input4 {
	fields := strings.Split(input, "\r\n")

	mapSize, _ := strconv.Atoi(fields[0])
	pirateMap := make([][]string, mapSize)
	for i := 0; i < mapSize; i++ {
		pirateMap[i] = strings.Split(fields[i+1], "")
	}

	coordinateNum, _ := strconv.Atoi(fields[mapSize+1])
	coordinatePairs := make([][]point2D, coordinateNum)
	for i := 0; i < coordinateNum; i++ {
		pair := strings.Split(fields[i+mapSize+2], " ")

		coordinatePair := make([]point2D, len(pair))
		for j, coords := range pair {
			coord := strings.Split(coords, ",")
			x, _ := strconv.Atoi(coord[0])
			y, _ := strconv.Atoi(coord[1])

			coordinatePair[j] = point2D{
				x: x,
				y: y,
			}
		}

		coordinatePairs[i] = coordinatePair
	}

	return input4{
		mapSize:         mapSize,
		pirateMap:       pirateMap,
		coordinateNum:   coordinateNum,
		coordinatePairs: coordinatePairs,
	}
}

func copyPirateMap4(pirateMap [][]string) [][]string {
	newPirateMap := make([][]string, len(pirateMap))

	for i := range pirateMap {
		newPirateMap[i] = make([]string, len(pirateMap[i]))
		copy(newPirateMap[i], pirateMap[i])
	}

	return newPirateMap
}

func findWaterConnection4(pirateMap [][]string, mapSize int, startPoint, endPoint point2D) []point2D {
	return nil
}
