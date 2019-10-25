package main

import (
	"fmt"
	"sync"
)

// guarantee a1 happens before b2
// guarantee b1 happens before a2
var wg = sync.WaitGroup{} // global wait group

func threadA(x, y chan bool) {
	defer wg.Done()

	fmt.Println("A1")
	x <- true // signal A1 is done

	<-y // wait for B1
	fmt.Println("A2")
}

func threadB(x, y chan bool) {
	defer wg.Done()

	fmt.Println("B1")
	y <- true // signal B1 is done

	<-x // wait for A1
	fmt.Println("B2")
}

func main() {
	wg.Add(2)

	chn1 := make(chan bool, 1)
	chn2 := make(chan bool, 1)

	go threadA(chn1, chn2)
	go threadB(chn1, chn2)

	wg.Wait()
}
