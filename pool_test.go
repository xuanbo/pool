package pool

import (
	"fmt"
	"testing"
	"time"
)

func TestNewWorkerPool(t *testing.T) {
	wp := NewWorkerPool(5, 10)
	wp.Start()
	for i := 0; i < 100; i++ {
		count := i
		wp.Add(func() {
			fmt.Printf("%d\n", count)
		})
	}
	wp.Stop()
	time.Sleep(2 * time.Second)
}
