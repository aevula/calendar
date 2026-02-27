package repository

import (
	"context"
	"fmt"

	"github.com/aevula/interview-hustlers-calendar/internal/databases"
	tasksDomain "github.com/aevula/interview-hustlers-calendar/internal/domain/tasks"
	tasksRepo "github.com/aevula/interview-hustlers-calendar/internal/repository/tasks"
)

type TaskRepository interface {
	Create(context.Context, tasksDomain.Task) (tasksDomain.Task, error)
	FindAllByStatus(ctx context.Context, filter tasksDomain.StatusFilter) ([]tasksDomain.Task, error)
	UpdateStatusByIds(ctx context.Context, ids []tasksDomain.TaskID, status tasksDomain.TaskStatus) error
}

type taskRepository struct {
	db databases.DB
}

func NewTaskRepository(db databases.DB) TaskRepository {
	return &taskRepository{db: db}
}

const taskCreateSQL = `
INSERT INTO tasks (
	name, queue_name, payload, start_at, status
) VALUES (
	$1,   $2,         $3,      $4,     $5 
)
RETURNING
	id, name, queue, payload, start_at, status, attempt, created_at, updated_at
`

func (r *taskRepository) Create(ctx context.Context, task tasksDomain.Task) (tasksDomain.Task, error) {
	rTask := tasksRepo.FromDomain(task)
	zero := tasksRepo.Task{}

	row, err := r.db.QueryRow(ctx, taskCreateSQL,
		rTask.Name,
		rTask.Queue,
		rTask.Payload,
		rTask.StartAt,
		rTask.Status,
	)

	if err != nil {
		return zero.ToDomain(), err
	}

	err = row.Scan(
		&zero.ID,
		&zero.Name,
		&zero.Queue,
		&zero.Payload,
		&zero.StartAt,
		&zero.Status,
		&zero.Attempt,
		&zero.CreatedAt,
		&zero.UpdatedAt,
	)

	return zero.ToDomain(), err
}

const findAllByStatusSQL = `
SELECT
	id, name, queue, payload, start_at, status, attempt, created_at, updated_at
FROM tasks
WHERE
	status = $1
`

func (r *taskRepository) FindAllByStatus(ctx context.Context, filter tasksDomain.StatusFilter) ([]tasksDomain.Task, error) {
	rows, err := r.db.Query(ctx, findAllByStatusSQL, filter.Status)
	if err != nil {
		return nil, err
	}

	tasks := make([]tasksDomain.Task, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}

		task := tasksRepo.Task{}
		err = rows.Scan(
			&task.ID,
			&task.Name,
			&task.Queue,
			&task.Payload,
			&task.StartAt,
			&task.Status,
			&task.Attempt,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task.ToDomain())
	}

	return tasks, nil
}

const updateStatusByIdsSQL = `
UPDATE events SET
	status = $2
WHERE
	id IN ($1)
`

func (r *taskRepository) UpdateStatusByIds(ctx context.Context, ids []tasksDomain.TaskID, status tasksDomain.TaskStatus) error {
	return r.db.Exec(ctx, updateStatusByIdsSQL, fmt.Sprint(ids), string(status))
}
