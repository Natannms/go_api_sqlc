package usecase

import "github.com/Natannms/go_sqlc/domain/entities"

type UserRepository interface {
	Save(user *entities.User) error
}

type CreateUserUseCase struct {
	userRepository UserRepository
}

func NewCreateUserUseCase(userRepository UserRepository) *CreateUserUseCase {
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
