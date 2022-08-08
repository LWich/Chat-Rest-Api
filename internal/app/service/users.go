package service

import (
	"errors"
	"time"

	"github.com/LWich/chat-rest-api/internal/app/model"
	"github.com/LWich/chat-rest-api/internal/app/store"
	"github.com/LWich/chat-rest-api/pkg/auth"
)

var (
	ErrPasswordOrEmailIncorrect = errors.New("password or email incorrect")
	ErrFailedToCreateTokens     = errors.New("failed to create tokens")
)

// Users ...
type UsersService struct {
	tokenManager    auth.TokenManager
	accessTokenTTL  int
	refreshTokenTTL int
	userRepository  *store.UserRepository
}

// NewUsersService ...
func NewUsersService(tokenManager auth.TokenManager, accessTokenTTL int, refreshTokenTTL int, repo *store.UserRepository) *UsersService {
	return &UsersService{
		tokenManager:    tokenManager,
		accessTokenTTL:  accessTokenTTL,
		userRepository:  repo,
		refreshTokenTTL: refreshTokenTTL,
	}
}

// SignUp ...
func (us *UsersService) SignUp(input UsersSignUpInput) (*model.User, error) {
	u := &model.User{
		Email:    input.Email,
		Password: input.Password,
	}

	if err := us.userRepository.Create(u); err != nil {
		return nil, err
	}

	u.Sanitize()

	return u, nil
}

// RefreshTokens
func (us *UsersService) RefreshTokens(refreshToken string) (*auth.Tokens, error) {
	u, err := us.userRepository.FindByRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	tokens, err := us.createTokens(u.Id)
	if err != nil {
		return nil, ErrFailedToCreateTokens
	}

	return tokens, nil
}

// SignIn ...
func (us *UsersService) SignIn(input UsersSignInInput) (*auth.Tokens, error) {
	u, err := us.userRepository.FindByEmail(input.Email)
	if err != nil || !u.ComparerPassword(input.Password) {
		return nil, ErrPasswordOrEmailIncorrect
	}

	tokens, err := us.createTokens(u.Id)
	if err != nil {
		return nil, ErrFailedToCreateTokens
	}

	return tokens, nil
}

func (us *UsersService) createTokens(userId int) (*auth.Tokens, error) {
	var (
		res auth.Tokens
		err error
	)

	res.AccessToken, err = us.tokenManager.NewJWT(userId, time.Duration(us.accessTokenTTL))
	if err != nil {
		return nil, err
	}

	res.RefreshToken = us.tokenManager.NewRefreshToken()

	return &res, us.userRepository.SetRefreshTokenAndExpiresIn(
		userId,
		us.refreshTokenTTL,
		res.RefreshToken,
	)
}
