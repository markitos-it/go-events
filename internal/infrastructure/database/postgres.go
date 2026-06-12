package database

import (
	"context"
	"errors"

	"govent/internal/domain/shared"
	"govent/internal/domain/types"

	"gorm.io/gorm"
)

type EventPostgresRepository struct {
	db *gorm.DB
}

func NewEventPostgresRepository(db *gorm.DB) types.EventRepository {
	return &EventPostgresRepository{db: db}
}

func (r *EventPostgresRepository) Create(ctx context.Context, event *types.Event) error {
	return r.db.WithContext(ctx).Create(event).Error
}

func (r *EventPostgresRepository) One(ctx context.Context, id *types.EventId) (*types.Event, error) {
	var event types.Event

	err := r.db.WithContext(ctx).
		First(&event, "id = ?", id.Value()).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shared.ErrEventNotFound
		}
		return nil, err
	}

	return &event, nil
}

func (r *EventPostgresRepository) AllByName(ctx context.Context, name *types.EventName) ([]*types.Event, error) {
	var events []*types.Event

	err := r.db.WithContext(ctx).
		Where("name = ?", name.Value()).
		Order("created_at DESC").
		Find(&events).
		Error

	if err != nil {
		return nil, shared.ErrEventBadRequest
	}

	return events, nil
}

func (r *EventPostgresRepository) AllByNameAndSource(ctx context.Context, name *types.EventName, source *types.EventSource) ([]*types.Event, error) {
	var events []*types.Event

	err := r.db.WithContext(ctx).
		Where("name = ? AND source = ?", name.Value(), source.Value()).
		Order("created_at DESC").
		Find(&events).
		Error

	if err != nil {
		return nil, shared.ErrEventBadRequest
	}

	return events, nil
}

func (r *EventPostgresRepository) Delete(ctx context.Context, id *types.EventId) error {
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
