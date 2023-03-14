package respositories

import "github.com/Natannms/go_sqlc/domain/entities"

type UserRepository interface {
	Save(user *entities.User) error
}
