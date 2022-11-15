package main

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

func level1(filename string, wg *sync.WaitGroup) {
	in, err := os.ReadFile(filename + ".in")
	if err != nil {
		panic(err)
	}

	counter := strings.Count(string(in), "C")

	os.WriteFile(filename+".out", []byte(strconv.Itoa(counter)), os.ModePerm)
	wg.Done()
}
