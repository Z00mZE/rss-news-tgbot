package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/wire"
)

func main() {
	ctx, ctxClose := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer ctxClose()

	app, appClose, appError := wire.InitApplication(ctx)
	defer appClose()

	if appError != nil {
		os.Exit(1)
	}

	app.Run()
}
