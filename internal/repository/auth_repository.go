package repository

import (
	"database/sql"
	"mypos-backend/internal/model"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) Login(req *model.LoginRequest) (*model.User, error) {
	var user model.User
	query := `SELECT id, username, full_name, role FROM users WHERE username = ? AND password = ?`
	err := r.db.QueryRow(query, req.Username, req.Password).Scan(&user.ID, &user.Username, &user.FullName, &user.Role)

	if err != nil {
		return nil, err
	}
	return &user, err
}
