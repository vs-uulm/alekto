package user

type Role string

const (
	Default Role = "all"
	Student Role = "student"
	Admin   Role = "admin"
	Teacher Role = "teacher"
)