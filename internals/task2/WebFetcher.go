package task2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func WebFetcher(url string, hasilChannel chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	randomDelay := rand.Intn(3500)
	time.Sleep(time.Duration(randomDelay) * time.Millisecond)

	hasilChannel <- fmt.Sprintf("fetched: %s", url)
}
