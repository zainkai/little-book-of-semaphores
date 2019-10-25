package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{} // global wait group

func threadA(chn chan bool) {
	defer wg.Done()

	fmt.Println("Printing from thread A!")
	chn <- true // release thread
}

func threadB(chn chan bool) {
	defer wg.Done()

	<-chn // block and wait
	fmt.Println("Printing from thread B!")
}

func main() {
	wg.Add(2)
	fmt.Println("Number of logical CPUs: ", runtime.GOMAXPROCS(0))

	chn := make(chan bool)
	go threadA(chn)
	go threadB(chn)

	wg.Wait() // ensure all goroutines finish
}
