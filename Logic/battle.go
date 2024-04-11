package Logic

import (
	"CandyStoreLabyrinth/Labyrinth"
	"fmt"
	"sync"
)

func Battle(magician string, witch *Labyrinth.Witch, results chan<- string, wg *sync.WaitGroup) {
	for witch.Health > 0 {
		witch.Health -= 1
		results <- fmt.Sprintf("%s struck %v and it's health is %v\n", magician, witch.Name, witch.Health)
		if witch.Health == 0 {
			wg.Done()
		}
		// time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
