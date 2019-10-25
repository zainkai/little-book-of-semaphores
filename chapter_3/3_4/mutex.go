package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} // global wait group

func threadA(mutex chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		count := <-mutex

		count++

		fmt.Println("threadA updated Count to: ", count)
		mutex <- count
	}
}

func threadB(mutex chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		count := <-mutex

		count++

		fmt.Println("threadB updated Count to: ", count)
		mutex <- count
	}
}

func main() {
	mu := make(chan int, 1)
	mu <- 0
	wg.Add(2)

	go threadA(mu) // add 100 to count
	go threadB(mu)

	wg.Wait()
}
