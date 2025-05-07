package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// add a mock string containing a match event every 200 milliseconds,
// and protects to matchEvents with a mutex
func matchRecorder(matchEvents *[]string, mutex *sync.Mutex) {
	for i := 0; ; i++ {
		mutex.Lock()
		*matchEvents = append(*matchEvents, "Match event "+strconv.Itoa(i))
		mutex.Unlock()
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Appended match event")
	}
}

// run the client handler as a goroutine，each handling a connected user
// This function locks the shared slice containing the game events and make a copy of every element in the slice
// Simulates building a response to send to the user.
func clientHandler(mEvents *[]string, mutex *sync.Mutex, st time.Time) {
	for i := 0; i < 100; i++ {
		mutex.Lock()
		allEvents := copyAllEvents(mEvents)
		mutex.Unlock()
		timeTaken := time.Since(st)
		fmt.Println(len(allEvents), "events copied in", timeTaken)
	}
}

func copyAllEvents(mathchEvents *[]string) []string {
	allEvents := make([]string, len(*mathchEvents))
	for _, e := range *mathchEvents {
		allEvents = append(allEvents, e)
	}
	return allEvents
}

func main() {
	// Initializes a new mutex
	mutex := sync.Mutex{}
	var matchEvents = make([]string, 0, 10000)
	// Pre-populate the matchEvents slice with 10000 events
	for j := 0; j < 1000; j++ {
		matchEvents = append(matchEvents, "Match event")
	}
	// Start the matchRecorder goroutine
	go matchRecorder(&matchEvents, &mutex)
	// Start the clientHandler goroutines
	start := time.Now()
	for j := 0; j < 5000; j++ {
		go clientHandler(&matchEvents, &mutex, start)
	}
	time.Sleep(100 * time.Second)

}
