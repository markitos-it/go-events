package services_test

import (
	"context"
	"testing"

	"go-vents/internal/domain/services"
	"go-vents/internal/domain/shared"
	"go-vents/internal/domain/types"

	"github.com/stretchr/testify/assert"
)

func TestCanCreateAUser(t *testing.T) {
	var event = types.Event{
		Slug: shared.RandomSlug(),
	}
	var request = services.EventCreateRequest{
		Slug: event.Slug,
	}

	var service = services.NewEventCreateService(repository)
	response, err := service.Do(context.TODO(), request)

	assert.Nil(t, err)
	assert.True(t, repository.CreateHaveBeenCalledWith(&request.Slug))
	assert.Equal(t, response.Slug, request.Slug)
	assert.NotEmpty(t, response.Id)
}

func TestCantCreateWithoutName(t *testing.T) {
	var request = services.EventCreateRequest{}

	var service = services.NewEventCreateService(repository)
	_, err := service.Do(context.TODO(), request)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, shared.ErrEventBadRequest)
	assert.False(t, repository.CreateHaveBeenCalledWith(&request.Slug))
}

func TestCantCreateWithoutValidName(t *testing.T) {
	var request = services.EventCreateRequest{
		Slug: "",
	}

	var service = services.NewEventCreateService(repository)
	_, err := service.Do(context.TODO(), request)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, shared.ErrEventBadRequest)
	assert.False(t, repository.CreateHaveBeenCalledWith(&request.Slug))
}
