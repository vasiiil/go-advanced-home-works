package auth

type LoginRequest struct {
	Phone string `json:"phone" validate:"required"`
}

type LoginResponse struct {
	SessionId string `json:"SessionId"`
}

type VerifyCodeRequest struct {
	SessionId string `json:"sessionId" validate:"required"`
	Code      uint16 `json:"code" validate:"required,min=1000,max=9999"`
}

type VerifyCodeResponse struct {
	Token string `json:"token"`
}
