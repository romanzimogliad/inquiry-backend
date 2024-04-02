package model

import "github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"

type IdName struct {
	Id   int32  `db:"id"`
	Name string `db:"name"`
}

type IdNames []*IdName

func (l *IdNames) ToDomain() []*domain.IdName {
	resp := make([]*domain.IdName, len(*l))
	for k, v := range *l {
		resp[k] = v.ToDomain()
	}
	return resp
}

func (l *IdName) ToDomain() *domain.IdName {
	return &domain.IdName{
		Id:   l.Id,
		Name: l.Name,
	}
}
