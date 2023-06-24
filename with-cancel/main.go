package main

import (
	"context"
	"fmt"
	"time"
)

const (
	key = "init-time"
)

func main() {
	// context.Background() is a root(parent) blank context
	ctx, cancel := context.WithCancel(context.WithValue(context.Background(), key, time.Now().UnixMilli()))
	defer cancel() // close it anyway at the end
	fmt.Print("Just started...\n")

	// create and execute a new goroutine
	go func(cancel context.CancelFunc) {
		time.Sleep(4 * time.Second) // after sometime cancel the context
		cancel()
	}(cancel)

	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	select {
	case <-ctx.Done(): // channel is closed when the returned cancel function is called or when the parent context's Done channel is closed, whichever happens first.
		fmt.Printf("Request was canceled. ctx value for key '%s': %v\n", key, ctx.Value(key))
	case <-time.After(5 * time.Second):
		fmt.Printf("Finished normally\n")
	}
}
