package repository

import (
	"context"
	"time"

	"github.com/aevula/interview-hustlers-calendar/internal/databases"
	domain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
	repo "github.com/aevula/interview-hustlers-calendar/internal/repository/events"
)

type EventRepository interface {
	Create(ctx context.Context, event domain.Event) (domain.Event, error)
	Update(ctx context.Context, event domain.Event) (domain.Event, error)
	Delete(ctx context.Context, id domain.EventID) error
	All(ctx context.Context) ([]domain.Event, error)
	AllNotifyable(ctx context.Context, from time.Time) ([]domain.Event, error)
}

type eventRepository struct {
	db databases.Db
}

func NewEventRepository(db databases.Db) EventRepository {
	return &eventRepository{db: db}
}

const eventCreateSQL = `
INSERT INTO events (
	title, description, user_id, start_at, duration, notify_at
) VALUES (
	$1,    $2, 			$3, 	 $4, 	   $5, 		 $6
) RETURNING
	id, title, description, user_id, start_at, duration, notify_at, notified_at, created_at, updated_at
`

func (r *eventRepository) Create(ctx context.Context, event domain.Event) (domain.Event, error) {
	rEvent := repo.FromDomain(event)
	zero := repo.Event{}

	row, err := r.db.QueryRow(ctx, eventCreateSQL,
		rEvent.Title,
		rEvent.Description,
		rEvent.UserId,
		rEvent.StartAt,
		rEvent.Duration,
		rEvent.NotifyAt,
	)
	if err != nil {
		return zero.ToDomain(), err
	}

	err = row.Scan(
		&zero.ID,
		&zero.Title,
		&zero.Description,
		&zero.UserId,
		&zero.StartAt,
		&zero.Duration,
		&zero.NotifyAt,
		&zero.NotifiedAt,
		&zero.CreatedAt,
		&zero.UpdatedAt,
	)
	return zero.ToDomain(), err
}

const eventUpdateSQL = `
UPDATE events SET
	title = $2, description = $3, start_at = $4, duration = $5, notify_at = $6, notified_at = $7
WHERE
	id = $1
RETURNING
	id, title, description, user_id, start_at, duration, notify_at, notified_at, created_at, updated_at
`

func (r *eventRepository) Update(ctx context.Context, event domain.Event) (domain.Event, error) {
	rEvent := repo.FromDomain(event)
	zero := repo.Event{}

	row, err := r.db.QueryRow(ctx, eventUpdateSQL,
		rEvent.ID,
		rEvent.Title,
		rEvent.Description,
		rEvent.StartAt,
		rEvent.Duration,
		rEvent.NotifyAt,
		rEvent.NotifiedAt,
	)
	if err != nil {
		return zero.ToDomain(), err
	}

	err = row.Scan(
		&zero.ID,
		&zero.Title,
		&zero.Description,
		&zero.UserId,
		&zero.StartAt,
		&zero.Duration,
		&zero.NotifyAt,
		&zero.NotifiedAt,
		&zero.CreatedAt,
		&zero.UpdatedAt,
	)
	return zero.ToDomain(), err
}

const eventDeleteSQL = `
DELETE FROM events
WHERE
	id = $1
RETURNING
	id
`

func (r *eventRepository) Delete(ctx context.Context, id domain.EventID) error {
	row, err := r.db.QueryRow(ctx, eventDeleteSQL, int(id))
	if err != nil {
		return err
	}

	var zero int
	err = row.Scan(&zero)
	return err
}

const eventAllSQL = `
SELECT
	id, title, description, user_id, start_at, duration, notify_at, notified_at, created_at, updated_at
FROM events
`

func (r *eventRepository) All(ctx context.Context) ([]domain.Event, error) {
	rows, err := r.db.Query(ctx, eventAllSQL)
	if err != nil {
		return nil, err
	}

	events := make([]domain.Event, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}

		event := repo.Event{}
		err = rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.UserId,
			&event.StartAt,
			&event.Duration,
			&event.NotifyAt,
			&event.NotifiedAt,
			&event.CreatedAt,
			&event.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		events = append(events, event.ToDomain())
	}

	return events, nil
}

const eventNotifyableSQL = `
SELECT
	id, title, description, user_id, start_at, duration
FROM events
WHERE notified_at is NULL AND notify_at <= $1
`

func (r *eventRepository) AllNotifyable(ctx context.Context, from time.Time) ([]domain.Event, error) {
	rows, err := r.db.Query(ctx, eventNotifyableSQL, from)
	if err != nil {
		return nil, err
	}

	events := make([]domain.Event, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}

		event := repo.Event{}
		err = rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.UserId,
			&event.StartAt,
			&event.Duration,
		)
		if err != nil {
			return nil, err
		}

		events = append(events, event.ToDomain())
	}

	return events, nil
}
