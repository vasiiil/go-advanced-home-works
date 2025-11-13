package auth

import (
	"api-orders/configs"
	"api-orders/pkg/sms"
	"fmt"
	"math/rand/v2"

	"github.com/google/uuid"
)

type AuthServiceDeps struct {
	Sms *configs.SmsConfig
}

type AuthService struct {
	sms      *configs.SmsConfig
	sessions map[string]uint16
}

func NewService(deps AuthServiceDeps) *AuthService {
	return &AuthService{
		sms:      deps.Sms,
		sessions: make(map[string]uint16),
	}
}

func (service *AuthService) Login(phone string) (string, error) {
	code := rand.UintN(9000) + 1000
	smsSender := sms.New(service.sms)
	err := smsSender.Send(phone, fmt.Sprintf("Код подтверждения: %d", code))
	if err != nil {
		return "", err
	}
	sessionId := uuid.New().String()
	service.sessions[sessionId] = uint16(code)
	return sessionId, nil
}

func (service *AuthService) VerifyCode(sessionId string, code uint16) bool {
	if !(code == service.sessions[sessionId]) {
		return false
	}
	delete(service.sessions, sessionId)
	return true
}

func (service *AuthService) PrintSessions(message string) {
	fmt.Printf("#### sessions %s ####\n", message)
	for k, v := range service.sessions {
		fmt.Printf("%s: %d\n", k, v)
	}
	fmt.Println()
}
