package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go-vents/internal/domain/shared"
	"go-vents/internal/domain/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PulledMessage struct {
	QueueId string
	EventId string
	Name    string
	Source  string
	Payload string
}

type EventPostgresRepository struct {
	db *gorm.DB
}

func NewEventPostgresRepository(db *gorm.DB) types.EventRepository {
	return &EventPostgresRepository{db: db}
}

func (r *EventPostgresRepository) Create(ctx context.Context, event *types.Event) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(event).Error; err != nil {
			return fmt.Errorf("error inserting main event: %w", err)
		}

		var subscribers []string
		err := tx.Table("subscriptions").
			Where("event_name = ? AND source = ?", event.Slug, event.Source).
			Pluck("subscriber_name", &subscribers).Error

		if err != nil {
			return fmt.Errorf("error searching subscriptions: %w", err)
		}

		for _, subscriber := range subscribers {
			queueId := uuid.New().String()

			msg, err := types.NewQueueMessage(queueId, subscriber, event.Id, event.Payload)
			if err != nil {
				return fmt.Errorf("error al construir QueueMessage: %w", err)
			}

			if err := tx.Table("queue").Create(msg).Error; err != nil {
				return fmt.Errorf("error al insertar en la cola de %s: %w", subscriber, err)
			}
		}

		return nil
	})
}

func (r *EventPostgresRepository) One(ctx context.Context, id *types.Id) (*types.Event, error) {
	var event types.Event

	err := r.db.WithContext(ctx).
		First(&event, "id = ?", id.Value()).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shared.ErrEventNotFound
		}
		return nil, shared.ErrEventBadRequest
	}

	return &event, nil
}

func (r *EventPostgresRepository) AllBySlugAndSource(ctx context.Context, slug *types.Slug, source *types.Slug) ([]*types.Event, error) {
	var events []*types.Event

	err := r.db.WithContext(ctx).
		Where("slug = ? AND source = ?", slug.Value(), source.Value()).
		Order("created_at DESC").
		Find(&events).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shared.ErrEventNotFound
		}
		return nil, shared.ErrEventBadRequest
	}

	return events, nil
}

func (r *EventPostgresRepository) Delete(ctx context.Context, id *types.Id) error {
	result := r.db.WithContext(ctx).
		Where("id = ?", id.Value()).
		Delete(&types.Event{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return shared.ErrEventNotFound
	}

	return nil
}

func (r *EventPostgresRepository) CreateSubscription(ctx context.Context, sub *types.Subscription) error {
	var count int64
	err := r.db.WithContext(ctx).Table("subscriptions").
		Where("subscriber_name = ? AND event_name = ? AND source = ?", sub.SubscriberName, sub.EventName, sub.Source).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	return r.db.WithContext(ctx).
		Table("subscriptions").
		Create(sub).Error
}

func (r *EventPostgresRepository) PullMessages(
	ctx context.Context,
	subscriberName *types.Slug,
	slug *types.Slug,
	source *types.Slug,
) ([]*types.Queue, error) {

	var results []*types.Queue

	err := r.db.WithContext(ctx).
		Table("queue q").
		Joins("JOIN events e ON q.event_id = e.id").
		Where("q.subscriber_name = ? AND e.slug = ? AND e.source = ? AND q.status = ?", subscriberName.Value(), slug.Value(), source.Value(), "pending").
		Clauses(clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *EventPostgresRepository) AckMessage(ctx context.Context, id *types.Id) error {
	result := r.db.WithContext(ctx).
		Table("queue").
		Where("id = ? AND status = ?", id.Value(), "pending").
		Updates(map[string]interface{}{
			"status":     "processed",
			"updated_at": time.Now(),
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return shared.ErrQueueMessageNotFound
	}

	return nil
}

func (r *EventPostgresRepository) AckMessages(ctx context.Context, ids []*types.Id) error {
	if len(ids) == 0 {
		return nil
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var values []string
		for _, id := range ids {
			values = append(values, id.Value())
		}

		result := tx.Table("queue").
			Where("id IN ? AND status = ?", values, "pending").
			Updates(map[string]interface{}{
				"status":     "processed",
				"updated_at": time.Now(),
			})

		if result.Error != nil {
			return result.Error
		}
		if int(result.RowsAffected) != len(ids) {
			return shared.ErrQueueMessageNotFound // Or another specific error if needed
		}

		return nil
	})
}
