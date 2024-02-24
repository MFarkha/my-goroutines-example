package waitgroups

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func StartWaitGroups() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var wg sync.WaitGroup
	counter := 0

	for i := 0; i < 5; i++ {
		counter++
		wg.Add(1)
		go func() {
			defer func() {
				fmt.Println(counter, "goroutines remaining")
				counter--
				wg.Done()
			}()
			duration := time.Duration(rand.Intn(500) * int(time.Millisecond))
			fmt.Println("Waiting for", duration)
			time.Sleep(duration)
		}()
	}
	fmt.Println("waiting for goroutines to finish:")
	wg.Wait()
	fmt.Println("Done")
}
