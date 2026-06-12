package services_test

import (
	"os"
	"testing"

	"govent/internal/domain/types"
	internal_test "govent/internal/testsuite/internal"
)

type MockSpyGoldenRepository struct {
	LastCreatedGoldenName     *string
	LastDeleteGoldenId        *string
	LastOneGoldenId           *string
	LastUpdatedGoldenId       *string
	LastUpdatedGoldenName     *string
	LastOneForUpdatedGoldenId *string
	LastAllHaveBeenCalled     bool
	LastUpdateHaveBeenCalled  bool
	LastSearchHaveBeenCalled  bool
}

func NewMockSpyGoldenRepository() *MockSpyGoldenRepository {
	return &MockSpyGoldenRepository{
		LastCreatedGoldenName:     nil,
		LastDeleteGoldenId:        nil,
		LastOneGoldenId:           nil,
		LastUpdatedGoldenId:       nil,
		LastUpdatedGoldenName:     nil,
		LastOneForUpdatedGoldenId: nil,
		LastAllHaveBeenCalled:     false,
		LastUpdateHaveBeenCalled:  false,
		LastSearchHaveBeenCalled:  false,
	}
}

func (m *MockSpyGoldenRepository) Create(golden *types.Golden) error {
	m.LastCreatedGoldenName = &golden.Name

	return nil
}

func (m *MockSpyGoldenRepository) CreateHaveBeenCalledWith(goldenName *string) bool {
	var result = m.LastCreatedGoldenName != nil && *m.LastCreatedGoldenName == *goldenName

	m.LastCreatedGoldenName = nil

	return result
}

func (m *MockSpyGoldenRepository) Delete(id *types.GoldenId) error {
	value := id.Value()
	m.LastDeleteGoldenId = &value

	return nil
}

func (m *MockSpyGoldenRepository) DeleteHaveBeenCalledWith(goldenId *string) bool {
	var result = m.LastDeleteGoldenId != nil && *m.LastDeleteGoldenId == *goldenId

	m.LastDeleteGoldenId = nil

	return result
}

func (m *MockSpyGoldenRepository) Update(golden *types.Golden) error {
	m.LastUpdateHaveBeenCalled = true
	m.LastUpdatedGoldenId = &golden.Id
	m.LastUpdatedGoldenName = &golden.Name
	m.LastOneForUpdatedGoldenId = &golden.Id

	return nil
}

func (m *MockSpyGoldenRepository) UpdateHaveBeenCalledWith(id, name string) bool {
	var matchCalled = m.LastUpdateHaveBeenCalled
	var matchId = *m.LastUpdatedGoldenId == id
	var matchName = *m.LastUpdatedGoldenName == name

	m.LastUpdatedGoldenId = nil
	m.LastUpdatedGoldenName = nil
	m.LastUpdateHaveBeenCalled = false

	return matchCalled && matchId && matchName
}

func (m *MockSpyGoldenRepository) UpdateHaveBeenCalled() bool {
	var matchCalled = m.LastUpdateHaveBeenCalled

	m.LastUpdateHaveBeenCalled = false
	m.LastUpdatedGoldenId = nil
	m.LastUpdatedGoldenName = nil

	return matchCalled
}

func (m *MockSpyGoldenRepository) UpdateHaveBeenCalledOneWith(id string) bool {
	var matchId = *m.LastOneForUpdatedGoldenId == id

	m.LastOneForUpdatedGoldenId = nil

	return matchId
}

func (m *MockSpyGoldenRepository) One(id *types.GoldenId) (*types.Golden, error) {
	value := id.Value()
	m.LastOneGoldenId = &value

	return internal_test.NewRandomGoldenWithCustomId(id), nil
}

func (m *MockSpyGoldenRepository) OneHaveBeenCalledWith(goldenId *string) bool {
	var result = m.LastOneGoldenId != nil && *m.LastOneGoldenId == *goldenId

	m.LastOneGoldenId = nil

	return result
}

func (m *MockSpyGoldenRepository) All() ([]*types.Golden, error) {
	m.LastAllHaveBeenCalled = true

	return nil, nil
}

func (m *MockSpyGoldenRepository) AllHaveBeenCalled() bool {
	result := m.LastAllHaveBeenCalled
	m.LastAllHaveBeenCalled = false

	return result
}

func (m *MockSpyGoldenRepository) SearchAndPaginate(
	searchTerm string,
	pageNumber int,
	pageSize int) ([]*types.Golden, error) {
	m.LastSearchHaveBeenCalled = true

	return nil, nil
}

func (m *MockSpyGoldenRepository) SearchHaveBeenCalled() bool {
	result := m.LastSearchHaveBeenCalled

	m.LastSearchHaveBeenCalled = false

	return result
}

var repository *MockSpyGoldenRepository

func TestMain(m *testing.M) {
	repository = NewMockSpyGoldenRepository()

	os.Exit(m.Run())
}
