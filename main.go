package main

import (
	"CandyStoreLabyrinth/Labyrinth"
	"CandyStoreLabyrinth/Logic"
	"fmt"
	"math/rand"
	"strconv"
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

	// finally, they all take down the queen bee at the heart of the labyrinth

	Queen_BEE := Labyrinth.Witch{
		"QueenBee",
		30, 300, rand.Intn(100) + 1, 90,
	}

	results_final := make(chan string, 101)
	wg_final := sync.WaitGroup{}
	wg_final.Add(1)
	for _, magician := range magicians {
		fmt.Println(magician + "------------------------------------------------------------")
	}

	var witchMutex sync.Mutex
	for _, magician := range magicians {
		go Logic.Concurrent_battle(magician, &Queen_BEE, results_final, &wg_final, witchMutex)
	}
	wg_results_final := sync.WaitGroup{}
	wg_results_final.Add(1)
	go func() {
		defer wg_results_final.Done()
		iteration := 1
		for result_final := range results_final {
			fmt.Println(strconv.Itoa(iteration) + " : " + result_final) // Convert inte to string
			iteration++
		}
	}()
	wg_final.Wait()
	close(results_final)
	wg_results_final.Wait()

}
