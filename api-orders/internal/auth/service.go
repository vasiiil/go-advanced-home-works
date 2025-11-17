package auth

import (
	"api-orders/internal/models"
	"api-orders/internal/user"
	"api-orders/pkg/sms"
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthServiceDeps struct {
	SmsSender      *sms.Sms
	UserRepository *user.UserRepository
}

type AuthService struct {
	SmsSender      *sms.Sms
	UserRepository *user.UserRepository
	sessions       TVerifyCodeMap
}

type TVerifyCodeMap map[string]uint16

func NewService(deps AuthServiceDeps) *AuthService {
	return &AuthService{
		SmsSender:      deps.SmsSender,
		UserRepository: deps.UserRepository,
		sessions:       make(TVerifyCodeMap),
	}
}

func (service *AuthService) Login(phone string) (string, error) {
	var sessionId string
	existsUser, err := service.UserRepository.GetByPhone(phone)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}
	if existsUser != nil {
		sessionId = existsUser.SessionId
	} else {
		sessionId = uuid.New().String()
		newUser := &models.User{
			Phone:     phone,
			SessionId: sessionId,
		}
		_, err = service.UserRepository.Create(newUser)
		if err != nil {
			return "", err
		}
	}

	code := rand.UintN(9000) + 1000
	err = service.SmsSender.Send(phone, fmt.Sprintf("Код подтверждения: %d", code))
	if err != nil {
		return "", err
	}

	service.sessions[sessionId] = uint16(code)
	return sessionId, nil
}

func (service *AuthService) VerifyCode(sessionId string, code uint16) bool {
	sessionCode, ok := service.sessions[sessionId]
	if !ok {
		return false
	}
	if code != sessionCode {
		return false
	}
	delete(service.sessions, sessionId)
	return true
}

func (service *AuthService) PrintSessions(message string) {
	fmt.Printf("#### sessions %s ####\n", message)
	for k, v := range service.sessions {
		fmt.Printf("%s: %v\n", k, v)
	}
	fmt.Println()
}
