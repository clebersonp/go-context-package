package main

import (
	"context"
	"log"
	"time"
)

const (
	ctxKey = "myKey"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	ctx := context.Background()

	ctx = context.WithValue(ctx, ctxKey, "myValue")

	doSomething(ctx)
}

func doSomething(ctx context.Context) {

	// automatically cancel the context 1500 milliseconds (1.5 seconds) after the function starts by using the time.Now function
	deadLine := time.Now().Add(1500 * time.Millisecond)
	ctx, cancelCtx := context.WithDeadline(ctx, deadLine)

	// the cancel function is still required to be called in order to clean up any resources that were used, so this is more of a safety measure.
	defer cancelCtx()

	printCh := make(chan int)
	go doAnother(ctx, printCh)

labelFor:
	for num := 1; num <= 20; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				log.Printf("<doSomenthing> err: %s\n", err)
			}
			log.Printf("<doSomething>: breaking\n")
			break labelFor
		}
	}

	log.Println("<doSomething>: Cancelling the context...")
	time.Sleep(2 * time.Second)

	// The code can now potentially end by calling cancelCtx directly or an automatic cancelation via the deadline
	cancelCtx()

	time.Sleep(300 * time.Millisecond)

	log.Printf("<doSomething>: %s's value is '%s'\n", ctxKey, ctx.Value(ctxKey))
	log.Printf("<doSomething>: finished\n")
}

func doAnother(ctx context.Context, printCh <-chan int) {
labelFor:
	for {
		select {
		// Once the deadline is exceeded, both the doAnother and doSomething functions finish running because they’re both watching for the ctx.Done channel to close.
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				log.Printf("<doAnother> err: %s\n", err)
			}
			log.Printf("<doAnother>: finished\n")
			break labelFor
		case num := <-printCh:
			log.Printf("<doAnother>: %d\n", num)
		}
	}
}
