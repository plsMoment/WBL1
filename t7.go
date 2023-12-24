package main

import (
	"fmt"
	"sync"
)

// asMap struct is concurrent safe map.
type asMap struct {
	data map[int]int
	m    sync.RWMutex
}

func (a *asMap) Add(key, value int) {
	a.m.Lock()
	a.data[key] = value
	a.m.Unlock()
}

func (a *asMap) Load(key int) int {
	a.m.RLock()
	value := a.data[key]
	a.m.RUnlock()
	return value
}

func main() {
	asyncMap := &asMap{data: make(map[int]int)}

	var wg sync.WaitGroup
	wg.Add(8)

	// fill map
	for i := 0; i < 8; i++ {
		go func(i int) {
			asyncMap.Add(i, i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	wg.Add(8)

	// get stored values from map
	for i := 0; i < 8; i++ {
		go func(key int) {
			fmt.Printf("key: %d, value: %d\n", key, asyncMap.Load(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
