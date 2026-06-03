package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

var database map[string]string

func RandomRange(min, max float64) string {
	return fmt.Sprintf("%.2f", min+rand.Float64()*(max-min))
}

func Write(ctx context.Context, mu *sync.RWMutex) {
	if database == nil {
		database = make(map[string]string)
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			mu.Lock()
			database[RandomRange(-90.0, 90.0)] = RandomRange(-180.0, 180.0)
			mu.Unlock()
		}
	}

}

func Read(ctx context.Context, mu *sync.RWMutex, reader string) {
	mu.RLock()
	for lat, lon := range database {
		select {
		case <-ctx.Done():
			return
		default:

			fmt.Println("reader=", reader, "lat=", lat, "lon=", lon)
		}

	}
	mu.RUnlock()
}

func main() {
	mu := new(sync.RWMutex)

	database = make(map[string]string)
	for i := 0; i < 100; i++ {
		database[RandomRange(-90.0, 90.0)] = RandomRange(-180.0, 180.0)
	}

	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 10; i++ {
		go Read(ctx, mu, strconv.Itoa(i+1))
	}

	go Write(ctx, mu)

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	cancel()
}
