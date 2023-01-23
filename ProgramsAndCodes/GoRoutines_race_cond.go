package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	counter := 0

	const num = 100
	var wg sync.WaitGroup
	wg.Add(num)

	var mu sync.Mutex

	for i := 0; i < num; i++ {
		go func() {
			mu.Lock()
			temp := counter
			runtime.Gosched()
			temp++
			counter = temp
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count:", counter)
}
