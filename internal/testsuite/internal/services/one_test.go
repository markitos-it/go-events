package services_test

import (
	"testing"

	"govent/internal/domain/services"
	"govent/internal/domain/shared"

	"github.com/stretchr/testify/assert"
)

func TestCanGetAResource(t *testing.T) {
	var request = services.GoldenOneRequest{
		Id: shared.UUIDv4(),
	}

	var service = services.NewGoldenOneService(repository)
	golden, err := service.Do(request)

	assert.Nil(t, err)
	assert.IsType(t, services.GoldenOneResponse{}, *golden)
	assert.True(t, repository.OneHaveBeenCalledWith(&request.Id))
}

func TestCantGetWithoutId(t *testing.T) {
	var request = services.GoldenOneRequest{}

	var service = services.NewGoldenOneService(repository)
	_, err := service.Do(request)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, shared.ErrGoldenBadRequest)
	assert.False(t, repository.OneHaveBeenCalledWith(&request.Id))
}

func TestCantGetWithoutInvalidId(t *testing.T) {
	var request = services.GoldenOneRequest{
		Id: "invalid-id",
	}

	var service = services.NewGoldenOneService(repository)
	_, err := service.Do(request)

	assert.ErrorIs(t, err, shared.ErrGoldenBadRequest)
	assert.False(t, repository.OneHaveBeenCalledWith(&request.Id))
}
