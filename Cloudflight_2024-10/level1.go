package main

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

type input1 struct {
	roomCount int
	rooms     []room1
}

type room1 struct {
	x int
	y int
}

func level1(filename string, wg *sync.WaitGroup) {
	in, err := os.ReadFile(filename + ".in")
	if err != nil {
		panic(err)
	}

	input := parseInput1(string(in))

	result := ""
	for _, room := range input.rooms {
		result += strconv.Itoa(room.y*room.x/3) + "\n"
	}

	os.WriteFile(filename+".out", []byte(result), os.ModePerm)
	wg.Done()
}

func parseInput1(input string) input1 {
	fields := strings.Fields(input)

	roomCount, _ := strconv.Atoi(fields[0])

	rooms := []room1{}

	for i := 0; i < len(fields[1:]); i += 2 {

		x, _ := strconv.Atoi(fields[1+i])
		y, _ := strconv.Atoi(fields[2+i])

		rooms = append(rooms, room1{x, y})
	}

	return input1{
		roomCount: roomCount,
		rooms:     rooms,
	}
}
