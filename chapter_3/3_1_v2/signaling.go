package main

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/zainkai/little-book-of-semaphores/signaler"
)

var wg = sync.WaitGroup{} // global wait group

func threadA(s *signaler.Signaler) {
	defer wg.Done()

	fmt.Println("Printing from thread A!")
	s.Signal()
}

func threadB(s *signaler.Signaler) {
	defer wg.Done()

	s.Wait()
	fmt.Println("Printing from thread B!")
}

func main() {
	wg.Add(2)
	fmt.Println("Number of logical CPUs: ", runtime.GOMAXPROCS(0))

	s := signaler.New()
	go threadA(s)
	go threadB(s)

	wg.Wait() // ensure all goroutines finish
}
