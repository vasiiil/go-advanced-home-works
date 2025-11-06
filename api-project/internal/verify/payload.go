package verify

type VerifyRequest struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required"`
}

type VerifyResponse struct {
	Success bool `json:"success"`
}