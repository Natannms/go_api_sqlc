package repositories

import "context"

type IUserRepository interface {
	CreateUser(ctx context.Context, arg ICreateUserParams) error
	// outras funções de usuário
}
