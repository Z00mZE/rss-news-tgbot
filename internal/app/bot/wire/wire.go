//go:build wireinject
// +build wireinject

package wire

import (
	"context"

	"github.com/google/wire"

	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot"
	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/config"
	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/pkg/logger"
	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/provider/publisher"
	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/repository/article"
	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/service/pubslihworker"
)

func InitApplication(ctx context.Context) (*bot.Application, func(), error) {
	panic(
		wire.Build(
			bot.NewApplication,
			config.NewConfig,
			logger.WireSet,
			publisher.WireSet,
			pubslihworker.WireSet,
			article.WireSet,
		),
	)
	return new(bot.Application), nil, nil
}
