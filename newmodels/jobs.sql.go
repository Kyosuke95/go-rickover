// Code generated by sqlc. DO NOT EDIT.
// source: jobs.sql

package newmodels

import (
	"context"
)

const createJob = `-- name: CreateJob :one
INSERT INTO jobs (
    name,
    delivery_strategy,
    attempts,
    concurrency)
VALUES ($1, $2, $3, $4)
RETURNING name, delivery_strategy, attempts, concurrency, created_at, auto_id
`

type CreateJobParams struct {
	Name             string           `json:"name"`
	DeliveryStrategy DeliveryStrategy `json:"delivery_strategy"`
	Attempts         int16            `json:"attempts"`
	Concurrency      int16            `json:"concurrency"`
}

func (q *Queries) CreateJob(ctx context.Context, arg CreateJobParams) (Job, error) {
	row := q.queryRow(ctx, q.createJobStmt, createJob,
		arg.Name,
		arg.DeliveryStrategy,
		arg.Attempts,
		arg.Concurrency,
	)
	var i Job
	err := row.Scan(
		&i.Name,
		&i.DeliveryStrategy,
		&i.Attempts,
		&i.Concurrency,
		&i.CreatedAt,
		&i.AutoID,
	)
	return i, err
}

const deleteAllJobs = `-- name: DeleteAllJobs :execrows
DELETE FROM jobs
`

func (q *Queries) DeleteAllJobs(ctx context.Context) (int64, error) {
	result, err := q.exec(ctx, q.deleteAllJobsStmt, deleteAllJobs)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getAllJobs = `-- name: GetAllJobs :many
SELECT name, delivery_strategy, attempts, concurrency, created_at, auto_id FROM jobs
`

func (q *Queries) GetAllJobs(ctx context.Context) ([]Job, error) {
	rows, err := q.query(ctx, q.getAllJobsStmt, getAllJobs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Job
	for rows.Next() {
		var i Job
		if err := rows.Scan(
			&i.Name,
			&i.DeliveryStrategy,
			&i.Attempts,
			&i.Concurrency,
			&i.CreatedAt,
			&i.AutoID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getJob = `-- name: GetJob :one
SELECT name, delivery_strategy, attempts, concurrency, created_at, auto_id
FROM jobs
WHERE name = $1
`

func (q *Queries) GetJob(ctx context.Context, name string) (Job, error) {
	row := q.queryRow(ctx, q.getJobStmt, getJob, name)
	var i Job
	err := row.Scan(
		&i.Name,
		&i.DeliveryStrategy,
		&i.Attempts,
		&i.Concurrency,
		&i.CreatedAt,
		&i.AutoID,
	)
	return i, err
}
