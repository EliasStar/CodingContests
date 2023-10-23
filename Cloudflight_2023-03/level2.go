package main

import (
	"os"
	"strings"
	"sync"
)

func level2(filename string, wg *sync.WaitGroup) {
	in, err := os.ReadFile(filename + ".in")
	if err != nil {
		panic(err)
	}

	tournaments := strings.Fields(string(in))[2:]

	var result string
	for _, tournament := range tournaments {
		var tmp string
		for i := 0; i < len(tournament); i += 2 {
			tmp += winnerOf(string(tournament[i]), string(tournament[i+1]))
		}

		for i := 0; i < len(tmp); i += 2 {
			result += winnerOf(string(tmp[i]), string(tmp[i+1]))
		}

		result += "\n"
	}

	os.WriteFile(filename+".out", []byte(result), os.ModePerm)
	wg.Done()
}
