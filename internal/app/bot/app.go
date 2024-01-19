package bot

import (
	"context"
	"log/slog"

	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/domain/service"
)

type Application struct {
	ctx           context.Context
	logger        *slog.Logger
	publishWorker service.PublishWorker
}

func NewApplication(ctx context.Context, publishWorker service.PublishWorker, logger *slog.Logger) *Application {
	return &Application{
		ctx:           ctx,
		logger:        logger,
		publishWorker: publishWorker,
	}
}

func (a *Application) Run() {
	a.logger.Info("start")
	a.publishWorker.Start(a.ctx)

loop:
	for {
		select {
		case <-a.ctx.Done():
			a.logger.Info("parent context was closed")
			break loop
			//case serverError := <-serverErrorCh:
			//	a.logger.Error("shutting down the server", sl.Error(serverError))
			//	break loop
		}
	}
}
