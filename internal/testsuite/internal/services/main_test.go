package services_test

import (
	"context"
	"os"
	"testing"

	"go-vents/internal/domain/types"
	internal_test "go-vents/internal/testsuite/internal"
)

type MockSpyEventRepository struct {
	LastCreatedEventName           *string
	LastDeleteEventId              *string
	LastOneEventId                 *string
	LastAllBySlugAndSource         []types.Event
	LastCreatedSubscriptionEvent   *string
	LastAckMessageId               *string
	LastAckMessagesIds             []string
	LastPullMessagesName           *string
	LastPullMessagesSource         *string
	LastPullMessagesSubscriberName *string
}

func NewMockSpyEventRepository() *MockSpyEventRepository {
	return &MockSpyEventRepository{
		LastCreatedEventName:           nil,
		LastDeleteEventId:              nil,
		LastOneEventId:                 nil,
		LastAllBySlugAndSource:         nil,
		LastCreatedSubscriptionEvent:   nil,
		LastAckMessageId:               nil,
		LastAckMessagesIds:             nil,
		LastPullMessagesName:           nil,
		LastPullMessagesSource:         nil,
		LastPullMessagesSubscriberName: nil,
	}
}

func (m *MockSpyEventRepository) PullMessages(ctx context.Context, subscriberName *types.Slug, slug *types.Slug, source *types.Source) ([]*types.Queue, error) {
	slugValue := slug.Value()
	sourceValue := source.Value()
	subscriberNameValue := subscriberName.Value()

	m.LastPullMessagesName = &slugValue
	m.LastPullMessagesSource = &sourceValue
	m.LastPullMessagesSubscriberName = &subscriberNameValue

	return []*types.Queue{}, nil
}

func (m *MockSpyEventRepository) AckMessage(ctx context.Context, id *types.Id) error {
	v := id.Value()
	m.LastAckMessageId = &v
	return nil
}

func (m *MockSpyEventRepository) AckMessages(ctx context.Context, ids []*types.Id) error {
	var strIds []string
	for _, id := range ids {
		strIds = append(strIds, id.Value())
	}
	m.LastAckMessagesIds = strIds
	return nil
}

func (m *MockSpyEventRepository) Create(ctx context.Context, event *types.Event) error {
	m.LastCreatedEventName = &event.Slug

	return nil
}

func (m *MockSpyEventRepository) CreateHaveBeenCalledWith(eventName *string) bool {
	var result = m.LastCreatedEventName != nil && *m.LastCreatedEventName == *eventName

	m.LastCreatedEventName = nil

	return result
}

func (m *MockSpyEventRepository) CreateSubscription(ctx context.Context, sub *types.Subscription) error {
	m.LastCreatedSubscriptionEvent = &sub.EventName
	return nil
}

func (m *MockSpyEventRepository) CreateSubscriptionHaveBeenCalledWith(eventName *string) bool {
	var result = m.LastCreatedSubscriptionEvent != nil && *m.LastCreatedSubscriptionEvent == *eventName

	m.LastCreatedSubscriptionEvent = nil

	return result
}

func (m *MockSpyEventRepository) Delete(ctx context.Context, id *types.Id) error {
	value := id.Value()
	m.LastDeleteEventId = &value

	return nil
}

func (m *MockSpyEventRepository) DeleteHaveBeenCalledWith(eventId *string) bool {
	var result = m.LastDeleteEventId != nil && *m.LastDeleteEventId == *eventId

	m.LastDeleteEventId = nil

	return result
}

func (m *MockSpyEventRepository) One(ctx context.Context, id *types.Id) (*types.Event, error) {
	value := id.Value()
	m.LastOneEventId = &value

	return internal_test.NewRandomEventWithCustomId(id), nil
}

func (m *MockSpyEventRepository) OneHaveBeenCalledWith(eventId *string) bool {
	var result = m.LastOneEventId != nil && *m.LastOneEventId == *eventId

	m.LastOneEventId = nil

	return result
}

func (m *MockSpyEventRepository) AllBySlugAndSource(ctx context.Context, slug *types.Slug, source *types.Source) ([]*types.Event, error) {
	anEvent := internal_test.NewRandomEventWithSlugAndSource(slug.Value(), source.Value())
	m.LastAllBySlugAndSource = append(m.LastAllBySlugAndSource, *anEvent)

	return []*types.Event{anEvent}, nil
}

func (m *MockSpyEventRepository) LastAllBySlugAndSourceHaveBeenCalled(slug *types.Slug, source *types.Source) bool {
	var result = m.LastAllBySlugAndSource[0].Slug == slug.Value() &&
		m.LastAllBySlugAndSource[0].Source == source.Value()

	m.LastAllBySlugAndSource = nil

	return result
}

var repository *MockSpyEventRepository

func TestMain(m *testing.M) {
	repository = NewMockSpyEventRepository()

	os.Exit(m.Run())
}
