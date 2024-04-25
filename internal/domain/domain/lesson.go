package domain

import "time"

const DefaultSize = 9
const DefaultPage = 1

type Lessons struct {
	Count   int32
	Lessons []*Lesson
}

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
	GradeId    int32
	SearchText string
}

type ListLessonsRequest struct {
	UserId int32
	Filter Filter
	Page   Page
}
type Page struct {
	Page int32
	Size int32
}

type GetLessonsRequest struct {
	UserId int32
	Id     string
}
