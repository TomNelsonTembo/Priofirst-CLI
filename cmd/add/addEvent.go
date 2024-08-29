package addevents

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"Priofirst-CLI/Internal/Models/sqlite"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

type addEventOptions struct {
}

type app struct {
	events *sqlite.EventsModel
}

func NewAddCmd() *cobra.Command {
	ops := &addEventOptions{}
	addCmd := &cobra.Command{
		Use:   "AddEvent",
		Short: "Add Events related utilities",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ops.run()
		},
	}

	return addCmd

}

func connectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./Internal/app.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func (o *addEventOptions) run() error {
	fmt.Println("Connecting to the database...")
	db := connectDB()
	defer db.Close()
	fmt.Println("Connected to the database....")

	app := app{
		events: &sqlite.EventsModel{DB: db},
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("[+] Enter the title of the event: ")
	eventTitle, _ := reader.ReadString('\n')
	eventTitle = strings.TrimSpace(eventTitle)

	var eventPriorityInt int
	fmt.Println("[+] Enter the priority of the event: ")
	_, err := fmt.Fscanln(reader, &eventPriorityInt)
	if err != nil {
		return fmt.Errorf("failed to read event priority: %w", err)
	}

	var eventAlertTypeInt int
	fmt.Println("[+] Enter the alert type of the event: ")
	_, err = fmt.Fscanln(reader, &eventAlertTypeInt)
	if err != nil {
		return fmt.Errorf("failed to read event alert type: %w", err)
	}

	fmt.Println("[+] Enter the due date of the event (in YYYY-MM-DD format): ")
	eventDueDate, _ := reader.ReadString('\n')
	dueDate, err := time.Parse("2006-01-02", strings.TrimSpace(eventDueDate))
	if err != nil {
		return fmt.Errorf("invalid date format: %w", err)
	}

	fmt.Println("[+] Enter the due time of the event (in HH:MM format): ")
	eventDueTime, _ := reader.ReadString('\n')
	dueTime, err := time.Parse("15:04", strings.TrimSpace(eventDueTime))
	if err != nil {
		return fmt.Errorf("invalid time format: %w", err)
	}

	fmt.Println("[+] Enter the note of the event: ")
	eventNote, _ := reader.ReadString('\n')
	eventNote = strings.TrimSpace(eventNote)

	err = app.events.InsertEvents(eventTitle, time.Now(), eventPriorityInt, eventAlertTypeInt, dueDate, dueTime, eventNote)
	if err != nil {
		return fmt.Errorf("failed to insert event: %w", err)
	}

	fmt.Println("[+] Event inserted successfully!")
	return nil
}
