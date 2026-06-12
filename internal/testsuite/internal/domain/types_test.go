package domain_test

import (
	"strings"
	"testing"

	"govent/internal/domain/types"
)

func TestCanCreateValidEventName(t *testing.T) {
	validNames := []string{
		"ValidName",
		"AnotherValidName",
		"Valid Name With Spaces",
		"Short",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"}
	for _, name := range validNames {
		if _, err := types.NewEventName(name); err != nil {
			t.Errorf("Expected valid name, but got invalid: %s", name)
		}
	}

	invalidNames := []string{
		" InvalidName",
		"InvalidName ",
		"Invalid Name ",
		" Invalid Name",
		"Invalid@Name",
		"Invalid#Name",
		"Invalid123Name",
		"Invalid Name With Spaces ",
		" Invalid Name With Spaces",
		"Invalid Name With Spaces And Symbols!",
	}
	for _, name := range invalidNames {
		if _, err := types.NewEventName(name); err == nil {
			t.Errorf("Expected valid name, but got invalid: %s", name)
		}
	}

	invalidLengthNames := []string{
		strings.Repeat("a", types.EVENT_NAME_MAX_LENGTH+1),
		strings.Repeat("b", types.EVENT_NAME_MIN_LENGTH-1),
		"",
	}
	for _, name := range invalidLengthNames {
		if _, err := types.NewEventName(name); err == nil {
			t.Errorf("Expected invalid name, but got invalid: %s", name)
		}
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
		if _, err := types.NewEventPayload(payload); err != nil {
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
		if _, err := types.NewEventPayload(payload); err == nil {
			t.Errorf("Expected invalid payload, but got valid for: %s", payload)
		}
	}
}
