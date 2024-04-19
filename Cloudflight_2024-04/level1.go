package main

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

type input1 struct {
	pathCount int
	paths     []string
}

func level1(filename string, wg *sync.WaitGroup) {
	in, err := os.ReadFile(filename + ".in")
	if err != nil {
		panic(err)
	}

	input := parseInput1(string(in))

	result := ""
	for _, path := range input.paths {
		result += strconv.Itoa(countLetter(path, 'W')) + " "
		result += strconv.Itoa(countLetter(path, 'D')) + " "
		result += strconv.Itoa(countLetter(path, 'S')) + " "
		result += strconv.Itoa(countLetter(path, 'A')) + "\n"
	}

	os.WriteFile(filename+".out", []byte(result), os.ModePerm)
	wg.Done()
}

func parseInput1(input string) input1 {
	fields := strings.Fields(input)

	pathCount, _ := strconv.Atoi(fields[0])

	return input1{
		pathCount: pathCount,
		paths:     fields[1 : pathCount+1],
	}
}

func countLetter(str string, letter rune) int {
	count := 0

	for _, char := range str {
		if char == letter {
			count++
		}
	}

	return count
}
