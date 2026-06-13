package services_test

import (
	"context"
	"testing"

	"go-vents/internal/domain/services"
	"go-vents/internal/domain/shared"

	"github.com/stretchr/testify/assert"
)

func TestCanGetAllResources(t *testing.T) {
	var service = services.NewEventAllService(repository)

	anyName := shared.RandomString()
	anySource := shared.RandomString()
	event, err := service.Do(context.TODO(), services.EventAllRequest{
		Slug:   anyName,
		Source: anySource,
	})

	assert.Nil(t, err)
	assert.IsType(t, services.EventAllResponse{}, *event)
	assert.True(t, repository.LastAllBySlugAndSourceHaveBeenCalled(event.Data[0].GetSlug(), event.Data[0].GetSource()))
}
