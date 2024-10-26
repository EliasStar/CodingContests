package main

import (
	"sync"
)

var filenames1 = []string{"level1/level1_example", "level1/level1_1", "level1/level1_2", "level1/level1_3", "level1/level1_4", "level1/level1_5"}
var filenames2 = []string{"level2/level2_example", "level2/level2_1", "level2/level2_2", "level2/level2_3", "level2/level2_4", "level2/level2_5"}

func main() {
	var wg sync.WaitGroup

	for _, v := range filenames1 {
		wg.Add(1)
		go level1(v, &wg)
	}

	for _, v := range filenames2 {
		wg.Add(1)
		go level2(v, &wg)
	}

	wg.Wait()
}
