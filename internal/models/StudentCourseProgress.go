package models

type StudentCourseProgress struct {
	ID        int  `db:"id"`
	StudentID int  `db:"student_id"`
	CourseID  int  `db:"course_id"`
	Progress  int  `db:"progress"`
	Completed bool `db:"completed"`
}
