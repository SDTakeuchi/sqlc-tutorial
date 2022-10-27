-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY due_date;

-- name: CreateTodo :one
INSERT INTO
  todos (
    title,
    memo,
    is_done,
    due_date
  )
VALUES
  ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateTodo :one
UPDATE
  todos
SET
  title = COALESCE(sqlc.narg(title), title),
  memo = COALESCE(sqlc.narg(memo), memo),
  is_done = COALESCE(sqlc.narg(is_done), is_done),
  due_date = COALESCE(sqlc.narg(due_date), due_date)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = @deleting_todo_id;
