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

	ctx, cancelCtx := context.WithCancel(ctx)
	defer cancelCtx()

	printCh := make(chan int)
	go doAnother(ctx, printCh)

	for num := 1; num <= 20; num++ {
		printCh <- num
		time.Sleep(300 * time.Millisecond)
	}

	log.Println("<doSomething>: Cancelling the context...")
	time.Sleep(2 * time.Second)
	cancelCtx()

	time.Sleep(300 * time.Millisecond)

	log.Printf("<doSomething>: %s's value is '%s'\n", ctxKey, ctx.Value(ctxKey))
	log.Printf("<doSomething>: finished\n")
}

func doAnother(ctx context.Context, printCh <-chan int) {
	labelFor:
	for {
		select {
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
