package domain

import "time"

type Lesson struct {
	Id          int32
	Name        string
	UserId      int64
	Description string
	GradeId     int32
	SubjectId   int32
	CreatedAt   time.Time
}

type ListLessonsRequest struct {
	UserId int64
}
