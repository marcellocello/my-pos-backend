package repository

import (
	"database/sql"
	"mypos-backend/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	query := `INSERT INTO users (username, password_hash, full_name, role) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, user.Username, user.PasswordHash, user.FullName, user.Role)
	return err
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	query := `SELECT id, username, password_hash, full_name, role, created_at, updated_at FROM users WHERE username = ?`
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.FullName, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepository) GetByID(id int) (*model.User, error) {
	var user model.User
	query := `SELECT id, username, password_hash, full_name, role FROM users WHERE id = ?`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.FullName, &user.Role)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, err
}

func (r *UserRepository) UpdatePassword(id int, newPasswordHash string) error {
	query := `UPDATE users SET password_hash = ? WHERE id = ?`
	_, err := r.db.Exec(query, newPasswordHash, id)

	return err
}
