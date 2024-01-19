package tgbot

import (
	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewBotAPIClient(cfg *config.Config) {
	bot, botError := tgbotapi.NewBotAPI(cfg.Telegram.Token)
}
