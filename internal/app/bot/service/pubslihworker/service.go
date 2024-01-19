package pubslihworker

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/wire"

	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/config"
	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/domain/provider"
	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/domain/repository"
	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/domain/service"
	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/service/pubslihworker/view"
	"github.com/Z00mZE/rss-news-tgbot/internal/pkg/logger/sl"
)

type PublishWorker struct {
	publisher         provider.Publisher
	logger            *slog.Logger
	publishInterval   time.Duration
	articleRepository repository.Article
}

var WireSet = wire.NewSet(
	NewPublishWorker,
	wire.Bind(new(service.PublishWorker), new(*PublishWorker)),
)

func NewPublishWorker(publisher provider.Publisher, articleRepository repository.Article, cfg *config.Config, logger *slog.Logger) *PublishWorker {
	return &PublishWorker{
		publishInterval:   cfg.Publisher.Interval,
		publisher:         publisher,
		articleRepository: articleRepository,
		logger:            logger,
	}
}
func (r *PublishWorker) Start(ctx context.Context) {
	r.logger.Info("start publish worker")
	ticker := time.NewTimer(r.publishInterval)
	go func() {
		defer ticker.Stop()
		defer r.logger.Info("stop publish worker")

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				r.publish(ctx)
				ticker.Reset(r.publishInterval)
			}
		}
	}()
}

func (r *PublishWorker) publish(ctx context.Context) {
	r.logger.Info("publish something into telegram")

	articles, articlesError := r.articleRepository.Articles(ctx, 1)
	if articlesError != nil {
		r.logger.Error("occurred error on read articles", sl.Error(articlesError))
		return
	}

	for _, article := range articles {
		if publishError := r.publisher.Send(ctx, view.ArticleRender(article)); publishError != nil {
			r.logger.Error("occurred error on publish article into TG", sl.Error(publishError))
		}
	}

}
