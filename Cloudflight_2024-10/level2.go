package main

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

type input2 struct {
	roomCount int
	rooms     []room2
}

type room2 struct {
	x          int
	y          int
	tableCount int
}

func level2(filename string, wg *sync.WaitGroup) {
	in, err := os.ReadFile(filename + ".in")
	if err != nil {
		panic(err)
	}

	input := parseInput2(string(in))

	result := ""
	for _, room := range input.rooms {
		id := 1

		for y := 0; y < room.y; y++ {
			for x := 0; x < room.x/3; x++ {

				result += strconv.Itoa(id) + " " + strconv.Itoa(id) + " " + strconv.Itoa(id) + " "
				id++
			}

			result += "\n"
		}

		result += "\n"
	}

	os.WriteFile(filename+".out", []byte(result), os.ModePerm)
	wg.Done()
}

func parseInput2(input string) input2 {
	fields := strings.Fields(input)

	roomCount, _ := strconv.Atoi(fields[0])

	rooms := []room2{}

	for i := 0; i < len(fields[1:]); i += 3 {

		x, _ := strconv.Atoi(fields[1+i])
		y, _ := strconv.Atoi(fields[2+i])
		tableCount, _ := strconv.Atoi(fields[3+i])

		rooms = append(rooms, room2{x, y, tableCount})
	}

	return input2{
		roomCount: roomCount,
		rooms:     rooms,
	}
}
