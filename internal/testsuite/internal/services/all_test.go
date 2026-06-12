package services_test

import (
	"testing"

	"govent/internal/domain/services"

	"github.com/stretchr/testify/assert"
)

func TestCanGetAllResources(t *testing.T) {
	var service = services.NewGoldenAllService(repository)
	golden, err := service.Do()

	assert.Nil(t, err)
	assert.IsType(t, services.GoldenAllResponse{}, *golden)
	assert.True(t, repository.AllHaveBeenCalled())
}
