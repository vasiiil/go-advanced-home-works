package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	secret string
}

func New(secret string) *Jwt {
	return &Jwt{
		secret: secret,
	}
}

func (j *Jwt) Create(sessionId string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sessionId": sessionId,
	})
	s, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}

	return s, nil
}
