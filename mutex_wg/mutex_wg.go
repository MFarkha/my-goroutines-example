//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package mutex_wg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type LetterCounter struct {
	count int
	sync.Mutex
}

func processWord(w string, wg *sync.WaitGroup, lc *LetterCounter, i int) {
	defer wg.Done()
	lc.Lock()
	defer lc.Unlock()
	lc.count += len(w)
}

func readInput() (words []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		words = append(words, strings.Fields(line)...)
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return words
}

func StartMutexWg() {
	words := readInput() // waiting for stdin
	var lc LetterCounter
	var wg sync.WaitGroup
	for i, w := range words {
		wg.Add(1)
		go processWord(w, &wg, &lc, i)
	}
	wg.Wait()
	lc.Lock()
	defer lc.Unlock()
	fmt.Println("Total count of characters:", lc.count)
}
