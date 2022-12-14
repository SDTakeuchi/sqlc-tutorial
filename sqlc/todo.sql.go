// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: todo.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO
  todos (
    title,
    memo,
    is_done,
    due_date
  )
VALUES
  ($1, $2, $3, $4)
RETURNING id, title, memo, is_done, due_date
`

type CreateTodoParams struct {
	Title   string         `db:"title" json:"title"`
	Memo    sql.NullString `db:"memo" json:"memo"`
	IsDone  bool           `db:"is_done" json:"is_done"`
	DueDate time.Time      `db:"due_date" json:"due_date"`
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo,
		arg.Title,
		arg.Memo,
		arg.IsDone,
		arg.DueDate,
	)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Memo,
		&i.IsDone,
		&i.DueDate,
	)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1
`

func (q *Queries) DeleteTodo(ctx context.Context, deletingTodoID int32) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, deletingTodoID)
	return err
}

const getTodo = `-- name: GetTodo :one
SELECT id, title, memo, is_done, due_date FROM todos
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTodo(ctx context.Context, id int32) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodo, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Memo,
		&i.IsDone,
		&i.DueDate,
	)
	return i, err
}

const listTodos = `-- name: ListTodos :many
SELECT id, title, memo, is_done, due_date FROM todos
ORDER BY due_date
`

func (q *Queries) ListTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Todo{}
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Memo,
			&i.IsDone,
			&i.DueDate,
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

const updateTodo = `-- name: UpdateTodo :one
UPDATE
  todos
SET
  title = COALESCE($1, title),
  memo = COALESCE($2, memo),
  is_done = COALESCE($3, is_done),
  due_date = COALESCE($4, due_date)
WHERE
  id = $5
RETURNING id, title, memo, is_done, due_date
`

type UpdateTodoParams struct {
	Title   sql.NullString `db:"title" json:"title"`
	Memo    sql.NullString `db:"memo" json:"memo"`
	IsDone  sql.NullBool   `db:"is_done" json:"is_done"`
	DueDate sql.NullTime   `db:"due_date" json:"due_date"`
	ID      int32          `db:"id" json:"id"`
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, updateTodo,
		arg.Title,
		arg.Memo,
		arg.IsDone,
		arg.DueDate,
		arg.ID,
	)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Memo,
		&i.IsDone,
		&i.DueDate,
	)
	return i, err
}
