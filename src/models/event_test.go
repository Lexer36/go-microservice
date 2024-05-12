package models

import (
	"fmt"
	"testing"
)

func TestEvent_Validate(t *testing.T) {
	t.Parallel()
	// Тестовые случаи
	testCases := []struct {
		name  string
		input Event
		err   error
	}{
		{
			name:  "OK",
			input: Event{EventType: "login", EventTime: "2006-01-02 15:04:05"},
			err:   nil,
		},
		{
			name:  "TypeErr",
			input: Event{EventType: "logout", EventTime: "2006-01-02 15:04:05"},
			err:   fmt.Errorf("неверный тип события eventType"),
		},
		{
			name:  "TimeFormatErr",
			input: Event{EventType: "login", EventTime: "01-02-2006 15:04:05"},
			err:   fmt.Errorf("ошибка формата даты eventTime"),
		},
	}

	// Итерация по тестовым случаям
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actualErr := tc.input.Validate()
			if actualErr.Error() != tc.err.Error() {
				t.Fatalf("expected error %t, got %s", tc.err, actualErr)
			}
		})
	}
}
