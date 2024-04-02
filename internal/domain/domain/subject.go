package domain

type Subject struct {
	Id   int32
	Name string
}

func (s *Subject) GetId() int32 {
	return s.Id
}
func (s *Subject) GetName() string {
	return s.Name
}
