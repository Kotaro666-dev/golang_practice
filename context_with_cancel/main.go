package main

import (
	"context"
	"fmt"
	"sync"
)

func sayHello(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		return "Hello", nil
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		defer wg.Done()

		word, err := sayHello(ctx)
		if err != nil {
			fmt.Printf("Cannot say hello: %v\n", err)
			cancel()
			return
		}
		fmt.Printf("%s\n", word)
	}()
	cancel() // context をキャンセルする
	wg.Wait()
}
