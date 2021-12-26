package text

import (
	"fmt"
	"sync"
)

func T1() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("有点强人锁男")
		mu.Lock()
		wg.Done()
	}()
	wg.Wait()
	mu.Unlock()
}
