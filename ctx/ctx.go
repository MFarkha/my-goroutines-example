package ctx

import (
	"context"
	"fmt"
	"time"
)

func sampleOperation(ctx context.Context, msg string, msDelay time.Duration) <-chan string {
	out := make(chan string)
	go func() {
		for {
			select {
			case <-time.After(msDelay * time.Millisecond):
				out <- fmt.Sprintf("%v operation completed", msg)
			case <-ctx.Done():
				out <- fmt.Sprintf("%v operation aborted", msg)
			}

		}
	}()
	return out
}

func StartContext() {
	ctx := context.Background()               // creates an empty `context.Context`
	ctx, cancelCtx := context.WithCancel(ctx) // returns `cancelCtx` function which allows to `cancel` the `ctx`
	webserver := sampleOperation(ctx, "webserver", 200)
	microservice := sampleOperation(ctx, "microservice", 500)
	database := sampleOperation(ctx, "database", 900)

MainLoop:
	for {
		select {
		case output := <-webserver:
			fmt.Println(output)
		case output := <-microservice:
			fmt.Println(output)
			fmt.Println("cancel context")
			cancelCtx()
			break MainLoop
		case output := <-database:
			fmt.Println(output)
		}
	}
	fmt.Println(<-database)
}
