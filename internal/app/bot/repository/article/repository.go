package article

import (
	"context"

	"github.com/google/wire"

	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/domain/model"
	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/domain/repository"
)

type Repository struct {
}

var WireSet = wire.NewSet(
	NewRepository,
	wire.Bind(new(repository.Article), new(*Repository)),
)

var _ repository.Article = (*Repository)(nil)

func NewRepository() *Repository {
	return new(Repository)
}

func (r *Repository) Articles(_ context.Context, _ uint) ([]model.Article, error) {
	dummy := []model.Article{
		{
			Link:    "https://z00mze.me",
			Title:   "Гайд: как подключить готового бота к чату или каналу в Телеграме",
			Summary: "Это только часть популярных действий, для которых необходимо добавление бота в Телеграме.",
		},
	}
	return dummy, nil
}
