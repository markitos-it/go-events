package domain_test

import (
	"log"
	"testing"

	"go-vents/internal/domain/shared"
	"go-vents/internal/domain/types"

	"github.com/stretchr/testify/assert"
)

func TestCanCreateValidSlug(t *testing.T) {
	validSlugs := []string{
		"valid-slug",
		"ValidSlug",
		"slug.with.dots",
		"slug_with_underscores",
		"a1",
		"a-b-c",
	}
	for _, s := range validSlugs {
		slug, err := types.NewSlug(s)
		assert.NoError(t, err, "Expected valid slug for: %s", s)
		assert.Equal(t, s, slug.Value())
	}

	invalidSlugs := []string{
		"",
		"1slug",
		"-slug",
		".slug",
		"slug-",
		"slug.",
		"slug@name",
	}
	for _, s := range invalidSlugs {
		_, err := types.NewSlug(s)
		assert.Error(t, err, "Expected error for invalid slug: %s", s)
	}
}

func TestCanCreateValidEventPayload(t *testing.T) {
	validPayloads := []string{
		"",
		"{}",
		`{"key": "value"}`,
		`[1, 2, 3]`,
		`"just a string"`,
		"123",
	}

	for _, payload := range validPayloads {
		if _, err := types.NewPayload(payload); err != nil {
			t.Errorf("Expected valid payload, but got error for: %s", payload)
		}
	}

	invalidPayloads := []string{
		`{bad json}`,
		`{"key": "value",}`,
		`[1, 2,, 3]`,
		`"unclosed string`,
	}

	for _, payload := range invalidPayloads {
		if _, err := types.NewPayload(payload); err == nil {
			t.Errorf("Expected invalid payload, but got valid for: %s", payload)
		}
	}
}

func TestCanCreateValidQueueMessage(t *testing.T) {
	id := shared.UUIDv4()
	subscriberName := shared.Slug()
	eventId := shared.UUIDv4()

	qm, err := types.NewQueueMessage(id, subscriberName, eventId)
	assert.Nil(t, err)
	assert.NotNil(t, qm)
	assert.Equal(t, id, qm.Id)
	assert.Equal(t, subscriberName, qm.SubscriberName)
	assert.Equal(t, eventId, qm.EventId)
	assert.Equal(t, types.StatusPending, qm.Status)
	assert.NotZero(t, qm.CreatedAt)
	assert.NotZero(t, qm.UpdatedAt)
}

func TestCantCreateQueueMessageWithEmptyFields(t *testing.T) {
	id := shared.UUIDv4()
	subscriberName := shared.Slug()
	eventId := shared.UUIDv4()

	_, err := types.NewQueueMessage("", subscriberName, eventId)
	assert.NotNil(t, err)

	_, err = types.NewQueueMessage(id, "", eventId)
	assert.NotNil(t, err)

	subs, err := types.NewQueueMessage(id, subscriberName, "")
	log.Println("---------------------------------")
	log.Println("subs", subs)
	log.Println("---------------------------------")

	assert.NotNil(t, err)
}

func TestCanCreateValidSubscription(t *testing.T) {
	id := shared.UUIDv4()
	subscriberName := "sub1"
	eventName := "event1"
	source := "source1"

	sub, err := types.NewSubscription(id, subscriberName, eventName, source)
	assert.Nil(t, err)
	assert.NotNil(t, sub)
	assert.Equal(t, id, sub.Id)
	assert.Equal(t, subscriberName, sub.SubscriberName)
	assert.Equal(t, eventName, sub.EventName)
	assert.Equal(t, source, sub.Source)
}

func TestCantCreateSubscriptionWithInvalidFields(t *testing.T) {
	id := shared.UUIDv4()
	subscriberName := "sub1"
	eventName := "event1"
	source := "source1"

	_, err := types.NewSubscription("invalid-uuid", subscriberName, eventName, source)
	assert.NotNil(t, err)

	_, err = types.NewSubscription(id, "", eventName, source)
	assert.NotNil(t, err)

	_, err = types.NewSubscription(id, subscriberName, "", source)
	assert.NotNil(t, err)

	_, err = types.NewSubscription(id, subscriberName, eventName, "")
	assert.NotNil(t, err)
}
