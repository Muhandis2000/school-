// internal/models/lesson.go
package models

type Lesson struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}
