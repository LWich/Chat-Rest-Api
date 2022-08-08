package tokenmanager

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// Manager ...
type Manager struct {
	signinKey string
}

// NewManager ...
func NewManager(signinKey string) *Manager {
	return &Manager{
		signinKey: signinKey,
	}
}

// NewJWT ...
func (m *Manager) NewJWT(userId int, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			Subject:   strconv.Itoa(userId),
		},
	)

	return token.SignedString([]byte(m.signinKey))
}

// NewRefreshToken ...
func (m *Manager) NewRefreshToken() string {
	return uuid.NewString()
}

// Parse ...
func (m *Manager) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", t.Header["alg"].(string))
		}

		if t.Valid {
			return nil, errors.New("token is invalid")
		}

		return []byte(m.signinKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}
