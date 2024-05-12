package models

import (
	"fmt"
	"testProj/src/utils"
)

var (
	EventTypes = map[string]int{
		"login": 1,
	}
)

type Event struct {
	EventID   int64
	EventType string `json:"eventType"`
	UserID    int64  `json:"userId"`
	EventTime string `json:"eventTime"`
	Payload   string `json:"payload"`
}

// валидация данных
func (e *Event) Validate() error {
	// ищем из справочника EventId по eventType (?)
	val, exists := EventTypes[e.EventType]
	if exists {
		e.EventID = int64(val)
	} else {
		return fmt.Errorf("неверный тип события eventType")
	}

	// валидируем формат даты
	if utils.ValidateTimeFormat(e.EventTime) != nil {
		return fmt.Errorf("ошибка формата даты eventTime")
	}

	return nil
}
