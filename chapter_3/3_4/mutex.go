package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} // global wait group

func threadN(mutex chan int, name int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		count := <-mutex

		count++

		fmt.Printf("thread_%d updated Count to: %d\n", name, count)
		mutex <- count
	}
}

func main() {
	mu := make(chan int, 1)
	mu <- 0

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go threadN(mu, i)
	}

	wg.Wait()
}
