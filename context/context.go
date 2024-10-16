package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	_, cancel := context.WithTimeout(ctx, 10*time.Second)
	cancel()

}
