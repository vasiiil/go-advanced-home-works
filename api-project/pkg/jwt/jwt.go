package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	secret string
}

type JwtData struct {
	Email string
}

func New(secret string) *Jwt {
	return &Jwt{
		secret: secret,
	}
}

func (j *Jwt) Create(data JwtData) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
	})
	s, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}

	return s, nil
}

func (j *Jwt) Parse(tokenString string) (bool, *JwtData) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return false, nil
	}
	email := token.Claims.(jwt.MapClaims)["email"]
	return token.Valid, &JwtData{Email: email.(string)}
}
