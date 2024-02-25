package worklist

type Entry struct {
	Path string
}

type Worklist struct {
	jobs chan Entry
}

func (wl *Worklist) Add(e Entry) {
	wl.jobs <- e
}

func (wl *Worklist) Next() Entry {
	return <-wl.jobs
}

func New(buffSize int) Worklist {
	if buffSize <= 0 {
		panic("wrong buffer size of a channel Worklist")
	}
	return Worklist{make(chan Entry, buffSize)}
}

func NewEntry(path string) Entry {
	return Entry{path}
}

// generate 'empty jobs' to signal workers(goroutines) it is time them to quit
// `there is no files available to process`
func (wl *Worklist) Finalize(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		wl.Add(Entry{""})
	}
}
