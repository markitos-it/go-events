package types

import (
	"context"
)

type EventRepository interface {
	Create(ctx context.Context, event *Event) error
	One(ctx context.Context, id *Id) (*Event, error)
	AllBySlugAndSource(ctx context.Context, slug *Slug, source *Slug) ([]*Event, error)
	Delete(ctx context.Context, id *Id) error
	CreateSubscription(ctx context.Context, sub *Subscription) error
	PullMessages(ctx context.Context, SubscriberName *Slug, slug *Slug, source *Slug) ([]*Queue, error)
	AckMessage(ctx context.Context, id *Id) error
	AckMessages(ctx context.Context, ids []*Id) error
}
