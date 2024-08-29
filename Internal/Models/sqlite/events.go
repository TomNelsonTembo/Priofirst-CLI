package sqlite

import (
	"database/sql"
	"fmt"
	"time"
)

type EventsModel struct {
	DB *sql.DB
}

func (m *EventsModel) InsertEvents(title string, alertType int, priority int, dueDate time.Time, dueTime time.Time, note string) error {
	fmt.Println("[+] Inserting events...")
	stmt := "INSERT INTO events (title, creation_date, alert_type, priority, due_date, due_time, note) VALUES (?, ?, ?, ?, ?, ?, ?)"

	_, err := m.DB.Exec(stmt, title, time.Now(), alertType, priority, dueDate, dueTime, note)
	if err != nil {
		return err
	}
	return nil
}
