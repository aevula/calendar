package repository

import (
	"context"

	"github.com/aevula/interview-hustlers-calendar/internal/databases"
	eventsDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/events"
	eventsRepo "github.com/aevula/interview-hustlers-calendar/internal/repository/events"
)

type EventRepository interface {
	FindById(ctx context.Context, id eventsDomain.EventID) (eventsDomain.Event, error)
	Create(ctx context.Context, event eventsDomain.Event) (eventsDomain.Event, error)
	Update(ctx context.Context, event eventsDomain.Event) (eventsDomain.Event, error)
	Delete(ctx context.Context, id eventsDomain.EventID) error
	FindAllByStartAt(ctx context.Context, filter eventsDomain.StartAtFilter) ([]eventsDomain.Event, error)
	DeleteAllByNotifiedAt(ctx context.Context, filter eventsDomain.NotifiedAtFilter) ([]eventsDomain.Event, error)
	FindAllNotNotifiedByNotifyAt(ctx context.Context, filter eventsDomain.NotifyAtFilter) ([]eventsDomain.Event, error)
}

type eventRepository struct {
	db databases.DB
}

func NewEventRepository(db databases.DB) EventRepository {
	return &eventRepository{db: db}
}

const findByIdSQL = `
SELECT
	id, title, description, user_id, start_at, duration, notify_at, notified_at, created_at, updated_at
FROM events
WHERE
	id = $1
`

func (r *eventRepository) FindById(ctx context.Context, id eventsDomain.EventID) (eventsDomain.Event, error) {
	zero := eventsRepo.Event{}

	row, err := r.db.QueryRow(ctx, findByIdSQL, int(id))
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

const createSQL = `
INSERT INTO events (
	title, description, user_id, start_at, duration, notify_at
) VALUES (
	$1,    $2, 			$3, 	 $4, 	   $5, 		 $6
) RETURNING
	id, title, description, user_id, start_at, duration, notify_at, notified_at, created_at, updated_at
`

func (r *eventRepository) Create(ctx context.Context, event eventsDomain.Event) (eventsDomain.Event, error) {
	rEvent := eventsRepo.FromDomain(event)
	zero := eventsRepo.Event{}

	row, err := r.db.QueryRow(ctx, createSQL,
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

const updateByIdSQL = `
UPDATE events SET
	title = $2, description = $3, start_at = $4, duration = $5, notify_at = $6, notified_at = $7
WHERE
	id = $1
RETURNING
	id, title, description, user_id, start_at, duration, notify_at, notified_at, created_at, updated_at
`

func (r *eventRepository) Update(ctx context.Context, event eventsDomain.Event) (eventsDomain.Event, error) {
	rEvent := eventsRepo.FromDomain(event)
	zero := eventsRepo.Event{}

	row, err := r.db.QueryRow(ctx, updateByIdSQL,
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

const deleteByIdSQL = `
DELETE FROM events
WHERE
	id = $1
RETURNING
	id
`

func (r *eventRepository) Delete(ctx context.Context, id eventsDomain.EventID) error {
	row, err := r.db.QueryRow(ctx, deleteByIdSQL, int(id))
	if err != nil {
		return err
	}

	var zero int
	err = row.Scan(&zero)
	return err
}

const findAllByStartAtSQL = `
SELECT
	id, title, description, user_id, start_at, duration, notify_at, notified_at, created_at, updated_at
FROM events
WHERE
	start_at >= $1 AND start_at <= $2
`

func (r *eventRepository) FindAllByStartAt(ctx context.Context, filter eventsDomain.StartAtFilter) ([]eventsDomain.Event, error) {
	rows, err := r.db.Query(ctx, findAllByStartAtSQL, filter.GTE, filter.LTE)
	if err != nil {
		return nil, err
	}

	events := make([]eventsDomain.Event, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}

		event := eventsRepo.Event{}
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

const findAllByNotifiedAtSQL = `
DELETE FROM events
WHERE
	notified_at <= $1
RETURNING
	id
`

func (r *eventRepository) DeleteAllByNotifiedAt(ctx context.Context, filter eventsDomain.NotifiedAtFilter) ([]eventsDomain.Event, error) {
	rows, err := r.db.Query(ctx, findAllByNotifiedAtSQL, filter.LTE)
	if err != nil {
		return nil, err
	}

	events := make([]eventsDomain.Event, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}

		event := eventsRepo.Event{}
		err = rows.Scan(
			&event.ID,
		)
		if err != nil {
			return nil, err
		}

		events = append(events, event.ToDomain())
	}

	return events, nil
}

const findAllNotNotifiedByNotifyAtSQL = `
SELECT
	id, title, description, user_id, start_at, duration
FROM events
WHERE
	notified_at is NULL AND notify_at <= $1
`

func (r *eventRepository) FindAllNotNotifiedByNotifyAt(ctx context.Context, filter eventsDomain.NotifyAtFilter) ([]eventsDomain.Event, error) {
	rows, err := r.db.Query(ctx, findAllNotNotifiedByNotifyAtSQL, filter.LTE)
	if err != nil {
		return nil, err
	}

	events := make([]eventsDomain.Event, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}

		event := eventsRepo.Event{}
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
