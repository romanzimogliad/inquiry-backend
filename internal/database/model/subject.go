package model

import "github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"

type Subject struct {
	Id   int32  `db:"id"`
	Name string `db:"name"`
}

type Subjects []*Subject

func (l *Subjects) ToDomain() []*domain.Subject {
	resp := make([]*domain.Subject, len(*l))
	for k, v := range *l {
		resp[k] = v.ToDomain()
	}
	return resp
}

func (l *Subject) ToDomain() *domain.Subject {
	return &domain.Subject{
		Id:   l.Id,
		Name: l.Name,
	}
}
