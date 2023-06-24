package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// after timeout it will be canceled
	// withTimeout it's used by open DB connections, call rest apis
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // cancleled anyway
	fmt.Println("Just started...")

	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	select {
	case <-ctx.Done(): // case the ctx was canceled, do something
		fmt.Println("Request was canceled by ctx!")
	case <-time.After(time.Second * 12): // if the ctx did not canceled after that time, so do this
		fmt.Println("Finished normally!")
	}
}
