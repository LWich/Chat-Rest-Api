package tokenmanager

import (
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
