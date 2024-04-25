package domain

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
