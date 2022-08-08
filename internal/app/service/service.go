package service

import (
	"github.com/LWich/chat-rest-api/internal/app/model"
	"github.com/LWich/chat-rest-api/internal/app/store"
	"github.com/LWich/chat-rest-api/pkg/auth"
)

// Users ...
type Users interface {
	SignUp(UsersSignUpInput) (*model.User, error)
	SignIn(UsersSignInInput) (*auth.Tokens, error)
	RefreshTokens(string) (*auth.Tokens, error)
}

// Service
type Service struct {
	Users Users

	store *store.Store
}

// UsersSignInInput ...
type UsersSignInInput struct {
	Email    string
	Password string
}

// UsersSignUpInput ...
type UsersSignUpInput struct {
	Email    string
	Password string
}

// New ...
func New(store *store.Store, tokenManager auth.TokenManager, accessTokenTTL int, refreshTokenTTL int) *Service {
	s := &Service{
		store: store,
	}

	s.Users = NewUsersService(tokenManager, accessTokenTTL, refreshTokenTTL, store.User())

	return s
}
