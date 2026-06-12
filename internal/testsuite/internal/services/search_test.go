package services_test

import (
	"testing"

	"govent/internal/domain/services"

	"github.com/stretchr/testify/assert"
)

func TestCanSearchResources(t *testing.T) {
	var service = services.NewEventsearchService(repository)
	event, err := service.Do(services.EventsearchRequest{})

	assert.Nil(t, err)
	assert.IsType(t, services.EventsearchResponse{}, *event)
	assert.True(t, repository.SearchHaveBeenCalled())
}
