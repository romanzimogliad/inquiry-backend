package model

type TableName string

func (t TableName) String() string {
	return string(t)
}

const (
	LessonTableName TableName = "lesson"
)
