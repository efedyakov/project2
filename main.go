package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/efedyakov/project2/blob/main/internal/logger"
)

const appName = "bannerrotation"

func main() {
	ctx := context.Background()
	defer logger.Logger().Sync()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

}
