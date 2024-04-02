package domain

type DictionaryType int32

const (
	UnknownDictionaryType DictionaryType = iota
	SubjectDictionaryType
	UnitDictionaryType
	ConceptDictionaryType
	SkillDictionaryType
)

type IdName struct {
	Id   int32
	Name string
}

func (s *IdName) GetId() int32 {
	return s.Id
}
func (s *IdName) GetName() string {
	return s.Name
}
