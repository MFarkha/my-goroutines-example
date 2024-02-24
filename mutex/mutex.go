package mutex

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Hits struct {
	count int
	sync.Mutex
}

func wait() {
	time.Sleep(time.Duration(rand.Intn(500) * int(time.Millisecond)))
}

func serve(wg *sync.WaitGroup, hc *Hits, i int) {
	wait()
	hc.Lock()
	defer hc.Unlock()
	defer wg.Done()
	hc.count++
	fmt.Println("Iteration ", i)
}

func StartMutex() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var wg sync.WaitGroup

	hitCounter := Hits{}

	for i := 0; i < 20; i++ {
		iteration := i
		wg.Add(1)
		go serve(&wg, &hitCounter, iteration)
	}

	fmt.Println("Waiting for goroutines...")
	wg.Wait()

	hitCounter.Lock()
	defer hitCounter.Unlock()

	fmt.Println("Total hits: ", hitCounter.count)
}
