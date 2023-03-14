package userUsecases

import (
	"github.com/Natannms/go_sqlc/domain/entities"
	"github.com/Natannms/go_sqlc/domain/respositories"
)

type CreateUserUseCase struct {
	userRepository respositories.UserRepository
}

func NewCreateUserUseCase(userRepository respositories.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository: userRepository}
}

func (uc *CreateUserUseCase) Execute(user *entities.User) error {
	// Implementar a lógica para validar os dados do usuário,
	// EX: verificar se o e-mail já está cadastrado ou se a senha é válida.

	// Se os dados estiverem válidos, salvar o usuário no repositório.
	if err := uc.userRepository.Save(user); err != nil {
		return err
	}

	return nil
}
