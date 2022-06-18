package main

import (
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func level4(filename string, wg *sync.WaitGroup) {
	in, err := ioutil.ReadFile(filename + ".in")

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

			curMin := float64(v)
			curMax := float64(v)

			for y := 0; y < side; y++ {
				for x := 0; x < side; x++ {
					offset := i + x + y*cols

					if offset >= len(vals) {
						break next
					}

					val := float64(vals[offset])

					curMax = math.Max(curMax, val)
					curMin = math.Min(curMin, val)
				}
			}

			if curMax-curMin <= 2 {
				curPoses = append(curPoses, i)
				found = true
			}
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

	ioutil.WriteFile(filename+".out", []byte(out), os.ModePerm)
	wg.Done()
}
