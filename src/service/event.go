package service

import (
	"encoding/json"
	"io"
	"testProj/src/models"
	"testProj/src/store"
)

func HandleEvent(req io.ReadCloser) error {
	// обработка входных данных
	event := &models.Event{}
	err := json.NewDecoder(req).Decode(&event)
	if err != nil {
		return err
	}
	// валидация
	err = event.Validate()
	if err != nil {
		return err
	}
	// записываем событие в базу
	return store.Insert(event)
}
