package generator

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandInt(min int, max int) <-chan int {
	out := make(chan int, 3)
	go func() {
		for {
			out <- rand.Intn(max-min+1) + min
		}
	}()
	return out
}

// certain amount of numbers to be generated - number defined by count
func generateRandIntn(count int, min int, max int) <-chan int {
	out := make(chan int, 1)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(max-min+1) + min
		}
		close(out)
	}()
	return out
}

func StartGenerator() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randInt := generateRandInt(1, 100)
	fmt.Printf("generateRandInt infinite:\n")
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)

	fmt.Println("----first option:")
	randIntRange1 := generateRandIntn(3, 1, 10)
	for i := range randIntRange1 {
		fmt.Println("generateRandIntn1: ", i)
	}
	fmt.Println("----second option:")
	randIntRange2 := generateRandIntn(3, 1, 10)
	for {
		n, open := <-randIntRange2
		if !open {
			break
		}
		fmt.Println("generateRandIntn2: ", n)
	}

}
