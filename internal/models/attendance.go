// internal/models/attendance.go
package models

import "time"

type Attendance struct {
	ID        int       `db:"id"`
	StudentID int       `db:"student_id"`
	LessonID  int       `db:"lesson_id"`
	Date      time.Time `db:"date"`
	Present   bool      `db:"present"`
}
