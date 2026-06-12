package services_test

import (
	"testing"

	"govent/internal/domain/services"

	"github.com/stretchr/testify/assert"
)

func TestCanGetAllResources(t *testing.T) {
	var service = services.NewEventAllService(repository)
	event, err := service.Do()

	assert.Nil(t, err)
	assert.IsType(t, services.EventAllResponse{}, *event)
	assert.True(t, repository.AllHaveBeenCalled())
}
