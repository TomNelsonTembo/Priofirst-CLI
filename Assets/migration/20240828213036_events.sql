-- +goose Up
-- +goose StatementBegin
CREATE TABLE "events" (
    "event_id" INTEGER NOT NULL UNIQUE,
    "title" VARCHAR(225) NOT NULL,
    "creation_date" DATE NOT NULL,
    "priority" INTEGER,
    "alert_type" INTEGER,
    "due_date" DATE NOT NULL,
    "due_time" TIME NOT NULL,
    "note" TEXT,
    PRIMARY KEY("event_id" AUTOINCREMENT)
);
INSERT INTO "events" ("event_id", "title", "creation_date", "priority", "alert_type", "due_date", "due_time", "note")
VALUES (1, 'Project Kickoff Meeting', '2024-08-01', 1, 2, '2024-08-10', '09:00:00', 'Initial meeting to discuss project scope and deliverables.');

INSERT INTO "events" ("event_id", "title", "creation_date", "priority", "alert_type", "due_date", "due_time", "note")
VALUES (2, 'Design Phase Deadline', '2024-08-02', 2, 1, '2024-09-01', '17:00:00', 'Final date for submission of design documents.');

INSERT INTO "events" ("event_id", "title", "creation_date", "priority", "alert_type", "due_date", "due_time", "note")
VALUES (3, 'Team Building Activity', '2024-08-05', 3, 3, '2024-08-15', '14:00:00', 'Outdoor activity to enhance team collaboration.');

INSERT INTO "events" ("event_id", "title", "creation_date", "priority", "alert_type", "due_date", "due_time", "note")
VALUES (4, 'Client Presentation', '2024-08-10', 1, 2, '2024-08-20', '11:00:00', 'Presentation of the project proposal to the client.');

INSERT INTO "events" ("event_id", "title", "creation_date", "priority", "alert_type", "due_date", "due_time", "note")
VALUES (5, 'Code Review', '2024-08-15', 2, 1, '2024-08-25', '10:00:00', 'Review of the code for the first phase of development.');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "events";
-- +goose StatementEnd
