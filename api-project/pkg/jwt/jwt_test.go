package jwt_test

import (
	"api-project/pkg/jwt"
	"testing"
)

func TestCreate(t *testing.T) {
	const email = "a@a.ru"
	secret := "6KzFmPtGZINF9pRs2L6hlOwwvFzGawrYF7PsImVjxNw="
	jwtService := jwt.New(secret)
	token, err := jwtService.Create(jwt.JwtData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}

	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token is invalid")
	}
	if data.Email != email {
		t.Fatalf("Email %s not equal %s", data.Email, email)
	}
}
