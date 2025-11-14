package auth

import (
	"api-orders/pkg/sms"
	"fmt"
	"math/rand/v2"

	"github.com/google/uuid"
)

type AuthServiceDeps struct {
	SmsSender *sms.Sms
}

type TVerifyCode struct {
	phone string
	code  uint16
}
type TVerifyCodeMap map[string]TVerifyCode

type AuthService struct {
	SmsSender *sms.Sms
	sessions  TVerifyCodeMap
}

func NewService(deps AuthServiceDeps) *AuthService {
	return &AuthService{
		SmsSender: deps.SmsSender,
		sessions:  make(TVerifyCodeMap),
	}
}

func (service *AuthService) Login(phone string) (string, error) {
	code := rand.UintN(9000) + 1000
	err := service.SmsSender.Send(phone, fmt.Sprintf("Код подтверждения: %d", code))
	if err != nil {
		return "", err
	}
	sessionId := uuid.New().String()
	service.sessions[sessionId] = TVerifyCode{
		phone: phone,
		code:  uint16(code),
	}
	return sessionId, nil
}

func (service *AuthService) VerifyCode(sessionId string, code uint16) string {
	sess, ok := service.sessions[sessionId]
	if !ok {
		return ""
	}
	if !(code == sess.code) {
		return ""
	}
	delete(service.sessions, sessionId)
	return sess.phone
}

func (service *AuthService) PrintSessions(message string) {
	fmt.Printf("#### sessions %s ####\n", message)
	for k, v := range service.sessions {
		fmt.Printf("%s: %v\n", k, v)
	}
	fmt.Println()
}
