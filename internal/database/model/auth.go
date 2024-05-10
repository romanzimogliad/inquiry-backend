package model

import "github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"

type User struct {
	Id       int32  `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	RoleId   int32  `db:"role_id"`
}

func (u *User) ToDomain() *domain.User {
	return &domain.User{
		Id:       u.Id,
		Name:     u.Name,
		Password: u.Password,
		RoleId:   u.RoleId,
	}
}

type MethodToRole struct {
	Id     int32  `db:"id"`
	Method string `db:"method"`
	RoleId int32  `db:"role_id"`
}
type MethodsToRole []*MethodToRole

func (l *MethodsToRole) ToDomain() []*domain.MethodToRole {
	resp := make([]*domain.MethodToRole, len(*l))
	for k, v := range *l {
		resp[k] = v.ToDomain()
	}
	return resp
}

func (m *MethodToRole) ToDomain() *domain.MethodToRole {
	return &domain.MethodToRole{
		Id:     m.Id,
		Method: m.Method,
		RoleId: m.RoleId,
	}
}
