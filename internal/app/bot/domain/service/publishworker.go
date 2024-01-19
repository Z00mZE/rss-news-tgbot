package service

import (
	"context"
)

type PublishWorker interface {
	Start(ctx context.Context)
}
