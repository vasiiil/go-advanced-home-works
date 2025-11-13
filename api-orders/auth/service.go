package auth

import (
	"fmt"
	"math/rand/v2"

	"github.com/google/uuid"
)

type AuthService struct {
	sessions map[string]uint16
}

func NewService() *AuthService {
	return &AuthService{
		sessions: make(map[string]uint16),
	}
}

func (service *AuthService) Login(phone string) string {
	code := rand.UintN(9000) + 1000
	sessionId := uuid.New().String()
	service.sessions[sessionId] = uint16(code)
	return sessionId
}

func (service *AuthService) VerifyCode(sessionId string, code uint16) bool {
	return code == service.sessions[sessionId]
}

func (service *AuthService) PrintSessions(message string) {
	fmt.Printf("#### sessions %s ####\n", message)
	for k, v := range service.sessions {
		fmt.Printf("%s: %d\n", k, v)
	}
	fmt.Println()
}
