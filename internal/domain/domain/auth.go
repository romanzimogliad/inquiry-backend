package domain

type LoginRequest struct {
	Login    string
	Password string
}
type LoginResponse struct {
	Token string
}
type User struct {
	Id       int32
	Name     string
	Password string
	RoleId   int32
}

type MethodToRole struct {
	Id     int32
	Method string
	RoleId int32
}
