package services_test

import (
	"testing"

	"govent/internal/domain/services"

	"github.com/stretchr/testify/assert"
)

func TestCanSearchResources(t *testing.T) {
	var service = services.NewGoldensearchService(repository)
	golden, err := service.Do(services.GoldensearchRequest{})

	assert.Nil(t, err)
	assert.IsType(t, services.GoldensearchResponse{}, *golden)
	assert.True(t, repository.SearchHaveBeenCalled())
}
