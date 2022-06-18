package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

func level2(filename string, wg *sync.WaitGroup) {
	in, err := ioutil.ReadFile(filename + ".in")

	if err != nil {
		panic(err)
	}

	input := strings.Fields(string(in))

	vals := make([]int, len(input[1:]))

	for i, v := range input[1:] {
		vals[i], _ = strconv.Atoi(v)
	}

	var positions []int
	var lens []int

	length := 1
	pos := 0
	baseVal := vals[0]

	for i, v := range vals[1:] {
		if baseVal == v {
			length++
		} else {
			positions = append(positions, pos)
			lens = append(lens, length)

			baseVal = v
			pos = i + 1
			length = 1
		}
	}

	positions = append(positions, pos)
	lens = append(lens, length)

	var lenest int
	var index int

	for i, v := range lens {
		if lenest < v {
			index = i
			lenest = v
		}
	}

	ioutil.WriteFile(filename+".out", []byte(strconv.Itoa(int(positions[index]))+" "+strconv.Itoa(int(lens[index]))), os.ModePerm)
	wg.Done()
}
