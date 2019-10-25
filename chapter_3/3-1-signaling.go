package main

import (
	"fmt"
	"time"
)

func threadA(chn chan bool) {
	fmt.Println("Printing from thread A!")
	chn <- true // release thread
}

func threadB(chn chan bool) {
	<-chn // block and wait
	fmt.Println("Printing from thread B!")
}

func main() {
	chn := make(chan bool, 1)
	go threadA(chn)
	go threadB(chn)

	time.Sleep(100000)
}
