package main

import (
	"context"
	"fmt"
	"time"

	"github.com/zrbecker/go-terminator"
)

func main() {
	ctx, t := terminator.WithTerminator(context.Background())
	defer t.Terminate()

	terminator.MustFromContext(ctx).AddTerminationTask(func() {
		fmt.Println("Performing cleanup...")
	})

	fmt.Println("Application active. Press Ctrl+C to stop or wait 10 seconds for automatic shutdown.")
	time.Sleep(10 * time.Second)
}
