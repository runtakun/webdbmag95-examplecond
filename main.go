package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	c := sync.NewCond(&mu)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer mu.Unlock()

			fmt.Printf("waiting %d\n", i)
			mu.Lock()
			c.Wait()
			fmt.Printf("go %d\n", i)
		}(i)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("signaling!\n")
		c.Signal()
	}

	wg.Wait()
}
