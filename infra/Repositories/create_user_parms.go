package repositories

type CreateUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p CreateUserParams) GetName() string {
	return p.Name
}

func (p CreateUserParams) GetEmail() string {
	return p.Email
}

func (p CreateUserParams) GetPassword() string {
	return p.Password
}
