package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func nightClubPatron(sem chan bool, name int) {
	for {
		<-sem // wait to enter night club
		fmt.Printf("Patron_%d entered night club.\n", name)

		t := rand.Int63n(10)
		time.Sleep(time.Duration(t) * time.Second)

		fmt.Printf("Patron_%d is leaving night club.\n", name)
		sem <- true // leave night club
	}
}

func main() {
	nightClubSize := 10
	nightClub := make(chan bool, nightClubSize)
	for i := 0; i < nightClubSize; i++ {
		nightClub <- true
	}

	for i := 0; i < nightClubSize*10; i++ {
		go nightClubPatron(nightClub, i)
	}

	runtime.Goexit()
}
