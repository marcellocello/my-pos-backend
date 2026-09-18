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
	query := `SELECT id, username, password_hash, full_name, role FROM users WHERE username = ?`
	err := r.db.QueryRow(query, req.Username).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.FullName, &user.Role)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}
