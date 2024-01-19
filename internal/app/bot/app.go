package bot

import (
	"context"
	"log/slog"
)

type Application struct {
	ctx    context.Context
	logger *slog.Logger
}

func NewApplication(ctx context.Context, logger *slog.Logger) *Application {
	return &Application{ctx: ctx, logger: logger}
}

func (a *Application) Run() {
	a.logger.Info("start")
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
