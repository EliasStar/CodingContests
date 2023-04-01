package main

import (
	"os"
	"strings"
	"sync"
)

func level1(filename string, wg *sync.WaitGroup) {
	in, err := os.ReadFile(filename + ".in")
	if err != nil {
		panic(err)
	}

	fights := strings.Fields(string(in))[1:]

	var result string
	for _, fight := range fights {
		result += winnerOf(string(fight[0]), string(fight[1])) + "\n"
	}

	os.WriteFile(filename+".out", []byte(result), os.ModePerm)
	wg.Done()
}

func winnerOf(style1, style2 string) string {
	switch style1 {
	case "R":
		if style2 == "P" {
			return "P"
		} else {
			return "R"
		}
	case "P":
		if style2 == "S" {
			return "S"
		} else {
			return "P"
		}
	case "S":
		if style2 == "R" {
			return "R"
		} else {
			return "S"
		}
	}

	return "_"
}
