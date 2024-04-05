package domain

import "time"

type Lesson struct {
	Id          string
	Unit        *IdName
	Duration    int32
	Text        string
	Name        string
	UserId      int32
	Concept     *IdName
	Skill       *IdName
	Description string
	GradeId     int32
	Subject     *IdName
	CreatedAt   time.Time
	Image       *File
	Files       []*File
}
type Filter struct {
	SubjectId  int32
	ConceptId  int32
	UnitId     int32
	SkillId    int32
	SearchText string
}

type ListLessonsRequest struct {
	UserId int32
	Filter Filter
}

type GetLessonsRequest struct {
	UserId int32
	Id     string
}
