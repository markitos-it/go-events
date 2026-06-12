package services_test

import (
	"testing"

	"govent/internal/domain/services"
	"govent/internal/domain/shared"
	"govent/internal/domain/types"

	"github.com/stretchr/testify/assert"
)

func TestCanCreateAUser(t *testing.T) {
	var golden = types.Golden{
		Name: shared.RandomPersonalName(),
	}
	var request = services.GoldenCreateRequest{
		Name: golden.Name,
	}

	var service = services.NewGoldenCreateService(repository)
	response, err := service.Do(request)

	assert.Nil(t, err)
	assert.True(t, repository.CreateHaveBeenCalledWith(&request.Name))
	assert.Equal(t, response.Name, request.Name)
	assert.NotEmpty(t, response.Id)
}

func TestCantCreateWithoutName(t *testing.T) {
	var request = services.GoldenCreateRequest{}

	var service = services.NewGoldenCreateService(repository)
	_, err := service.Do(request)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, shared.ErrGoldenBadRequest)
	assert.False(t, repository.CreateHaveBeenCalledWith(&request.Name))
}

func TestCantCreateWithoutValidName(t *testing.T) {
	var request = services.GoldenCreateRequest{
		Name: "",
	}

	var service = services.NewGoldenCreateService(repository)
	_, err := service.Do(request)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, shared.ErrGoldenBadRequest)
	assert.False(t, repository.CreateHaveBeenCalledWith(&request.Name))
}
