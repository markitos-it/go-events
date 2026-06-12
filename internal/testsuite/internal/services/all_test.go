package services_test

import (
	"context"
	"log"
	"testing"

	"govent/internal/domain/services"
	"govent/internal/domain/shared"

	"github.com/stretchr/testify/assert"
)

func TestCanGetAllResources(t *testing.T) {
	var service = services.NewEventAllService(repository)

	anyName := shared.RandomString()
	anySource := shared.RandomString()
	event, err := service.Do(context.TODO(), services.EventAllRequest{
		Name:   anyName,
		Source: anySource,
	})

	assert.Nil(t, err)
	assert.IsType(t, services.EventAllResponse{}, *event)

	log.Println("----------------------------")
	log.Println("event.Data[0].GetName()", event.Data[0].GetName())
	log.Println("anyName", anyName)
	log.Println("anySource", anySource)
	log.Println("event.Data[0].GetSource()", event.Data[0].GetSource())
	log.Println("----------------------------")

	assert.True(t, repository.LastAllByNameAndSourceHaveBeenCalled(event.Data[0].GetName(), event.Data[0].GetSource()))
}
