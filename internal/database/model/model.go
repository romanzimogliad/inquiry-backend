package model

type TableName string

func (t TableName) String() string {
	return string(t)
}

const (
	LessonTableName           TableName = "lesson"
	SubjectTableName          TableName = "subject"
	ConceptTableName          TableName = "concept"
	SkillTableName            TableName = "skill"
	UnitTableName             TableName = "unit"
	MaterialToLessonTableName TableName = "material_to_lesson"
	GradeTableName            TableName = "grade"
)
