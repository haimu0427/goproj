package usermanager

type User struct {
	Name     string
	Password string
}

func NewUser(name string, password string) *User {
	return &User{
		Name:     name,
		Password: password,
	}
}
