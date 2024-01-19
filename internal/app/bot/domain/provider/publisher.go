package provider

import (
	"context"
)

type Publisher interface {
	Send(context.Context, string) error
}
