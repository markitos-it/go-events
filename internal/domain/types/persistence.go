package types

import (
	"context"
)

type EventRepository interface {
	Create(ctx context.Context, event *Event) error
	One(ctx context.Context, id *SharedId) (*Event, error)
	AllBySlugAndSource(ctx context.Context, slug *Slug, source *Source) ([]*Event, error)
	Delete(ctx context.Context, id *SharedId) error
	CreateSubscription(ctx context.Context, sub *Subscription) error
	PullMessages(ctx context.Context, SubscriberName *Name, slug *Slug, source *Source) ([]*Queue, error)
	AckMessage(ctx context.Context, id *SharedId) error
	AckMessages(ctx context.Context, ids []*SharedId) error
}
