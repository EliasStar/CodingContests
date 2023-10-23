package main

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

type Tournament struct {
	rockCount int
	paperCount int
	scissorCount int
}

func level3(filename string, wg *sync.WaitGroup) {
	in, err := os.ReadFile(filename + ".in")
	if err != nil {
		panic(err)
	}

	input := strings.Fields(string(in))[2:]

	var tournaments []Tournament
	for i := 0; i < len(input); i+=3 {
		rockCount, _ := strconv.Atoi(input[i][:len(input[i])-1])
		paperCount, _ := strconv.Atoi( input[i+1][:len( input[i+1])-1])
		scissorCount, _ := strconv.Atoi(input[i+2][:len(input[i+2])-1])

		tournaments = append(tournaments, Tournament{rockCount, paperCount, scissorCount})
	}

	var result string
	for _, tournament := range tournaments {
		for tournament.rockCount > 2 {
			result += "RRRP"
			tournament.rockCount -= 3
			tournament.paperCount--
		}

		if tournament.rockCount == 2 {
			result += "RPR"
			tournament.rockCount -= 2
			tournament.paperCount--
		} else if tournament.rockCount == 1 {
			result += "R"
			tournament.rockCount--
		}

		for tournament.paperCount > 0 {
			result += "P"
			tournament.paperCount--
		}

		for tournament.scissorCount > 0 {
			result += "S"
			tournament.scissorCount--
		}

		result += "\n"
	}

	os.WriteFile(filename+".out", []byte(result), os.ModePerm)
	wg.Done()
}
