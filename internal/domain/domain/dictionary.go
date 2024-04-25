package domain

type DictionaryType int32

const (
	UnknownDictionaryType DictionaryType = iota
	SubjectDictionaryType
	UnitDictionaryType
	ConceptDictionaryType
	SkillDictionaryType
	GradeDictionaryType
	AllDictionaryType
)

type Dictionary struct {
	Type  DictionaryType
	Pairs []*IdName
}
