package main

import (
	"context"
	"fmt"
	"time"

	"github.com/zrbecker/terminator"
)

func main() {
	ctx, t := terminator.WithTerminator(context.Background())
	defer t.Terminate()

	for i := 0; i < 10; i++ {
		go func(i int) {
			terminator.MustFromContext(ctx).AddTerminationTask(func() {
				fmt.Printf("Performing cleanup for %d...\n", i)
			})
		}(i)
	}

	fmt.Println("Application active. Press Ctrl+C to stop or wait 10 seconds for automatic shutdown.")
	time.Sleep(10 * time.Second)
}
