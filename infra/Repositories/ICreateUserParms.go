package repositories

type ICreateUserParams interface {
	GetName() string
	GetEmail() string
	GetPassword() string
}


