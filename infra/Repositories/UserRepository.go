package repositories

import "database/sql"


type userRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
    return &userRepository{
        db: db,
    }
}


