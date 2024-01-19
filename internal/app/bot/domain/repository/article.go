package repository

import (
	"context"

	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/domain/model"
)

type Article interface {
	Articles(ctx context.Context, limit uint) ([]model.Article, error)
}
