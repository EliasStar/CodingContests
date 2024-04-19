package main

import (
	"sync"
)

var filenames1 = []string{"level1/level1_example", "level1/level1_1", "level1/level1_2", "level1/level1_3", "level1/level1_4", "level1/level1_5"}
var filenames2 = []string{"level2/level2_example", "level2/level2_1", "level2/level2_2", "level2/level2_3", "level2/level2_4", "level2/level2_5"}
var filenames3 = []string{"level3/level3_example", "level3/level3_1", "level3/level3_2", "level3/level3_3", "level3/level3_4", "level3/level3_5"}
var filenames4 = []string{"level4/level4_example", "level4/level4_1", "level4/level4_2", "level4/level4_3", "level4/level4_4", "level4/level4_5"}

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

	for _, v := range filenames3 {
		wg.Add(1)
		go level3(v, &wg)
	}

	for _, v := range filenames4 {
		wg.Add(1)
		go level4(v, &wg)
	}

	wg.Wait()
}
