package database

import (
	"context"
	"errors"
	"go-vents/internal/domain/shared"
	"go-vents/internal/domain/types"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventMariaDBRepository struct {
	db *gorm.DB
}

func NewEventMariaDBRepository(db *gorm.DB) types.EventRepository {
	return &EventMariaDBRepository{db: db}
}

func (r *EventMariaDBRepository) Create(ctx context.Context, event *types.Event) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(event).Error; err != nil {
			return err
		}

		var subscribers []string
		err := tx.Table("subscriptions").
			Where("event_name = ? AND source = ?", event.Slug, event.Source).
			Pluck("subscriber_name", &subscribers).Error
		if err != nil {
			return err
		}

		for _, subscriber := range subscribers {
			queueId := uuid.New().String()
			msg, err := types.NewQueueMessage(queueId, subscriber, event.Id)
			if err != nil {
				return err
			}
			if err := tx.Table("queue").Create(msg).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *EventMariaDBRepository) One(ctx context.Context, id *types.SharedId) (*types.Event, error) {
	var event types.Event
	err := r.db.WithContext(ctx).First(&event, "id = ?", id.Value()).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shared.ErrEventNotFound
		}
		return nil, shared.ErrEventBadRequest
	}
	return &event, nil
}

func (r *EventMariaDBRepository) AllBySlugAndSource(ctx context.Context, slug *types.Slug, source *types.Source) ([]*types.Event, error) {
	var events []*types.Event
	err := r.db.WithContext(ctx).
		Where("slug = ? AND source = ?", slug.Value(), source.Value()).
		Order("created_at DESC").
		Find(&events).Error
	return events, err
}

func (r *EventMariaDBRepository) Delete(ctx context.Context, id *types.SharedId) error {
	result := r.db.WithContext(ctx).Where("id = ?", id.Value()).Delete(&types.Event{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return shared.ErrEventNotFound
	}
	return nil
}

func (r *EventMariaDBRepository) CreateSubscription(ctx context.Context, sub *types.Subscription) error {
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

func (r *EventMariaDBRepository) PullMessages(ctx context.Context, slug *types.Slug, source *types.Source) ([]*types.Queue, error) {
	var results []*types.Queue
	err := r.db.WithContext(ctx).
		Table("queue q").
		Joins("JOIN events e ON q.event_id = e.id"). // MariaDB soporta este JOIN estándar perfectamente
		Where("e.slug = ? AND e.source = ? AND q.status = ?", slug.Value(), source.Value(), "pending").
		Find(&results).Error
	return results, err
}

func (r *EventMariaDBRepository) AckMessage(ctx context.Context, id *types.SharedId) error {
	result := r.db.WithContext(ctx).
		Model(&types.Queue{}).
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

func (r *EventMariaDBRepository) AckMessages(ctx context.Context, ids []*types.SharedId) error {
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
			return shared.ErrQueueMessageNotFound
		}

		return nil
	})
}
