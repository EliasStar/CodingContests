package main

import (
	"sync"
)

func level5(filename string, wg *sync.WaitGroup) {
	wg.Done()
}
