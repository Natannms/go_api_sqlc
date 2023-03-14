package entities

import "fmt"

type User struct {
	ID       int    `json:"ID"`
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func (u *User) SetPassword(Password string) error {
	if len(Password) < 6 {
		return fmt.Errorf("a senha deve conter pelo menos 6 caracteres")
	}

	u.Password = Password

	return nil
}
