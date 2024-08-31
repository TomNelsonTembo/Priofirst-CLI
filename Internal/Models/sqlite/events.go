package sqlite

import (
	"database/sql"
	"fmt"
	"time"
)

type EventsModel struct {
	DB *sql.DB
}

func (m *EventsModel) InsertEvents(title string, creationDate time.Time, priority int, alertType int, dueDate time.Time, dueTime time.Time, note string) error {
	fmt.Println("[+] Inserting events...")
	stmt := `INSERT INTO "events" ("title", "creation_date", "priority", "alert_type", "due_date", "due_time", "note")
VALUES (?, ?, ?, ?, ?, ?, ?);`
	result, err := m.DB.Exec(stmt, title, creationDate, priority, alertType, dueDate, dueTime, note)
	if err != nil {
		return err
	}
	eventID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Println("Generated event ID:", eventID)

	return nil
}
