package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	errCh := make(chan error, 1)
	errCh <- Run(ctx)

	select {
	case err := <-errCh:
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}
