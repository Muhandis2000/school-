// internal/models/grade.go
package models

import "time"

type Grade struct {
	ID        int       `db:"id"`
	StudentID int       `db:"student_id"`
	TeacherID int       `db:"teacher_id"`
	LessonID  int       `db:"lesson_id"`
	Value     string    `db:"value"`
	Date      time.Time `db:"date"`
}
