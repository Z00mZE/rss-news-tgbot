package logger

import (
	"log/slog"

	"github.com/google/wire"

	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/config"
	"github.com/Z00mZE/rss-news-tgbot/internal/pkg/logger"
)

var WireSet = wire.NewSet(NewLogger)

func NewLogger(cfg *config.Config) *slog.Logger {
	return logger.NewSlogLogger(logger.ParseLogLevel(cfg.LogLevel))
}
