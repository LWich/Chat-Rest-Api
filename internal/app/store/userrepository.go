package store

import (
	"github.com/LWich/chat-rest-api/internal/app/model"
	"github.com/gorilla/sessions"
)

// UserRepository
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.Id)
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email=$1",
		email,
	).Scan(
		&u.Id,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		return nil, err
	}

	return u, nil
}

// SetCookie ...
func (r *UserRepository) SetRefreshTokenBySession(userId int, session *sessions.Session) error {
	refreshToken := session.Values["refreshToken"]
	expiresIn := session.Options.MaxAge

	_, err := r.store.db.Exec(
		"UPDATE users SET refresh_token=$1, expires_in=$2 WHERE id=$3",
		refreshToken,
		expiresIn,
		userId,
	)

	return err
}
