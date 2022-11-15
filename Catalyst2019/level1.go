package main

import (
	"math"
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

	input := strings.Fields(string(in))

	vals := make([]int, len(input[1:]))
	for i, v := range input[1:] {
		vals[i], _ = strconv.Atoi(v)
	}

	var curMax = float64(vals[0])
	var curMin = float64(vals[0])
	for _, v := range vals {
		curMax = math.Max(curMax, float64(v))
		curMin = math.Min(curMin, float64(v))
	}

	os.WriteFile(filename+".out", []byte(strconv.Itoa(int(curMin))+" "+strconv.Itoa(int(curMax))), os.ModePerm)
	wg.Done()
}
