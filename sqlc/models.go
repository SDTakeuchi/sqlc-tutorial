// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID      int32          `db:"id" json:"id"`
	Title   string         `db:"title" json:"title"`
	Memo    sql.NullString `db:"memo" json:"memo"`
	IsDone  bool           `db:"is_done" json:"is_done"`
	DueDate time.Time      `db:"due_date" json:"due_date"`
}
