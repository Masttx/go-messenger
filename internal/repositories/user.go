package repositories

import (
	"database/sql"
	"fmt"
	"go-messenger/internal/database/queries"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(name string, email string, password string) (sql.Result, error) {
	result, err := r.db.Exec(queries.CreateUserQuery, name, email, password)
	if err != nil {
		return nil, fmt.Errorf("Error to insert user: %v", err)
	}

	return result, nil
}
