package main

import (
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func level3(filename string, wg *sync.WaitGroup) {
	in, err := os.ReadFile(filename + ".in")
	if err != nil {
		panic(err)
	}

	input := strings.Fields(string(in))

	rows, _ := strconv.Atoi(input[0])
	cols, _ := strconv.Atoi(input[1])

	vals := make([]int, len(input[2:]))
	for i, v := range input[2:] {
		vals[i], _ = strconv.Atoi(v)
	}

	found := false

	var lenest int
	var lenestPoses []int

	for side := 2; side <= int(math.Min(float64(rows), float64(cols))); side++ {

		curPoses := make([]int, 0)

	next:
		for i, v := range vals {

			for y := 0; y < side; y++ {
				for x := 0; x < side; x++ {
					offset := i + x + y*cols

					if offset >= len(vals) {
						break next
					}

					val := vals[offset]

					if v != val {
						continue next
					}
				}
			}

			found = true
			curPoses = append(curPoses, i)
		}

		if !found {
			break
		}

		lenest = side
		lenestPoses = curPoses
		found = false
	}

	out := strconv.Itoa(lenest)

	for _, v := range lenestPoses {
		y := v / cols
		x := v % cols
		out += "\n" + strconv.Itoa(x) + " " + strconv.Itoa(y)
	}

	os.WriteFile(filename+".out", []byte(out), os.ModePerm)
	wg.Done()
}
