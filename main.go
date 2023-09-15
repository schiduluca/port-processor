package main

import (
	"context"
	"fmt"
	"github.com/schiduluca/port-processor/repo"
	"github.com/schiduluca/port-processor/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	done := make(chan error, 1)

	f, err := os.Open("ports.json")

	if err != nil {
		fmt.Printf("could not read the file: %v", err)
		return
	}
	defer f.Close()

	storage := repo.NewMemDB()
	processor := service.NewJSONProcessor(storage)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err = processor.Process(ctx, f)
		if err != nil {
			done <- err
			return
		}
		done <- nil
	}()

	fmt.Println("awaiting processing...")
	select {
	case innerErr := <-done:
		if innerErr != nil {
			fmt.Println("processing exited with an error:", innerErr)
			return
		}
		fmt.Println("processing done")
	case <-sigc:
		fmt.Println("unexpected exit, gracefully shutting down")
		cancel()
	}
}
