package store

import (
	databases "testProj/src/db"
	"testProj/src/models"
	"time"
)

// вставка данных в таблицу events
func Insert(e *models.Event) error {
	query := "INSERT INTO events VALUES (?, ?, ?, ?, ?);"
	_, err := databases.DbSession().Exec(query, e.EventID, e.EventType, e.UserID, e.EventTime, e.Payload)
	return err
}

// Вывод событий по заданному eventType и временному диапазону
func SelectByTypeAndTime(eventType string, timeStart, timeEnd time.Time) ([]*models.Event, error) {
	query := `SELECT *
			FROM events
			WHERE eventType = ? AND eventTime BETWEEN ? AND ?`
	rows, err := databases.DbSession().Query(query, eventType, timeStart, timeEnd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*models.Event
	for rows.Next() {
		event := &models.Event{}
		if err := rows.Scan(&event.EventID, &event.EventType, &event.UserID, &event.EventTime, &event.Payload); err == nil {
			result = append(result, event)
		}
	}
	return result, nil
}
