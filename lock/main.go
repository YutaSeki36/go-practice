package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	var mu sync.Mutex
	// wg := sync.WaitGroup{}
	ch := make(chan string)
	cend := make(chan struct{})

	// mu.Lockを外すとgoroutineでエラーが発生する。
	for i := 0; i < 1000; i++ {
		// wg.Add(1)
		go func() {
			mu.Lock()
			defer mu.Unlock()
			counter++
			ch <- "*"
			// wg.Done()

			if counter >= 1000 {
				close(cend)
			}
		}()
	}

	for {
		select {
		case str := <-ch:
			fmt.Println(str)
		case <-cend:
			fmt.Printf("\n%d\n", counter)

			return
		}
	}
	// wg.Wait()
}
