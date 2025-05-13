package user

type User struct {
	ID   int
	Name string
}

type UserRepo interface {
	FindByID(id int) (*User, error)
}
