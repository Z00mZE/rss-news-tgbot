//go:build wireinject
// +build wireinject

package wire

import (
	"context"

	"github.com/google/wire"

	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot"
	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/config"
	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/pkg/logger"
)

func InitApplication(ctx context.Context) (*bot.Application, func(), error) {
	panic(
		wire.Build(
			wire.NewSet(
				bot.NewApplication,
				config.NewConfig,
				logger.WireSet,
			),
		),
	)
	return new(bot.Application), nil, nil
}
