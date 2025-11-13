package repository

import (
	"context"
	"time"

	"github.com/aevula/interview-hustlers-calendar/internal/databases"
)

type Event struct {
	ID           int
	Title        string
	Description  string
	UserId       int
	StartAt      time.Time
	Duration     time.Duration
	NotifyOffset time.Duration
	NotifiedAt   time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type EventRepository interface {
	Create(ctx context.Context, event Event) (Event, error)
	Update(ctx context.Context, event Event) (Event, error)
	Delete(ctx context.Context, event Event) error
	List(ctx context.Context) ([]Event, error)
}

type eventRepository struct {
	db databases.Db
}

func NewEventRepository(db databases.Db) EventRepository {
	return &eventRepository{db: db}
}

const eventCreateSQL = `
INSERT INTO events (
	title, description, user_id, start_at, duration, notify_offset
) VALUES (
	$1,    $2, 			$3, 	 $4, 	   $5, 		 $6
) RETURNING
	id, title, description, user_id, start_at, duration, notify_offset, notified_at, created_at, updated_at
`

func (r *eventRepository) Create(ctx context.Context, event Event) (Event, error) {
	ev := Event{}

	row, err := r.db.QueryRow(ctx, eventCreateSQL,
		event.Title,
		event.Description,
		event.UserId,
		event.StartAt,
		event.Duration,
		event.NotifyOffset,
	)
	if err != nil {
		return ev, err
	}

	err = row.Scan(
		&ev.ID,
		&ev.Title,
		&ev.Description,
		&ev.UserId,
		&ev.StartAt,
		&ev.Duration,
		&ev.NotifyOffset,
		&ev.NotifiedAt,
		&ev.CreatedAt,
		&ev.UpdatedAt,
	)
	return ev, err
}

const eventUpdateSQL = `
UPDATE events SET (
	title = $2, description = $3, user_id = $4, start_at = $5, duration = $6, notify_offset = $7, notified_at = $8
) WHERE
	id = $1
RETURNING
	id, title, description, user_id, start_at, duration, notify_offset, notified_at, created_at, updated_at
`

func (r *eventRepository) Update(ctx context.Context, event Event) (Event, error) {
	ev := Event{}

	row, err := r.db.QueryRow(ctx, eventUpdateSQL,
		event.ID,
		event.Title,
		event.Description,
		event.UserId,
		event.StartAt,
		event.Duration,
		event.NotifyOffset,
		event.NotifiedAt,
	)
	if err != nil {
		return ev, err
	}

	err = row.Scan(
		&ev.ID,
		&ev.Title,
		&ev.Description,
		&ev.UserId,
		&ev.StartAt,
		&ev.Duration,
		&ev.NotifyOffset,
		&ev.NotifiedAt,
		&ev.CreatedAt,
		&ev.UpdatedAt,
	)
	return ev, err
}

const eventDeleteSQL = `
DELETE FROM events
WHERE
	id = $1
RETURNING
	id
`

func (r *eventRepository) Delete(ctx context.Context, event Event) error {
	row, err := r.db.QueryRow(ctx, eventDeleteSQL,
		event.ID,
	)
	if err != nil {
		return err
	}

	var id int
	err = row.Scan(&id)

	return err
}

const eventListSQL = `
SELECT (
	id, title, description, user_id, start_at, duration, notify_offset, notified_at, created_at, updated_at
) FROM events`

func (r *eventRepository) List(ctx context.Context) ([]Event, error) {
	rows, err := r.db.Query(ctx, eventListSQL)
	if err != nil {
		return nil, err
	}

	events := make([]Event, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}

		event := Event{}
		err = rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.UserId,
			&event.StartAt,
			&event.Duration,
			&event.NotifyOffset,
			&event.NotifiedAt,
			&event.CreatedAt,
			&event.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
