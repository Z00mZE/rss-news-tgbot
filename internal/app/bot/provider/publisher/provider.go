package publisher

import (
	"context"
	"log/slog"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/wire"

	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/config"
	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/domain/provider"
)

type Provider struct {
	ctx             context.Context
	logger          *slog.Logger
	publishInterval time.Duration
	channelID       int64
	bot             *tgbotapi.BotAPI
}

var WireSet = wire.NewSet(
	NewProvider,
	wire.Bind(new(provider.Publisher), new(*Provider)),
)

func NewProvider(cfg *config.Config, logger *slog.Logger) (*Provider, error) {
	const op = `publisher`

	botAPI, clientError := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if clientError != nil {
		return nil, clientError
	}

	self := &Provider{
		logger:          logger.With(slog.String("op", op)),
		bot:             botAPI,
		channelID:       cfg.Telegram.ChannelID,
		publishInterval: cfg.Publisher.Interval,
	}
	return self, nil
}

func (p *Provider) Send(_ context.Context, message string) error {
	msgSettings := tgbotapi.NewMessage(p.channelID, message)
	msgSettings.ParseMode = "MarkdownV2"

	_, msgError := p.bot.Send(msgSettings)
	return msgError
}
