package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} // global wait group

func threadA(data *int, mutex chan bool) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-mutex
		*data++

		fmt.Println("threadA updated Count to: ", *data)
		mutex <- true
	}
}

func threadB(data *int, mutex chan bool) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-mutex
		*data++

		fmt.Println("threadB updated Count to: ", *data)
		mutex <- true
	}
}

func main() {
	count := 0
	mu := make(chan bool, 1)
	mu <- true
	wg.Add(2)

	go threadA(&count, mu) // add 100 to count
	go threadB(&count, mu)

	wg.Wait()
}
