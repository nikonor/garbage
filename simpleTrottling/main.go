package main

import (
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go foo(ctx)

	for i := 0; i < 20; i++ {
		println("i=", i)
		if true && i == 7 {
			cancel()
		}
		time.Sleep(time.Second)
	}
}

func foo(ctx context.Context) {
	println("start::" + time.Now().String())
	select {
	case <-ctx.Done():
	case <-time.After(10 * time.Second):
	}
	println("finish::" + time.Now().String())
}
