package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key1", "value1")
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	val1 := ctx.Value("key1")
	fmt.Println(val1)

	chanel := make(chan int)
	for i := 0; i < 10; i++ {
		go func(ctx context.Context) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
			chanel <- i
		}(ctx)
	}

	select {
	case msg := <-chanel:
		fmt.Println("Received:", msg)
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}
