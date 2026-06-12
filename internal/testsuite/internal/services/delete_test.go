package services_test

import (
	"testing"

	"govent/internal/domain/services"
	"govent/internal/domain/shared"

	"github.com/stretchr/testify/assert"
)

func TestCanDeleteAUser(t *testing.T) {
	var request = services.GoldenDeleteRequest{
		Id: shared.UUIDv4(),
	}

	var service = services.NewGoldenDeleteService(repository)
	err := service.Do(request)
	assert.Nil(t, err)
	assert.True(t, repository.DeleteHaveBeenCalledWith(&request.Id))
}

func TestCantDeleteWithoutId(t *testing.T) {
	var request = services.GoldenDeleteRequest{}

	var service = services.NewGoldenDeleteService(repository)
	err := service.Do(request)

	assert.ErrorIs(t, err, shared.ErrGoldenBadRequest)
	assert.NotNil(t, err)
	assert.False(t, repository.DeleteHaveBeenCalledWith(&request.Id))
}

func TestCantDeleteWithInvalidId(t *testing.T) {
	var request = services.GoldenDeleteRequest{
		Id: "invalid-id",
	}

	var service = services.NewGoldenDeleteService(repository)
	err := service.Do(request)

	assert.ErrorIs(t, err, shared.ErrGoldenBadRequest)
	assert.False(t, repository.DeleteHaveBeenCalledWith(&request.Id))
}
