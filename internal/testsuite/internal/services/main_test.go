package services_test

import (
	"os"
	"testing"

	"govent/internal/domain/types"
	internal_test "govent/internal/testsuite/internal"
)

type MockSpyEventRepository struct {
	LastCreatedEventName     *string
	LastDeleteEventId        *string
	LastOneEventId           *string
	LastUpdatedEventId       *string
	LastUpdatedEventName     *string
	LastOneForUpdatedEventId *string
	LastAllHaveBeenCalled    bool
	LastUpdateHaveBeenCalled bool
	LastSearchHaveBeenCalled bool
}

func NewMockSpyEventRepository() *MockSpyEventRepository {
	return &MockSpyEventRepository{
		LastCreatedEventName:     nil,
		LastDeleteEventId:        nil,
		LastOneEventId:           nil,
		LastUpdatedEventId:       nil,
		LastUpdatedEventName:     nil,
		LastOneForUpdatedEventId: nil,
		LastAllHaveBeenCalled:    false,
		LastUpdateHaveBeenCalled: false,
		LastSearchHaveBeenCalled: false,
	}
}

func (m *MockSpyEventRepository) Create(event *types.Event) error {
	m.LastCreatedEventName = &event.Name

	return nil
}

func (m *MockSpyEventRepository) CreateHaveBeenCalledWith(eventName *string) bool {
	var result = m.LastCreatedEventName != nil && *m.LastCreatedEventName == *eventName

	m.LastCreatedEventName = nil

	return result
}

func (m *MockSpyEventRepository) Delete(id *types.EventId) error {
	value := id.Value()
	m.LastDeleteEventId = &value

	return nil
}

func (m *MockSpyEventRepository) DeleteHaveBeenCalledWith(eventId *string) bool {
	var result = m.LastDeleteEventId != nil && *m.LastDeleteEventId == *eventId

	m.LastDeleteEventId = nil

	return result
}

func (m *MockSpyEventRepository) Update(event *types.Event) error {
	m.LastUpdateHaveBeenCalled = true
	m.LastUpdatedEventId = &event.Id
	m.LastUpdatedEventName = &event.Name
	m.LastOneForUpdatedEventId = &event.Id

	return nil
}

func (m *MockSpyEventRepository) UpdateHaveBeenCalledWith(id, name string) bool {
	var matchCalled = m.LastUpdateHaveBeenCalled
	var matchId = *m.LastUpdatedEventId == id
	var matchName = *m.LastUpdatedEventName == name

	m.LastUpdatedEventId = nil
	m.LastUpdatedEventName = nil
	m.LastUpdateHaveBeenCalled = false

	return matchCalled && matchId && matchName
}

func (m *MockSpyEventRepository) UpdateHaveBeenCalled() bool {
	var matchCalled = m.LastUpdateHaveBeenCalled

	m.LastUpdateHaveBeenCalled = false
	m.LastUpdatedEventId = nil
	m.LastUpdatedEventName = nil

	return matchCalled
}

func (m *MockSpyEventRepository) UpdateHaveBeenCalledOneWith(id string) bool {
	var matchId = *m.LastOneForUpdatedEventId == id

	m.LastOneForUpdatedEventId = nil

	return matchId
}

func (m *MockSpyEventRepository) One(id *types.EventId) (*types.Event, error) {
	value := id.Value()
	m.LastOneEventId = &value

	return internal_test.NewRandomEventWithCustomId(id), nil
}

func (m *MockSpyEventRepository) OneHaveBeenCalledWith(eventId *string) bool {
	var result = m.LastOneEventId != nil && *m.LastOneEventId == *eventId

	m.LastOneEventId = nil

	return result
}

func (m *MockSpyEventRepository) All() ([]*types.Event, error) {
	m.LastAllHaveBeenCalled = true

	return nil, nil
}

func (m *MockSpyEventRepository) AllHaveBeenCalled() bool {
	result := m.LastAllHaveBeenCalled
	m.LastAllHaveBeenCalled = false

	return result
}

func (m *MockSpyEventRepository) SearchAndPaginate(
	searchTerm string,
	pageNumber int,
	pageSize int) ([]*types.Event, error) {
	m.LastSearchHaveBeenCalled = true

	return nil, nil
}

func (m *MockSpyEventRepository) SearchHaveBeenCalled() bool {
	result := m.LastSearchHaveBeenCalled

	m.LastSearchHaveBeenCalled = false

	return result
}

var repository *MockSpyEventRepository

func TestMain(m *testing.M) {
	repository = NewMockSpyEventRepository()

	os.Exit(m.Run())
}
