package utils

import (
	"context"
	"log"
	"sync"
	"time"
)

func Startup(ctx context.Context, f func() error, success string) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer cancel()
		if err := f(); err != nil {
			log.Fatalln("Startup failed:", err)
		}
	}()
	if <-ctx.Done(); ctx.Err() == context.DeadlineExceeded {
		log.Println(success)
	}
	wg.Wait()
}
