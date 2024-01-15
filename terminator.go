package terminator

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Terminator interface {
	AddTerminationTask(task func())
	Terminate()
}

type terminatorKey struct{}

func WithTerminator(ctx context.Context) (context.Context, Terminator) {
	t := &terminator{}

	ctx = context.WithValue(ctx, terminatorKey{}, t)
	ctx, t.cancel = signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-ctx.Done()
		t.executeTasks()
		os.Exit(0)
	}()

	return ctx, t
}

func MustFromContext(ctx context.Context) Terminator {
	t, ok := ctx.Value(terminatorKey{}).(Terminator)
	if !ok {
		panic("terminator not found in context")
	}
	return t
}

func FromContext(ctx context.Context) (Terminator, bool) {
	t, ok := ctx.Value(terminatorKey{}).(Terminator)
	return t, ok
}

type terminator struct {
	mu     sync.Mutex
	tasks  []func()
	cancel context.CancelFunc
}

func (t *terminator) executeTasks() {
	t.mu.Lock()
	defer t.mu.Unlock()
	for _, task := range t.tasks {
		task()
	}
}

func (t *terminator) AddTerminationTask(task func()) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.tasks = append(t.tasks, task)
}

func (t *terminator) Terminate() {
	t.cancel()
}
