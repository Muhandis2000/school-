// internal/models/schedule.go
package models

import "time"

type Schedule struct {
	ID        int       `db:"id"`
	LessonID  int       `db:"lesson_id"`
	TeacherID int       `db:"teacher_id"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
}
