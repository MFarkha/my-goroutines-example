//--Summary:
//  Create a grep clone that can do simple substring searching
//  within files. It must auto-recurse into subdirectories.
//
//--Requirements:
//* Use goroutines to search through the files for a substring match
//* Display matches to the terminal as they are found
//  * Display the line number, file path, and complete line containing the match
//* Recurse into any subdirectories looking for matches
//* Use any synchronization method to ensure that all files
//  are searched, and all results are displayed before the program
//  terminates.
//
//--Notes:
//* Program invocation should follow the pattern:
//    mgrep search_string search_dir

package mgrep

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/MFarkha/my-goroutines-example/mgrep/worker"
	"github.com/MFarkha/my-goroutines-example/mgrep/worklist"
	"go.wit.com/dev/alexflint/arg"
)

func discoverDirs(wl *worklist.Worklist, path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		nextPath := filepath.Join(path, e.Name())
		if e.IsDir() {
			discoverDirs(wl, nextPath)
		} else {
			wl.Add(worklist.NewEntry(nextPath))
		}
	}
}

var args struct {
	SearchTerm string `arg:"positional,required"` // required argument, otherwise it will abort
	SearchDir  string `arg:"positional"`          // if a user did not provide a directory, current directory will be used
}

func StartMgrep() {
	arg.MustParse(&args)
	var workersWg sync.WaitGroup

	wl := worklist.New(100)
	resultChan := make(chan worker.Result, 100)
	numWorkers := 10

	// start and end of the processing
	workersWg.Add(1)
	go func() {
		defer workersWg.Done()
		discoverDirs(&wl, args.SearchDir)
		wl.Finalize(numWorkers)
	}()

	// main processing
	for i := 0; i < numWorkers; i++ {
		workersWg.Add(1)
		go func() {
			defer workersWg.Done()
			for {
				workEntry := wl.Next()
				if workEntry.Path != "" {
					workerResults := worker.FindInFile(workEntry.Path, args.SearchTerm)
					if workerResults != nil {
						for _, result := range workerResults.Inner {
							resultChan <- result
						}
					}
				} else {
					return
				}
			}
		}()
	}
	blockWorkersWg := make(chan struct{})
	go func() {
		workersWg.Wait()
		close(blockWorkersWg)
	}()

	var displayWg sync.WaitGroup
	displayWg.Add(1)
	go func() {
		for {
			select {
			case r := <-resultChan:
				fmt.Printf("%s[%d]:%s\n", r.Path, r.LineNum, r.Line)
			case <-blockWorkersWg:
				if len(resultChan) == 0 {
					displayWg.Done()
					return
				}
			}
		}
	}()
	displayWg.Wait()
}
