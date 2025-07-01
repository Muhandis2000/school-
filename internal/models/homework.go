// internal/models/homework.go
package models

import "time"

type Homework struct {
	ID          int       `db:"id"`
	StudentID   int       `db:"student_id"`
	LessonID    int       `db:"lesson_id"`
	Content     string    `db:"content"`
	SubmittedAt time.Time `db:"submitted_at"`
	Grade       string    `db:"grade"`
}
