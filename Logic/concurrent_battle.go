package Logic

import (
	"CandyStoreLabyrinth/Labyrinth"
	"fmt"
	"sync"
)

func Concurrent_battle(magician string, witch *Labyrinth.Witch, results chan<- string, wg *sync.WaitGroup, witchMutex sync.Mutex) {
	for witch.Health > 0 {
		witchMutex.Lock()
		witch.Health -= 1
		if witch.Health == 0 {
			wg.Done()
		}
		witchMutex.Unlock()
		results <- fmt.Sprintf("%s struck %v and it's health is %v\n", magician, witch.Name, witch.Health)

		// time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
