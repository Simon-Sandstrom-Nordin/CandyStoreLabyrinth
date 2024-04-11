package main

import (
	"CandyStoreLabyrinth/Labyrinth"
	"CandyStoreLabyrinth/Logic"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// learning objectives:
// wait group usage
// (un)buffered channels
// select statements
// data race coordination of some form
// prevent some kind of deadlock
// close channels purpose
func main() {
	rand.Seed(time.Now().UnixNano())

	Labyrinth.Entrance() // print welcome statement

	w1 := Labyrinth.Witch{
		"ShadowMiniRClone1",
		2, 3, rand.Intn(5) + 1, 7,
	}
	w2 := Labyrinth.Witch{
		"ShadowMiniRClone2",
		2, 3, rand.Intn(5) + 1, 7,
	}
	w3 := Labyrinth.Witch{
		"ShadowMiniRClone3",
		2, 3, rand.Intn(5) + 1, 8,
	}

	witches := []*Labyrinth.Witch{&w1, &w2, &w3}
	magicians := []string{"M", "H", "N"}

	// Create a channel to receive the results of the battles
	results := make(chan string)

	var wg sync.WaitGroup
	wg.Add(len(magicians))
	// entrance battle!
	for index, magician := range magicians {
		go Logic.Battle(magician, witches[index], results, &wg)
	}

	wg_results := sync.WaitGroup{}
	wg_results.Add(1)
	go func() {
		defer wg_results.Done()
		for result := range results {
			fmt.Println(result)
		}
	}()
	wg.Wait()
	close(results)
	wg_results.Wait()
}
