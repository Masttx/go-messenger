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

func (r *UserRepository) Read(id int64) (types.User, error) {
	var user types.User
	err := r.db.Get(&user, queries.ReadUserQuery, id)
	if err != nil {
		return types.User{}, fmt.Errorf("Error to find user: %v", err)
	}

	return user, nil
}

func (r *UserRepository) Update(id int64, nickname string) error {
	err := r.db.Exec(queries.UpdateUserQuery, nickname, id)
	if err != nil {
		return fmt.Errorf("Error to update user: %v", err)
	}

	return nil
}

func (r *UserRepository) Delete(id int64) error {
	res, err := r.db.Exec(queries.DeleteUserQuery, id)
	if err != nil {
		return fmt.Errorf("Error to delete user: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error checking rows affected: %v", err)
	}

	if rowsAffected == 0 {
		// return fmt.Errorf("Nenhuma linha deletada. Usuário com ID %d não encontrado", id) CORRIGIR
	}

	return nil
}

func (r *UserRepository) List(ctx context.Context) ([]types.User, error) {
	var result []types.User
	err := r.db.SelectContext(ctx, &result, queries.ListUsersQuery)
	if err != nil {
		return nil, fmt.Errorf("Error to list users: %v", err)
	}

	return result, nil
}

